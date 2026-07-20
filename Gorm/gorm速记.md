# GORM V2 全套速查表

说明：所有代码可直接复制使用，重点用法标红，适配日常开发+面试查阅。

## 一、初始化 & 基础配置（必写）

```go
import (
  "gorm.io/driver/mysql"   // MySQL驱动（其他数据库替换对应驱动）
  "gorm.io/gorm"
  "gorm.io/gorm/logger"
  "time"
)

// 1. MySQL DSN（核心，替换自己的数据库信息）
dsn := "user:pass@tcp(127.0.0.1:3306)/db_name?charset=utf8mb4&parseTime=True&loc=Local"

// 2. 初始化DB连接（全局唯一）
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
  PrepareStmt: true,  // 全局预编译，提升重复执行效率
  // 关闭自动创建外键（可选，多数项目禁用）
  DisableForeignKeyConstraintWhenMigrating: true,
  // 日志级别（开发用Info，生产用Silent）
  Logger: logger.Default.LogMode(logger.Info), 
})

if err != nil {
  panic("数据库连接失败: " + err.Error())
}

// 可选：获取底层sql.DB，设置连接池
sqlDB, _ := db.DB()
sqlDB.SetMaxIdleConns(10)    // 最大空闲连接
sqlDB.SetMaxOpenConns(100)   // 最大打开连接
sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大存活时间
```

## 二、Session 常用配置（临时生效，不污染全局）

```go
// 临时会话（仅当前链式调用生效）
tx := db.Session(&gorm.Session{
  DryRun:      true,    // 只生成SQL不执行，用于调试
  PrepareStmt: true,    // 会话级预编译（局部加速）
  SkipHooks:   true,    // 跳过钩子函数（新增/修改/删除时）
  Context:     ctx,     // 上下文（超时控制、追踪等）
})

// 常用：获取生成的SQL（DryRun模式下）
stmt := tx.First(&user, 1).Statement
sqlStr := stmt.SQL.String() // 生成的SQL语句
sqlVars := stmt.Vars        // SQL参数（防注入）
```

## 三、模型定义（结构体标签，核心）

```go
import "gorm.io/gorm"

// 基础模型（可嵌入，简化字段）
type BaseModel struct {
  ID        uint           `gorm:"primaryKey;autoIncrement"` // 主键自增
  CreatedAt time.Time      // 自动填充创建时间
  UpdatedAt time.Time      // 自动填充更新时间
  DeletedAt gorm.DeletedAt `gorm:"index"`                  // 软删除标记（开启软删除必须加）
}

// 示例：用户模型（嵌入BaseModel，含软删除、约束、索引）
type User struct {
  BaseModel                // 嵌入基础模型，继承ID/CreatedAt等字段
  Name      string         `gorm:"column:user_name;size:32;not null;unique"` // 列名、长度、非空、唯一
  Age       int            `gorm:"check:age_gt_0,age > 0;default:18"`        // 检查约束、默认值
  Email     string         `gorm:"uniqueIndex;size:128"`                     // 唯一索引
  Status    int            `gorm:"type:tinyint;default:1"`                  // 字段类型、默认值
  Phone     string         `gorm:"index"`                                   // 普通索引
  // 关联字段（后续关联查询用）
  AddressID uint           `gorm:"foreignKey:AddressID"`                    // 外键（关联Address表）
  Address   Address        `gorm:"foreignKey:AddressID;references:ID"`      // 一对一关联
  Orders    []Order        `gorm:"foreignKey:UserID"`                       // 一对多关联
}

// 可选：自定义表名（默认是结构体复数形式，如User→users）
func (User) TableName() string {
  return "sys_user"
}
```

## 四、迁移 / 表结构操作（AutoMigrate & Migrator）

```go
// 1. 自动迁移（最常用，安全：只增不减，不删列/删表）
db.AutoMigrate(&User{}, &Order{}, &Address{})

// 2. 创建表时指定引擎（MySQL专属）
db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

// 3. 手动操作表结构（Migrator，精细控制）
m := db.Migrator()

// 表操作
m.HasTable(&User{})        // 判断表是否存在（返回bool）
m.CreateTable(&User{})     // 创建表
m.DropTable(&User{})       // 删除表（危险，谨慎使用）
m.RenameTable(&User{}, "sys_user") // 重命名表

// 列操作
m.AddColumn(&User{}, "avatar")     // 新增列
m.DropColumn(&User{}, "avatar")    // 删除列（危险）
m.AlterColumn(&User{}, "age")      // 修改列（类型、长度等）
m.RenameColumn(&User{}, "avatar", "avatar_url") // 重命名字段
m.HasColumn(&User{}, "avatar")    // 判断列是否存在

// 约束操作
m.CreateConstraint(&User{}, "age_gt_0") // 创建约束（对应模型中check标签）
m.DropConstraint(&User{}, "age_gt_0")  // 删除约束
m.HasConstraint(&User{}, "age_gt_0")   // 判断约束是否存在

// 索引操作
m.CreateIndex(&User{}, "idx_phone")    // 创建索引
m.DropIndex(&User{}, "idx_phone")     // 删除索引
m.HasIndex(&User{}, "idx_phone")      // 判断索引是否存在
```

## 五、CRUD 基础操作（高频）

### 5.1 创建（Create）

```go
// 1. 单条创建
user := User{Name: "张三", Age: 20, Email: "zhangsan@xxx.com"}
db.Create(&user) // 创建后，user.ID会自动填充（自增主键）

// 2. 批量创建（高效，推荐）
users := []User{
  {Name: "李四", Age: 22},
  {Name: "王五", Age: 21},
}
db.CreateInBatches(users, 100) // 每次批量创建100条

// 3. 自定义创建字段（只创建指定字段）
db.Select("name", "age").Create(&user)
```

### 5.2 查询（Read）

```go
var user User
var users []User

// 单条查询
db.First(&user, 1)                // 根据主键查询（id=1），无数据返回ErrRecordNotFound
db.Take(&user, "name = ?", "张三") // 条件查询（随机返回一条）
db.Last(&user)                    // 查询最后一条数据

// 批量查询
db.Find(&users)                   // 查询所有
db.Where("age > ?", 18).Find(&users) // 条件查询（age>18）

// 条件组合
db.Where("age = ?", 20).Or("name = ?", "张三").Find(&users) // 或条件
db.Not("age = 18").Find(&users)                          // 非条件

// 排序、分页、指定字段
db.Order("age desc").Limit(10).Offset(20).Select("name", "age").Find(&users)
// Limit(10)：每页10条；Offset(20)：跳过前20条（第3页）

// 聚合查询
var count int64
db.Model(&User{}).Where("age > 18").Count(&count) // 统计数量
db.Model(&User{}).Max("age", &maxAge)             // 最大值
db.Model(&User{}).Min("age", &minAge)             // 最小值
db.Model(&User{}).Avg("age", &avgAge)             // 平均值
```

### 5.3 更新（Update）

```go
// 1. 单条更新（根据主键）
db.Model(&User{}).Where("id = ?", 1).Update("age", 21)

// 2. 批量更新（条件更新）
db.Model(&amp;User{}).Where("age < 18").Update("status", 0)

// 3. 多字段更新（结构体）
db.Model(&user).Updates(User{Name: "张三2", Age: 22}) // 忽略零值
db.Model(&user).Updates(map[string]interface{}{"name": "张三2", "age": 22}) // 不忽略零值

// 4. 指定更新字段（只更新name）
db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "张三3", "age": 23})

// 5. 忽略更新字段（不更新age）
db.Model(&user).Omit("age").Updates(map[string]interface{}{"name": "张三3", "age": 23})
```

### 5.4 删除（Delete）

```go
// 1. 物理删除（彻底删除，谨慎使用）
db.Delete(&User{}, 1)                  // 根据主键删除
db.Where("age < 18").Delete(&User{})   // 条件物理删除

// 2. 软删除（开启后，默认查询不显示已删除数据）
// 前提：模型中必须有 DeletedAt gorm.DeletedAt `gorm:"index"` 字段
db.Delete(&user, 1) // 软删除，实际是更新DeletedAt字段为当前时间
db.Unscoped().Delete(&user, 1) // 强制物理删除（软删除模型下）

// 3. 恢复软删除数据
db.Unscoped().Model(&User{}).Where("id = ?", 1).Update("deleted_at", nil)
```

## 六、事务（核心，保证数据一致性）

```go
// 方式1：手动控制事务（推荐，灵活）
func UpdateUserAndOrder(db *gorm.DB) error {
  // 1. 开启事务
  tx := db.Begin()
  defer func() {
    // 异常回滚（防止panic导致事务未关闭）
    if r := recover(); r != nil {
      tx.Rollback()
    }
  }()

  // 2. 检查事务开启是否成功
  if tx.Error != nil {
    return tx.Error
  }

  // 3. 事务内操作（所有操作都用tx，而非db）
  // 示例：修改用户年龄 + 创建订单
  var user User
  if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&user, 1).Error; err != nil {
    tx.Rollback() // 操作失败，回滚
    return err
  }
  user.Age += 1
  if err := tx.Save(&user).Error; err != nil {
    tx.Rollback()
    return err
  }

  order := Order{UserID: user.ID, OrderNo: "20240501001", Amount: 99.9}
  if err := tx.Create(&order).Error; err != nil {
    tx.Rollback()
    return err
  }

  // 4. 提交事务（所有操作成功，提交）
  return tx.Commit().Error
}

// 方式2：自动事务（简单场景用）
db.Transaction(func(tx *gorm.DB) error {
  // 事务内操作，用法和tx一致
  var user User
  if err := tx.First(&user, 1).Error; err != nil {
    return err // 返回错误，自动回滚
  }
  user.Age += 1
  return tx.Save(&user).Error // 返回nil，自动提交
})
```

## 七、锁（并发控制，避免脏数据）

```go
import "gorm.io/gorm/clause"

// 核心：锁必须在事务内使用，事务结束自动释放（无需手动解锁）
tx := db.Begin()
defer func() {
  if r := recover(); r != nil {
    tx.Rollback()
  }
}()

// 1. 行锁（FOR UPDATE，最常用，锁定查询行，其他事务无法修改/删除）
tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&user, 1)

// 2. 共享锁（FOR SHARE，只读锁，其他事务可查，不可改）
tx.Clauses(clause.Locking{Strength: "SHARE"}).First(&user, 1)

// 3. 避免死锁：添加NOWAIT（无锁则立即返回错误，不阻塞）
tx.Clauses(clause.Locking{Strength: "UPDATE", Options: "NOWAIT"}).First(&user, 1)

// 事务操作...
tx.Commit() // 提交/回滚后，锁自动释放
```

## 八、关联查询（一对一、一对多、多对多）

### 8.1 模型关联定义（先定义关联，再查询）

```go
// 1. 一对一：User ↔ Address（一个用户一个地址）
type Address struct {
  BaseModel
  Province string
  City     string
  UserID   uint // 外键，关联User.ID
}

type User struct {
  BaseModel
  Name      string
  AddressID uint      // 外键（可选，一对一可省略，自动匹配）
  Address   Address   `gorm:"foreignKey:AddressID;references:ID"` // 一对一关联
}

// 2. 一对多：User ↔ Order（一个用户多个订单）
type Order struct {
  BaseModel
  OrderNo string
  UserID  uint     // 外键，关联User.ID
  User    User     `gorm:"foreignKey:UserID;references:ID"` // 多对一关联（反向）
}

type User struct {
  BaseModel
  Name   string
  Orders []Order   `gorm:"foreignKey:UserID;references:ID"` // 一对多关联
}

// 3. 多对多：User ↔ Role（一个用户多个角色，一个角色多个用户）
type Role struct {
  BaseModel
  Name  string
  Users []User `gorm:"many2many:user_roles;"` // 多对多，中间表user_roles（自动创建）
}

type User struct {
  BaseModel
  Name  string
  Roles []Role `gorm:"many2many:user_roles;"` // 多对多关联
}
```

### 8.2 关联查询常用方法

```go
var user User
var users []User

// 1. 一对一查询：预加载Address（避免N+1问题）
db.Preload("Address").First(&user, 1) // user.Address 就是关联的地址信息

// 2. 一对多查询：预加载Orders
db.Preload("Orders").First(&user, 1) // user.Orders 就是该用户的所有订单

// 3. 多对多查询：预加载Roles
db.Preload("Roles").First(&user, 1) // user.Roles 就是该用户的所有角色

// 4. 嵌套预加载（订单关联商品，预加载订单+商品）
db.Preload("Orders.Goods").First(&user, 1)

// 5. 条件预加载（只加载状态为1的订单）
db.Preload("Orders", "status = ?", 1).First(&user, 1)

// 6. 关联查询（根据关联条件查主表）
// 查有地址的用户
db.Joins("Address").Find(&users)
// 查有订单的用户（订单金额>100）
db.Joins("JOIN orders ON orders.user_id = users.id").Where("orders.amount > 100").Find(&users)
```

## 九、钩子函数（Hook，生命周期回调）

说明：钩子是模型的方法，在创建/更新/删除/查询时自动执行，用于拦截操作、补充逻辑。

```go
type User struct {
  BaseModel
  Name string
  Age  int
}

// 1. 创建钩子（BeforeCreate：创建前执行；AfterCreate：创建后执行）
func (u *User) BeforeCreate(tx *gorm.DB) error {
  // 示例：创建前给Name加前缀
  u.Name = "prefix_" + u.Name
  return nil
}
func (u *User) AfterCreate(tx *gorm.DB) error {
  // 示例：创建后记录日志
  fmt.Printf("用户%d创建成功\n", u.ID)
  return nil
}

// 2. 更新钩子（BeforeUpdate：更新前；AfterUpdate：更新后）
func (u *User) BeforeUpdate(tx *gorm.DB) error {
  // 示例：禁止修改年龄小于18的用户
  if u.Age < 18 {
    return errors.New("未成年人不能修改信息")
  }
  return nil
}

// 3. 删除钩子（BeforeDelete：删除前；AfterDelete：删除后）
func (u *User) BeforeDelete(tx *gorm.DB) error {
  // 示例：禁止删除管理员
  if u.Name == "admin" {
    return errors.New("管理员不能删除")
  }
  return nil
}

// 4. 查询钩子（BeforeFind：查询前；AfterFind：查询后）
func (u *User) AfterFind(tx *gorm.DB) error {
  // 示例：查询后处理数据（如加密字段解密）
  u.Name = strings.TrimSpace(u.Name)
  return nil
}

// 注意：跳过钩子的方法（Session配置SkipHooks: true）
db.Session(&gorm.Session{SkipHooks: true}).Create(&user)
```

## 十、软删除（Soft Delete）

```go
// 1. 开启软删除（模型必须添加DeletedAt字段）
type User struct {
  BaseModel                // 已包含 DeletedAt gorm.DeletedAt `gorm:"index"`
  Name string
}

// 2. 软删除操作（实际更新DeletedAt字段）
db.Delete(&user, 1) // 软删除，db.Find(&users) 不会查询到该数据

// 3. 查询软删除数据（必须用Unscoped()）
db.Unscoped().Find(&users) // 查询所有数据（含已软删除）
db.Unscoped().Where("deleted_at IS NOT NULL").Find(&users) // 只查软删除数据

// 4. 恢复软删除数据
db.Unscoped().Model(&User{}).Where("id = ?", 1).Update("deleted_at", nil)

// 5. 强制物理删除（软删除模型下）
db.Unscoped().Delete(&user, 1)
```

## 十一、日志（自定义 & 级别）

```go
// 1. 全局设置日志级别
// 级别：logger.Silent（静音）、logger.Error（只打错误）、logger.Warn（错误+警告）、logger.Info（所有SQL）
db.Logger = logger.Default.LogMode(logger.Silent)

// 2. 自定义Logger（需实现logger.Interface接口）
type MyLogger struct{}

func (l MyLogger) LogMode(level logger.LogLevel) logger.Interface { return l }
func (l MyLogger) Info(ctx context.Context, msg string, args ...interface{}) {
  fmt.Printf("[INFO] %s\n", fmt.Sprintf(msg, args...))
}
func (l MyLogger) Warn(ctx context.Context, msg string, args ...interface{}) {
  fmt.Printf("[WARN] %s\n", fmt.Sprintf(msg, args...))
}
func (l MyLogger) Error(ctx context.Context, msg string, args ...interface{}) {
  fmt.Printf("[ERROR] %s\n", fmt.Sprintf(msg, args...))
}
// 核心：打印SQL（执行时间、SQL语句、影响行数）
func (l MyLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
  sql, rows := fc()
  fmt.Printf("耗时：%v | SQL：%s | 影响行数：%d\n", time.Since(begin), sql, rows)
}

// 使用自定义Logger
db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
  Logger: MyLogger{},
})
```

## 十二、常用技巧速记（必背）

- DryRun：只生成SQL不执行，用于调试（Session配置）
- PrepareStmt：预编译SQL，提升重复执行效率（全局/会话均可配置）
- AutoMigrate：自动同步表结构，安全（只增不减）
- Migrator：手动控制表/列/索引/约束，灵活但需谨慎
- 事务：必须用tx操作，提交/回滚自动释放锁，异常需回滚
- 关联查询：用Preload避免N+1问题，多对多需指定中间表
- 软删除：需添加DeletedAt字段，查询需用Unscoped()查已删除数据
- 钩子：可拦截操作，跳过钩子用Session{SkipHooks: true}
- 日志：开发用Info（看SQL），生产用Silent（静音）
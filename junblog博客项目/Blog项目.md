# **JUN Blog**

## **项目整体结构：**

### **前端：**

blog-frontend/

│

├── src/

│  ├── main.js

│  │

│  ├── api/        // 所有请求接口

│  │  └── article.js

│  │

│  ├── views/       // 页面

│  │  ├── Home.vue

│  │  ├── Login.vue

│  │  └── Article.vue

│  │

│  ├── components/    // 组件

│  │  └── Navbar.vue

│  │

│  ├── router/      // 路由

│  │  └── index.js

│  │

│  ├── store/       // 状态管理（Pinia/Vuex）

│  │

│  └── utils/       // axios封装

│    └── request.js

### **后端：**

blog-backend/

│

├── main.go         // 项目入口

├── go.mod

│

├── config/        // 配置管理

│  └── config.go

│  └── config.yaml

│  └── init.go

│

├── router/        // 路由注册

│  └── router.go

│

├── controller/      // 控制层（接收请求）

│  ├── user_controller.go

│  └── article_controller.go

│

├── service/        // 业务逻辑层（核心）

│  ├── user_service.go

│  └── article_service.go

│

├── dao/          // 数据访问层

│  ├── user_dao.go

│  └── article_dao.go

│

├── dto/          // 数据传输

│  ├── user_dto.go

│  └── article_dao.go

 

├── model/         // 数据模型（数据库结构）

│  ├── user.go

│  └── article.go

│

├── middleware/      // 中间件

│  ├── jwt.go

│  └── cors.go

│

├── utils/         // 工具类

│  ├── response.go

│  └── jwt.go

│

├── pkg/          // 可复用组件（可选）

│

└── uploads/        // 文件上传（头像、图片）

 

## **Model结构体定义：**

User:

```go
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"size:50;unique;not null"`
	Password  string    `json:"-"` // 不返回给前端
	Email     string    `json:"email" gorm:"size:100;unique"`
	Phone     string    `json:"phone" gorm:"size:20"`
	Role      string    `json:"role" gorm:"default:user"`
	Status    bool      `json:"status" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
```

Article:

```go
type Tag struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"size:50;unique;not null"`
	Slug      string    `json:"slug" gorm:"size:50;unique;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Articles []Article `json:"articles" gorm:"many2many:article_tags;"`
}
```

Comment:

```go
type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content" gorm:"type:text;not null"`                 //评论内容
	UserID    uint      `json:"user_id" gorm:"not null"`                           // 评论者id
	ArticleID uint      `json:"article_id" gorm:"not null"`                        //所评论文章的id
	ParentID  *uint     `json:"parent_id"`                                         //被回复人的评论id,用于实现评论回复
	Static    string    `json:"static" gorm:"type:varchar(20);default:'approved'"` // 当前评论的状态：approved(已处理), pending(待处理), deleted(已删除)
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User    User     `json:"user" gorm:"foreignKey:UserID"`
	Article Article  `json:"article" gorm:"foreignKey:ArticleID"`
	Parent  *Comment `json:"parent" gorm:"foreignKey:ParentID"` // 自关联，用于回复
}
```

Category:

```go
type Category struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"size:50;unique;not null"` // 分类名称
	Slug        string    `json:"slug" gorm:"size:50;unique;not null"` // 分类别名，用于URL
	Description string    `json:"description" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	
	Articles []Article `json:"articles" gorm:"foreignKey:CategoryID"`
}
```



Like:

```go
type Like struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"userID" gorm:"uniqueIndex:idx_user_article_like"`    //是谁点的赞
	ArticleID uint      `json:"articleId" gorm:"uniqueIndex:idx_user_article_like"` //点赞的是哪篇文章
	CreatedAt time.Time `json:"createdAt"`

	User    User    `json:"user" gorm:"foreignKey:UserID"`
	Article Article `json:"article" gorm:"foreignKey:ArticleID"`
}
```



Favorite:

```go
type Favorite struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"userID" gorm:"uniqueIndex:idx_user_article_fav"`    //谁收藏的
	ArticleID uint      `json:"articleID" gorm:"uniqueIndex:idx_user_article_fav"` //点赞的哪个文章
	CreatedAt time.Time `json:"createdAt"`

	User    User    `json:"user" gorm:"foreignKey:UserID"`
	Article Article `json:"article" gorm:"foreignKey:ArticleID"`
}
```



## 功能模块

### 一、游客模式

1.只能搜索和查看文章



### 二、访客模式

1. 登录、注册功能
2. 个人信息管理
3. 查看文章
4. 文章评论、点赞、收藏、
5. 回复评论
6. 查看点赞和收藏的文章



### 三、管理员模式

1. 登录注册
2. 个人信息管理
3. 发布文章
4. 管理文章
5. 删除文章
6. 修改文章
7. 回复评论
8. 统计文章数据（浏览、点赞、收藏数）
9. 后台访客信息管理
10. 访客数据统计





## 模块功能

### dao

dao层只关心操作数据库数据的逻辑

对数据的操作：

增

删

改

查

### router:

router层：只负责路由分发，在router中注册路由



### controller:

controller层负责处理HTTP相关的业务：

参数绑定：解析JSON和URL路径中把数据解析出来，把参数绑定到dto层

输入校验：利用binding的标签参数是否合法

进行身份校验：从JWT token中提取当前用户ID

分发任务：调用对应的Service函数

统一响应：根据Service返回的结构，格式化成统一的JSON



### Service

只负责处理核心业务逻辑，不管HTTP业务，不管使用的是什么框架写的逻辑

数据转换:把DTO转换成数据库Model

复杂逻辑处理：生成Slug、权限验证

多表联动：例如点赞时，既要给Like表加记录，也要给Article表的中数量也加一，处理多个表之间的关系

错误处理：返回具体的业务错误




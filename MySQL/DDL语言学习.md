# Mysql数据库---数据定义语言

数据定义语言

库和表的管理：创建（create），修改(alter)，删除(drop)



## 库的管理：

### 库的创建：

语法：create database 库名;



案例：创建库Books

```mysql
CREATE DATABASE books;
CREATE DATABASE IF NOT EXISTS books;
```



### 库的修改

库一般不进行修改

可以更改库的字符集：

```mysql
ALTER DATABASE books CHARACTER SET gbk;
```



### 库的删除

语法:

```mysql
DROP DATABASE books;
DROP DATABASE IF EXISTS books;
```







## 表的管理

### 表的创建

语法：

```mysql
create table 表名(
	列名 列的类型[(长度) 约束],
    列名 列的类型[(长度) 约束],
    列名 列的类型[(长度) 约束],
    ...
    列名 列的类型[(长度) 约束]
)
```



### 表的修改

语法：

```mysql
alter table 表名 add|drop|modify|change column 列名 （列类型 约束）
```





#### 修改列名 

```mysql
ALTER TABLE book 
CHANGE COLUMN(可省略) bname bookName VARCHAR(20);

```





#### 修改列的类型或约束

```mysql
ALTER TABLE book MODIFY COLUMN publishdata TIMESTAMP;
ALTER TABLE book MODIFY COLUMN 列名 要修改成的类型;
```



#### 添加新的列

```mysql
ALTER TABLE book ADD COLUMN publishdate TIMESTAMP;
ALTER TABLE book ADD COLUMN 列名 类型;
```



#### 删除列

```mysql
ALTER TABLE book DROP COLUMN A(要删除的列);
```



#### 修改表名

```mysql
ALTER TABLE book RENAME TO book_info;
```





### 表的删除

语法：

```mysql
DROP TABLE book_info;
```



通用的写法：

```mysql
DROP DATABASE IF EXISTS 旧库名;
CREATE DATABASE 新库名；

DROP TABLE IF EXISTS 旧表名；;
CREATE TABLE 表名();
```



### 表的复制



1. 仅仅复制表的结构

   ```mysql
   CREATE TABLE copy 
   LIKE author;
   ```

   

2. 复制表的结构+数据

   ```mysql
   CREATE TABLE copy2 
   SELECT * FROM author;
   ```

3. 只复制部分数据

   ```mysql
   CREATE TABLE copy3
   SELECT id,au_name 
   FROM author
   WHERE nation='中国';
   ```

   

4. 仅仅复制某些字段	

   ```mysql
   CREATE TABLE copy3
   SELECT id,au_name 
   FROM author
   WHERE 1=2;
   ```

   



## 数据类型管理

### 常见的数据类型

数值型：

​	整数

​	小数：定点数，浮点数

字符型：

​	较短的文本：char、varchar

​	较长的文本：text、blob（较长的二进制数据）

日期型：



### 整型

分类：tinyint、smallint、mediumint、int/integer、bigint

字节	1		2		3			4		8



特点：

1. 如果不设置无符号还是有符号，默认是有符号，如果想设置无符号，需要添加unsigned关键字

2. 如果插入的数值超出了整型的范围，会报out of range异常，并且插入的是临界值

3. 如果不设置长度，会设置默认的长度

   长度代表了显示的最大宽度，如果不够会用0在左边填充，但必须搭配zerofill使用



1. 如何设置无符号和有符号

   ```mysql
   CREATE TABLE tab_int(
   	t1 INT, #有符号
       t2 INT UNSIGNED #无符号
   );
   
   #ZEROFILL零填充，在数值不足设置的长度的时候，会在前方补零
   CREATE TABLE tab_int(
   	t1 INT ZEROFILL, #有符号
       t2 INT UNSIGNED #无符号
   );
   ```

   

### 小数

```mysql
CREATE TABLE tab_float(
	f1 FLOAT(5,2),
	f2 DOUBLE(5,2),
	f3 DECIMAL(5,2)


) ;
INSERT INTO tab_float VALUES(123.45,123.45,123.45);

INSERT INTO tab_float VALUES(123.456,123.456,123.456);

```



特点：

1. M和D的意思

   M:整数部位+小数部位

   D:小数部位

   如果超出范围，则插入临界值

2. M和D都可以省略

   如果是decimal,则M默认为10,D默认为0

   如果是float和double,则会根据插入的数值的精度来决定精度

3. 定点型的精确度较高，如果要求插入的数值的精度较高，如货币运算等则

#### 浮点型

类型：float(M,D)、double(M,D)

字节：	4	8



#### 定点型

decimal(M,D)简写为DEC(M,D),最大取指范围与double相同



选择的原则：

所选择的类型越简单越好，能保存数值的类型越小越好





### 字符型

#### 较短的文本：

char(M),固定长度的字符，比较耗费空间，效率高

varchar(M)，可变长度的字符，比较节省空间，效率低

M的意思：最大的字符数

char的M可以省略，默认为1，而varchar的M不可以省略



#### 其他

binary和varbinary用于保存较短的二进制



#### 较长的文本:

text

blob(较大的二进制)



#### Enum类型

枚举类型，要求插入的值必须属于列表中指定的值之一



#### Set类型

保存集合类型的，一次可以选举多个成员，而Enum只能选一个

根据成员个数不同，存储所占字节也不同





### 日期型

| 日期和时间类型 | 字节 | 最小值              | 最大值              |
| :------------- | :--- | :------------------ | :------------------ |
| date           | 4    | 1000-01-01          | 9999-12-31          |
| datetime       | 8    | 1000-01-01 00:00:00 | 9999-12-31 23:59:59 |
| timestamp      | 4    | 19700101080001      | 2038年的某个时刻    |
| time           | 3    | -838:59:59          | 838:59:59           |
| year           | 1    | 1901                | 2155                |

分类：

date:只保存日期

time:只保存时间

year:值保存年

datetime:保存日期+时间

timestmap:保存日期+时间



**datetime和timestamp的区别：**

1. Timestamp支持的时间范围较小，取值范围：
   19700101080001——2038年的某个时间

   Datetime的取值范围：1000-1-1——9999-12-31

2. timestamp和实际时区有关，更能反映实际的日期，而datetime则只能反映出插入时的当地时区

3. timestamp的属性受Mysql版本和SQLMode的影响很大





## 常见的约束

含义：一种限制，用于限制表中的数据吗，为了保证表中数据的准确和可靠性



分类：

1. NOT NULL:非空，用于保证该字段的值不能为空，例如姓名，学号

2. DEFAULT：默认，用于保证该字段有默认值，比如性别

3. PRIMARY KEY:主键，用于保证该字段的值具有唯一性，并且非空，例如学号，员工编号

4. UNIQUE:唯一，用于保证该字段的值具有唯一性，可以为空，比如座位号

5. CHECK:检查约束（MySQL中不支持）

6. FOREIGN KEY:外键，用于限制两个表的关系，用于保证该字段的值必须来自于主表的关联列的值，在从表添加外检约束，用于引用主表中某列的值

   比如学生表的专业编号，员工表的部门编号，员工表的工种编号



添加约束的时机：

​	1. 创建表时

​	2. 修改表时

约束添加的分类：

 1. 列级约束：

    ​	六大约束语法上都支持，但外键约束没有效果

 2. 表级约束：

    ​	除了非空和默认，其他的都支持



### 创建表时添加约束

主键和唯一的对比：

|      | 保证唯一性 | 是否允许为空 | 一个表中可以有多个 | 是否允许组合   |
| ---- | ---------- | ------------ | ------------------ | -------------- |
| 主键 | 允许       | 不允许       | 至多有一个         | 允许，但不推荐 |
| 唯一 | 允许       | 允许         | 可以有多个         | 允许，但不推荐 |



外键：

1. 要求在从表设置外键关系
2. 从表的外键列的类型和主表的关联列的类型要求一致或兼容，名称无要求
3. 主表的关联列必须是一个key(一般是主键或唯一键)
4. 插入数据时，先插入主表，再插入从表

#### 添加列级约束

语法：

直接在字段名和类型后面追加，约束类型即可，只支持：默认、非空、主键、唯一

```mysql
CREATE TABLE stuinfo(
	id INT PRIMARY KEY,#主键
	stuName VARCHAR(20) NOT NULL,#非空
	gender CHAR(1),
	seat INT UNIQUE,#唯一
	age INT DEFAULT 18,	#默认约束
	majorId INT,
	FOREIGN KEY (majorId) REFERENCES major(id)#外键
);

CREATE TABLE major(
	id INT PRIMARY KEY,
	majorName VARCHAR(20)
);
DESC stuinfo;
# 查看stuinfo表中所有的索引，包括主键，外键，唯一等；
SHOW INDEX FROM stuinfo;
```



#### 添加表级约束

语法：在各个字段的最下面

[constraint] 约束名 约束类型（字段名）



```mysql
CREATE TABLE stuinfo(
	id INT ,
	stuName VARCHAR(20) ,
	gender CHAR(1),
	seat INT ,
	age INT ,	
	majorId INT,
	
	CONSTRAINT pk PRIMARY KEY(id),#主键
	CONSTRAINT uq UNIQUE(seat),#唯一键 
	CONSTRAINT ck CHECK(gender='男' OR gender='女'), #检查
	CONSTRAINT fk_stuinfo_major FOREIGN KEY(majorId) REFERENCES major(id) # 外键
);
SHOW INDEX FROM stuinfo;
```



通用的写法：

```mysql
CREATE TABLE IF NOT EXISTS stuinfo(
	id INT PRIMARY KEY,
	stuName VARCHAR(20) NOT NULL,
	sex CHAR(1),
	seat INT UNIQUE,
	age INT DEFAULT 18,	
	majorId INT,
	
	CONSTRAINT fk_stuinfo_major FOREIGN KEY(majorId) REFERENCES major(id) # 外键
);
```





### 修改表时添加约束

```mysql
CREATE TABLE stuinfo(
	id INT ,
	stuName VARCHAR(20) ,
	gender CHAR(1),
	seat INT ,
	age INT ,	
	majorId INT
);
```

1.添加列级约束

```mysql
ALTER TABLE 表名 MODIFY COLUMN 字段名 字段类型 新约束;
```

2.添加表级约束

```mysql
ALTER TABLE stuinfo ADD (CONSTRAINT 约束名） 约束类型(字段名)；
```





#### 添加非空约束

```mysql
ALTER TABLE stuinfo MODIFY COLUMN stuName VARCHAR(20) NOT NULL;
```



#### 添加默认约束

```mysql
ALTER TABLE stuinfo MODIFY COLUMN age INT DEFAULT 18;
```



#### 添加主键

```mysql
# 列级约束
ALTER TABLE stuinfo MODIFY COLUMN id INT PRIMARY KEY;

#表级约束
ALTER TABLE stuinfo ADD PRIMARY KEY(id);
```



#### 添加唯一键

```mysql
# 列级约束
ALTER TABLE stuinfo MODIFY COLUMN seat INT UNIQUE;

#表级约束
ALTER TABLE stuinfo ADD UNIQUE(seat);
```



#### 添加外键

```mysql
ALTER TABLE stuinfo ADD FOREIGN KEY(majoriD) REFERENCES major(id);
```

## ⚡ 外键的参照动作（重要！）：

### 1. `ON DELETE` - 当主表数据被删除时

| 动作          | 说明             | 示例                               |
| :------------ | :--------------- | :--------------------------------- |
| `RESTRICT`    | 拒绝删除（默认） | 如果有员工，不能删除部门           |
| `CASCADE`     | 级联删除         | 删除部门时，自动删除该部门所有员工 |
| `SET NULL`    | 设为NULL         | 删除部门时，员工dept_id设为NULL    |
| `NO ACTION`   | 无动作           | 同RESTRICT                         |
| `SET DEFAULT` | 设为默认值       | 删除部门时，dept_id设为默认值      |

### 2. `ON UPDATE` - 当主表主键更新时

| 动作       | 说明     |
| :--------- | :------- |
| `CASCADE`  | 级联更新 |
| `SET NULL` | 设为NULL |
| `RESTRICT` | 拒绝更新 |

示例:

```mysql 
CREATE TABLE employees (
    emp_id INT PRIMARY KEY AUTO_INCREMENT,
    emp_name VARCHAR(50),
    dept_id INT,
    FOREIGN KEY (dept_id) REFERENCES departments(dept_id)
    ON DELETE CASCADE      -- 部门删除时，员工也删除
    ON UPDATE CASCADE      -- 部门ID更新时，员工dept_id同步更新
);
```





#### 修改表时删除约束

1.删除非空约束

```mysql
ALTER TABLE stuinfo MODIFY COLUMN stuName VARCHAR(20) NULL;
```



2.删除默认约束

```mysql 
ALTER TABLE stuinfo MODIFY COLUMN age INT;
```



3.删除主键

```mysql
ALTER TABLE stuinfo DROP PRIMARY KEY;
```



4.删除唯一键

```mysql
ALTER TABLE stuinfo DROP INDEX seat;
```



5.删除外键

```mysql
ALTER TABLE stuinfo DROP FOREIGN KEY majorId;
```





## 标识列

含义：

可以不用手动的插入值，系统提供默认的序列值





### 创建表时设置标识列

特点：

1. 标识列不一定要和主键搭配，但是要求是一个人key

2. 一个表中至多只能有一个标识列

3. 标识列的类型只能是数值型

4. 标识列可以通过`SET auto_increment_increment=3`设置步长

   可以通过手动插入值，来达到设置起始值的目的





```mysql
DROP TABLE IF EXISTS tab_identity;
CREATE TABLE tab_identity(
	id INT PRIMARY KEY auto_increment, #添加标识列
	Name VARCHAR(20)
);


INSERT INTO tab_identity VALUES(NULL,'john');
#设置每次增长的步长值
SET auto_increment_increment=3;
```



### 修改表时设置标识列

```mysql
ALTER TABLE tab_identity MODIFY COLUMN id INT PRIMARY KEY AUTO_INCREMENT;

```



### 修改表时删除标识列

```mysql
ALTER TABLE tab_identity MODIFY COLUMN id INT;
```


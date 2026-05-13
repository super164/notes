# MySQL知识复盘

本文主要是对于学习MySQL的一些知识点的复盘，对有记不清楚以及遗漏知识点的查漏补缺。

## 一.常见概念:

### 1.缩写概念：

DB:数据库（DATABASE），存储数据的仓库

DBMS:数据库管理系统，数据库是通过DBMS创建和操作的

SQL：结构化查询语句，专门用来与数据库通信的语言



DDL:数据定义语言，定义和管理数据库中的结构和模式

DML:数据操作语言，对数据库中的具体数据记录进行增删改查操作

DCL:数据控制语言，控制对数据库的访问权限和安全性

DQL:数据查询语言，顾名思义用来查询数据的语句

TCL:事务控制语言，管理DML操作组成的事务，确保数据的完整性



### 2. 数据库存储数据的特点

1. 将数据放在表中，表在放到库中
2. 一个数据库中可以有多个表，每个表都有一个的名字，用来表示自己。表明具有唯一性。
3. 表具有一些特性，这些特性定义了数据在表中如何存储，类似java中"类"的设计。

4. 表由列组成，我们也成为字段。所有表都是由一个或多个列组成的，每一列类似于java中的"属性"
5. 表中的数据是按行存储的，每一行类似于java中的“对象”。



### 3. 语法规范

1. 不区分带小写，建议关键字大写，表名列名小写
2. 每条命令用分号结尾
3. 每条命令根据需要，可以进行缩进或者换行（建议关键字单独一行）
4. 注释：
   1. 单行注释：#注释文字
   2. 单行注释：-- 注释文字
   3. 多行注释：/* 注释文字 */



下面将根据DDL、DML、DQL、DCL、以及TCL语言的类型进行总结复盘：



## 二.DQL语言--查询语句

### 1. 基础查询

```
select 查询列表 from 表名;
1.查询列表可以是，表中的字段、常量值、表达式、函数
2.查询的结果是一个虚拟的表格
```



### 2. 条件查询

语法： select 查询列表 from 表名 where 筛选条件;

筛选方法：

**一、按照条件表达式筛选**

条件运算符：> 、<、 =、 !=或<> 、>=、<=

**二、按照逻辑表达式筛选**

逻辑运算符：&&、 ||、 ! （and or not）

**三、模糊查询**

like 、between and、 in、is null

1. like一般和通配符配合使用

   通配符:

   % 任意多个字符，包含0个字符

   _ 任意单个字符

案例：

查询员工名中包含字符a 的员工信息

```mysql
SELECT * FROM employees WHERE last_name LIKE 'a'; 
```

2. between and

   包含临界值，并且临界值不要调换顺序

案例：

查询员工编号在100到120之间的员工信息

```mysql
SELECT * FROM employees WHERE employees BETWEEN 100 AND 120;
```

3. in

   判断某字段的值是否属于in列表中的

案例：

查询员工的工种编号是：IT_PROG、AD_VP、AD_PRES中的一个员工名和工种编号

```mysql
SELECT 
	last_name, job_id 
FROM 
	employees 
WHERE 
	job_id IN('IT_PROG','AD_VP','AD_PRES');
```

4. is null

   "="或"!="不能用于判断null值

   is null 或is not null可以判断null的值

   判断某字段或表达式是否为null，如果是，则返回1，否则返回0

案例：

查询没有奖金的员工名和奖金率

``` mysql
SELECT 
	last_name,
	comission_pct 
FROM 
	employees 
WHERE 
	comission_pct IS NULL;
```

5. 安全等于<=>

   可以代替 IS NULL的IS,判断是否等于

### 3. 排序查询

语法：

```
select 查询列表 
from 表 [where 筛选条件] 
order by 排序列表 [asc(升)|desc(降)]

1. 如果不写asc/desc 默认为升序
2. order by子句中可以支持单个字段，多个字段，表达式，函数，别名
3. order by 子句一般放在查询语句的最后面，limit子句除外
```



### 4. 分组查询

GROUP BY 子句语法

```mysql
SELECT column,group_function(column)
FROM table
[WHERE condition]
[GROUP BY group_by_expression]
[ORDER BY column];
```

分组查询顾名思义，具有相同属性的分为一组。

其中对于分组后的结果进行条件筛选可以使用HAVING关键字进行筛选

**简单的分组查询：**

案例1：查询每个工种的最高工资

```mysql
SELECT MAX(salary),job_id
FROM employess
GROUP BY job_id;
```

**按多个字段分组：**

案例1:查询每个部门每个工种的员工的平均工资

```mysql
SELECT AVG(salary),department_id,job_id
FROM employees
GROUP BY department_id,job_id;
```



### 5. 连接查询

连接查询又称多表查询，当查询的字段来自于多个表的时候，就会用到连接查询

一般按照sql99标准：支持内连接+外连接(左外和右外)+交叉连接

语法：

```mysql
select 查询列表
from 表1 别名 [连接条件]
join 表2 别名
on 连接条件
WHERE 筛选条件
GROUP BY 筛选条件
HAVING 筛选条件
ORDER BY 排序列表
```

分类：

内连接:inner

外连接

​	左外：left outer

​	右外：right outer

​	全外：full outer



##### 内连接：

案例1:查询员工名、部门名

```mysql
SELECT last_name,department_name
FROM employees e
INNER JOIN departments d
ON e.department_id=d.department_id;
```

##### 外连接:

应用场景：用于查询一个表中有，另一个表没有的记录

特点：

1. 外连接的查询结果为主表中的所有记

   如果从表中有和它匹配的，则显示匹配的值

   如果从表中没有和它匹配的，则显示null

   外连接查询结果=内连接结果+主表中有而从表中没有的记录

2. 左外连接:left join左边的是主表

3. 右外连接:right join右边的是主表

4. 左外和右外交换两个表的顺序，可以实现同样的效果

**左右连接：**

案例1:查询那个部门没有员工

左：

```mysql
SELECT d.*,e.employee_id 
FROM departments d
LEFT OUTER JOIN employees e
ON d.department_id=e.employee_id
WHERE e.employee_id IS NULL;
```

右:

```mysql
SELECT d.*,e.employee_id 
FROM employees e
RIGHT OUTER JOIN departments d
ON d.department_id=e.employee_id
WHERE e.employee_id IS NULL;
```



### 6.子查询:

含义：

出现在其他语句中的select语句，成为子查询或内查询

内部嵌套其他selec语句t的查询，称为主查询或外查询

分类：

按子查询出现的位置：

​	select后面：只支持标量子查询

​	from后面：支持表子查询

​	where或having后面：支持标量子查询、列子查询、行子查询（较少）

​	exists后面（相关子查询）：表子查询

按案结果集的行列数不同：

​	标量子查询（结果集只有一列一行）

​	列子查询（结果集只有一列多行

​	行子查询（结果集有一行多列）

​	表子查询（结果集一般为多行多列）



#### where或having后面

标量子查询（单行子查询,标量子查询）

列子查询（多行子查询）

行子查询（多列多行）

特点：

1. 子查询放在小括号内
2. 子查询一般放在条件的右侧
3. 表两字查询，一般搭配着单行操作符使用：>< >= = <>
4. 子查询执行优先于主查询执行的，主要查询条件用到了子查询的结果

列子查询，一般搭配着多行操作符使用：IN，ANY/SOME，ALL

##### 标量子查询

案例1：谁的工资比Abel高？

```mysql
#1.查询Abel的工资
SELECT last_name,salary
FROM employees
WHERE last_name='Abel';
#2.查询员工信息，满足salary>1的结果
SELECT *
FROM employees
WHERE salary>(
	SELECT salary
	FROM employees
	WHERE last_name='Abel'
);
```

##### 列子查询：

（多行子查询）

| 操作符          | 功能描述                   | 说明                                                     |
| :-------------- | :------------------------- | :------------------------------------------------------- |
| **IN / NOT IN** | 等于列表中的任意一个       | 用于判断某个值是否在指定的值列表中                       |
| **ANY / SOME**  | 和子查询返回的某一个值比较 | ANY 和 SOME 是同义词，满足子查询结果中的任意一个条件即可 |
| **ALL**         | 和子查询返回的所有值比较   | 必须满足子查询结果中的所有条件                           |

案例1:返回location_id 是1400或1700的部门中的所有员工姓名

```mysql
SELECT last_name
FROM employees
WHERE department_id IN(
	SELECT DISTINCT department_id
	FROM departments
	WHERE location_id=1400
	OR location_id=1700
);
```

#### select后面

案例：查询每个部门的员工个数

```mysql
SELECT d.*,(
	SELECT COUNT(*)
	FROM employees e
	WHERE e.department_id=d.department_id
)
FROM departments d;
```

#### from后面

将子查询的结果充当一张表，要求必须起别名

```mysql
SELECT av_dep.*,g.grade_level
FROM (
	SELECT AVG(salary) av,department_id
	FROM employees
	GROUP BY department_id
) av_dep
INNER JOIN job_grades g
ON av_dep.av BETWEEN lowest_sal AND highest_sal;

```



#### exists后面

相关子查询

语法exists(完整的查询语句)

结果：1或0，是一个布尔类型



案例1：查询有员工的部门名

```mysql
SELECT department_name
FROM departments d
WHERE EXISTS(
	SELECT *
	FROM employees e
	WHERE d.department_id=e.department_id
);
```



## 三. DDL语言--数据定义语言：

数据定义语言

库和表的管理：创建（create），修改(alter)，删除(drop)



### 库的管理：

#### 库的创建：

语法：create database 库名;



案例：创建库Books

```mysql
CREATE DATABASE books;
CREATE DATABASE IF NOT EXISTS books;
```



#### 库的修改

库一般不进行修改

可以更改库的字符集：

```mysql
ALTER DATABASE books CHARACTER SET gbk;
```



#### 库的删除

语法:

```mysql
DROP DATABASE books;
DROP DATABASE IF EXISTS books;
```







### 表的管理

#### 表的创建

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



#### 表的修改

语法：

```mysql
alter table 表名 add|drop|modify|change column 列名 （列类型 约束）
```





##### 修改列名 

```mysql
ALTER TABLE book 
CHANGE COLUMN(可省略) bname bookName VARCHAR(20);

```





##### 修改列的类型或约束

```mysql
ALTER TABLE book MODIFY COLUMN publishdata TIMESTAMP;
ALTER TABLE book MODIFY COLUMN 列名 要修改成的类型;
```



##### 添加新的列

```mysql
ALTER TABLE book ADD COLUMN publishdate TIMESTAMP;
ALTER TABLE book ADD COLUMN 列名 类型;
```



##### 删除列

```mysql
ALTER TABLE book DROP COLUMN A(要删除的列);
```



##### 修改表名

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

   

### 常见的约束：

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

#### 创建表时添加约束

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

##### 添加列级约束

语法：

直接在字段名和类型后面追加，约束类型即可，只支持：默认、非空、主键、唯一

##### 添加表级约束

语法：在各个字段的最下面

[constraint] 约束名 约束类型（字段名）

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

#### 修改表时添加约束

1.添加列级约束

```mysql
ALTER TABLE 表名 MODIFY COLUMN 字段名 字段类型 新约束;
```

2.添加表级约束

```mysql
ALTER TABLE stuinfo ADD (CONSTRAINT 约束名） 约束类型(字段名)；
```

#### 修改表时删除约束

```mysql
1.删除非空约束
ALTER TABLE stuinfo MODIFY COLUMN stuName VARCHAR(20) NULL;
2.删除默认约束
ALTER TABLE stuinfo MODIFY COLUMN age INT;
3.删除主键
ALTER TABLE stuinfo DROP PRIMARY KEY;
4.删除唯一键
ALTER TABLE stuinfo DROP INDEX seat;
5.删除外键
ALTER TABLE stuinfo DROP FOREIGN KEY majorId;
```



## 四. DML语言--数据操作语言

数据操作语言：

插入：insert

修改：update

删除：delete



### 插入语句

#### 方式一：

语法：

insert into 表名(列名,...) values(值1,...);

1. 插入的值的类型要与列的类型一致或兼容

```mysql
INSERT INTO beauty(id,NAME,sex,borndate,phone,photo,boyfriend_id)
VALUES(13,'唐艺昕','女','1990-4-23','18988888888',null,2);
```



2. 不可以为null的列可以为null的列如何插入值？

​	方式一：

```mysql
INSERT INTO beauty(id,NAME,sex,borndate,phone,photo,boyfriend_id)
VALUES(13,'唐艺昕','女','1990-4-23','18988888888',null,2);
```

​	方式二：

```mysql
INSERT INTO beauty(id,NAME,sex,borndate,phone,boyfriend_id)
VALUES(13,'唐艺昕','女','1990-4-23','18988888888',2);
```



3. 列的顺序是可以调换的,但是需要与属性一一对应

4. 列数和值的个数必须一致

5. 可以省略列名，默认所有列，而且列的顺序和表中的顺序一致

   ```mysql
   INSERT INTO beauty
   VALUES(13,'唐艺昕','女','1990-4-23','18988888888',null,2);
   ```

   

#### 方式二：

语法：

insert into 表名

set 列名=值,列名=值,...

```mysql
INSERT INTO beauty 
SET id=14,NAME='刘涛',phone='12134556677';
```



#### 两种方式比较

1. 方式一，支持一次插入多行

   ```mysql
   INSERT INTO beauty(id,NAME,sex,borndate,phone,photo,boyfriend_id)
   VALUES(13,'唐艺昕','女','1990-4-23','18988888888',null,2),
   (13,'唐艺昕','女','1990-4-23','18988888888',null,2),
   (13,'唐艺昕','女','1990-4-23','18988888888',null,2);
   ```

   

2. 方式一：支持子查询，方式二不支持

   ```mysql
   INSERT INTO beauty(id,NAME,phone)
   SELECT 20,'宋茜','13324536545';
   #这样也可以添加信息到表里
   ```

   

### 修改语句

#### 修改单标的记录

语法：

update 表名

set 列=新值,列=新值,...

where 筛选条件



案例1：修改beauty表中姓唐的女神的电话为13123143534

```mysql
UPDATE beauty SET phone = '13123143534'
WHERE NAME LIKE '唐%';
```



#### 修改多表记录

语法：

sql92语法

update 表1 别名，表2 别名

set 列=值,...

where 连接条件

and 筛选条件；



sql99语法：

update 表1 别名

inner|left|right join 表2 别名

on 连接条件

set 列=值,...

where 连接条件；



案例1：修改张无忌女朋友的手机号为114

```mysql
UPDATE boys bo
INNER JOIN beauty b
ON bo.id=b.boyfriend_id
SET b.phone='114'
WHERE bo.boyName='张无忌';
```



### 删除语句

#### 方式一：delect

##### 单表的删除：

语法：

delect from 表名 where 筛选条件



案例1：删除手机号以9结尾的女神信息

```mysql
DELECT FROM beauty WHERE phone LIKE '%9';
```





##### 多表的删除

sql99语法：

```mysql
delect 表1的别名，表2的别名
from 表1 别名
inner|left|right join 表2 别名
on 连接条件
where 筛选条件；
```





案例1：删除张无忌的女朋友的信息

```mysql
DELECT b
FROM beauty b
INNER JOIN boys bo
ON b.boyfriend_id=bo.id
WHERE bo.boyName='张无忌';
```



#### 方式二truncate

语法：truncate table 表名

该语句中不能添加筛选语句

案例：将魅力值>100的男神信息删除

```mysql
TRUNCATE TABLE boys #直接删除该表中全部数据
```



#### delect 和truncate的比较

1. delect 可以家WHERE 条件，TRUNCATE不能加

2. truncate删除，效率高一点

3. 加入要删除的表中有自增长列，

   如果delect删除后，在插入数据，自增长列的值从断点开始，

   而truncate删除后，在插入数据，自增长列的值从1开始

4. truncate删除没有返回值，delect删除有返回值

5. truncate删除不能回滚，delect删除可以回滚



## 五.TCL语言--事务控制语言

### 事务

概念：一个或一组sql语句组成一个执行单元，这个执行单元要么全部执行，要么全部不执行，每个MySQL语句是相互依赖的。

**并不是所有的存储引擎都支持事务，其中有innodb支持事务，而myisan、memory等不支持事务**



#### 实数的ACID属性

1. 原子性：

   原子性是指事务是一个不可分割的工作单位，事务的操作要么都发生，要么都不发生

2. 一致性：

   事务必须是数据库从一个一致性状态变换到另一个一致性状态

3. 隔离性：

   事务的隔离性是指一个事务的执行不能被其他事务干扰，即一个事务内部的操作即使用的数据对并发的其他事务是隔离的，并发执行的各个事务之间不能互相干扰

4. 持久性：

   持久性是一个事务一旦被提交，它对数据库中数据的改变就是永久性的，接下来的其他操作和数据库故障不应该对其有任何影响



#### 事务的创建

##### 隐式事务：

事务没有明显的开启和结束的标记，比如insert、update、delete语句



##### 显式事务

事务具有明显的开启和结束的标记

前提:必须先设置提交功能为禁用

```mysql
SHOW VARIABLES LIKE 'autocommit';#查看状态
SET autocommit=0; # 设置功能为禁用
```



1. 开启事务;

   ```mysql
   SET autocommit=0;
   START TRANSACTION;#可选是否执行
   ```

   

2. 编写事务中sql语句(select、insert、update、delete)

   ```mysql
   语句1;
   语句2;
   ...
   ```

   

3. 结束事务

   ```mysql
   commit;#提交事务
   rollback;#回滚事务
   
   
   savepoint 节点名;# 设置保存点
   ```

   



#### 数据库的隔离级别

对于同时运行的多个事务，当这些事务访问数据库中相同的数据时，如果没有采取必要的隔离机制，就会导致各种并发问题：

1. 脏读：对于两个事务T1,T2,T1读取了已经被T2更新但没有被提交的字段之后，若T2回滚，T1读取的内容就是临时且无效的。
2. 不可重复读：对于两个事务T1,T2,T1读取了一个字段，然后T2更新该字段之后，若T2回滚，T1再次读取同一个字段，值就不同了。

3. 幻读：对于两个事务T1,T2,T1从一个表中读取了一个字段，然后T2在该表中插入了一些新的行，之后，如果T1再次读取同一个表，就会多出几行



##### 数据库的事务隔离级别

| 隔离级别                        | 描述                                                         |
| :------------------------------ | :----------------------------------------------------------- |
| READ UNCOMMITTED (读未提交数据) | 允许事务读取未被其他事务提交的变更。脏读、不可重复读和幻读的问题都会出现。 |
| READ COMMITTED (读已提交数据)   | 只允许事务读取已经被其它事务提交的变更。可以避免脏读，但不可重复读和幻读问题仍然可能出现。 |
| REPEATABLE READ (可重复读)      | 确保事务可以多次从一个字段中读取相同的值。在这个事务持续期间，禁止其他事务对这个字段进行更新。可以避免脏读和不可重复读，但幻读的问题仍然存在。 |
| SERIALIZABLE (串行化)           | 确保事务可以从一个表中读取相同的行。在这个事务持续期间，禁止其他事务对该表执行插入、更新和删除操作。所有并发问题都可以避免，但性能十分低下。 |

MySQL支持4中事务隔离级别，默认的事务隔离级别为REPEATABLE READ



```mysql
#查看隔离级别
select @@transaction_isolation;

#设置当前MySQL连接的隔离级别:
set transaction isolation level read uncommitted;

#设置数据库系统的全局的隔离级别:
set global transaction isolation level raed committed;
```

​            

#### delect和truncate在事务使用时的区别

**演示delect**

```mysql
SET autocommit=0;
START TRANSACTION;
DELECT FROM account;
ROLLBACK;
```



**演示truncate**

```mysql
SET autocommit=0;
START TRANSACTION;
TRUNCATE TABLE account;
ROLLBACK;
```

不支持回滚



#### 演示savepoint的使用

```mysql
SET autocommit=0;
START TRANSACTION;
DELECT FROM account WHERE id =25; 
SAVEPOINT a;#设置保存点
TRUNCATE TABLE account;
ROLLBACK TO a;# 回滚到保存点
```



#### 不支持回滚的操作

##### **1. 通常不支持回滚的操作**

这类操作一旦执行，会立即永久生效，无法使用 `ROLLBACK` 命令撤销。

​	**1. DDL - 数据定义语言**

命令：CREATE, ALTER, DROP, TRUNCATE, RENAME

​	**2. DCL - 数据控制语言**

命令：`GRANT`, `REVOKE`



## 六. 视图：

含义：虚拟表，和普通的表一样使用

一种虚拟存在的表，行和列的数据来定义视图的查询中使用的表，并且在使用视图是动态生成的，只保存了sql逻辑，不保存查询结果。

优势：

1. 重用sql语句
2. 简化复杂的sql操作,不必知道查询细节
3. 保护数据，提高安全性
4. 

### 创建视图

语法：

```mysql
CREATE VIEW 视图名
AS
查询语句;
```



1. 查询邮箱中包含a字符的员工名、部门名和工种信息

   ```mysql
   #创建视图
   CREATE VIEW myv1
   AS
   SELECT last_name,department_id,job_title
   FROM employees e
   JOIN departments d ON e.department_id=d.department_id
   JOIN jobs j ON j.job_id=e.job_id;
   
   #使用视图
   SELECT * FROM myv1 WHERE last_name LIKE '%a%';
   ```

### 修改视图

#### 方式一：

```mysql
#若视图存在则进行替代，拖视图不存在则进行创造
CREATE OR REPLACE vIEW 视图名
AS
查询语句;

CREATE OR REPLACE VIEW myv3
AS 
SELECT AVG(salary),job_id
FROM employees
GROUP BY job_id;
```



#### 方式二

```mysql
ALTER VIEW 视图名
as
查询语句;
```



### 删除视图

语法：DROP VIEW 视图名,视图名,...;

```mysql
DROP VIEW myv1,myv2,myv3;
```



### 查看视图

```mysql
# 查看视图创建的过程

SHOW CREATE VIEW myv3;
```





### 视图更新

#### 可修改的

```mysql
CREATE OR REPLACE VIEW myv1
AS 
SELECT last_name,email
FROM employees;
```

1. 插入

   ```mysql
   INSERT INTO myv1 VALUES('张飞','zf@qq.com');
   ```

   上述语句在表格中插入新的信息，视图中会进行更改，在原始的表中也会进行更改

2. 修改

   ```mysql
   UPDATE myv1 SET last_name='张无忌' WHERE last_name='张飞';
   ```

3. 删除

   ```mysql
   DELETE FROM myv1 WHERE last_name='张无忌';
   ```

   

#### 不可修改的

具备以下特点的视图不允许更新:

1. 包含以下关键字的sql语句，分组函数、distinct、group by、having、union、或者union ail
2. 常量视图
3. SELECT中包含子查询
4. from一个不能更新的视图
5. WHERE子句的子查询引用了from子句中的表





### 视图和表的对比

|      | 创建语法的关键字 | 是否实际占用物理空间 | 使用                     |
| :--- | :--------------- | :------------------- | :----------------------- |
| 视图 | create view      | 只是保存了 SQL 逻辑  | 增删改查，一般不能增删改 |
| 表   | create table     | 保存了数据           | 增删改查                 |



## 七. DCL--数据控制语言:

管理数据库系统的**安全性**和**权限控制**。

#### **主要命令（只有两个，但功能强大）**

##### **1. GRANT**

- **功能**：**授予**权限给用户或角色。

- **语法**：

  ```mysql
  GRANT 权限列表 ON 数据库对象 TO 用户或角色 [WITH GRANT OPTION];
  ```

##### **2. REVOKE**

- **功能**：**收回**之前授予用户或角色的权限。

- **语法**：

  ```mysql
  REVOKE 权限列表 ON 数据库对象 FROM 用户或角色;
  ```



## 八.查漏补缺

### 1. 创建索引

```mysql
-- 创建普通索引
CREATE INDEX idx_customer_name ON customers (last_name, first_name);

-- 创建唯一索引
CREATE UNIQUE INDEX idx_customer_email ON customers (email);

-- 创建全文索引（用于文本搜索）
CREATE FULLTEXT INDEX idx_product_description ON products (description);
```



### 2. 表引擎：

常见的几个表引擎有:InnoDB(一般默认)，MyISAM(过时了),(Memory)(内存引擎),Archive(归档引擎)

#### **如何查看和设置表引擎**

**1. 查看某个表的引擎**

```mysql
SHOW TABLE STATUS LIKE 'table_name';
```

或者

```mysql
-- 更清晰的方式
SELECT TABLE_NAME, ENGINE 
FROM information_schema.TABLES 
WHERE TABLE_SCHEMA = 'your_database_name';
```

**2. 创建表时指定引擎**

```mysql
CREATE TABLE my_table (
    id INT PRIMARY KEY,
    data VARCHAR(100)
) ENGINE=InnoDB; -- 这里指定引擎
```

**3. 修改已有表的引擎**

```mysql
ALTER TABLE my_table ENGINE = InnoDB;
```
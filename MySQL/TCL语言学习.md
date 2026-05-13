# MySQL数据库---事务控制语言



## 事务

概念：一个或一组sql语句组成一个执行单元，这个执行单元要么全部执行，要么全部不执行，每个MySQL语句是相互依赖的。

**并不是所有的存储引擎都支持事务，其中有innodb支持事务，而myisan、memory等不支持事务**



### 事务的ACID属性

1. 原子性：

   原子性是指事务是一个不可分割的工作单位，事务的操作要么都发生，要么都不发生

2. 一致性：

   事务必须是数据库从一个一致性状态变换到另一个一致性状态

3. 隔离性：

   事务的隔离性是指一个事务的执行不能被其他事务干扰，即一个事务内部的操作即使用的数据对并发的其他事务是隔离的，并发执行的各个事务之间不能互相干扰

4. 持久性：

   持久性是一个事务一旦被提交，它对数据库中数据的改变就是永久性的，接下来的其他操作和数据库故障不应该对其有任何影响



### 事务的创建

#### 隐式事务：

事务没有明显的开启和结束的标记，比如insert、update、delete语句



#### 显式事务

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

   

演示:

```mysql
DROP TABLE IF EXISTS account;

CREATE TABLE account(
	id INT PRIMARY KEY AUTO_INCREMENT,
	userName VARCHAR(20),
	balance DOUBLE
);

INSERT INTO account(username,balance)
VALUES('张无忌',1000),('赵敏',1000);

SET autocommit=0;
START TRANSACTION;

UPDATE account SET balance=1000 WHERE username='张无忌';
UPDATE account SET balance=1000 WHERE username='赵敏';

ROLLBACK;# 回滚
COMMIT;# 提交
```





### 数据库的隔离级别

对于同时运行的多个事务，当这些事务访问数据库中相同的数据时，如果没有采取必要的隔离机制，就会导致各种并发问题：

1. 脏读：对于两个事务T1,T2,T1读取了已经被T2更新但没有被提交的字段之后，若T2回滚，T1读取的内容就是临时且无效的。
2. 不可重复读：对于两个事务T1,T2,T1读取了一个字段，然后T2更新该字段之后，若T2回滚，T1再次读取同一个字段，值就不同了。

3. 幻读：对于两个事务T1,T2,T1从一个表中读取了一个字段，然后T2在该表中插入了一些新的行，之后，如果T1再次读取同一个表，就会多出几行



#### 数据库的事务隔离级别

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

### delect和truncate在事务使用时的区别

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



### 演示savepoint的使用

```mysql
SET autocommit=0;
START TRANSACTION;
DELECT FROM account WHERE id =25; 
SAVEPOINT a;#设置保存点
TRUNCATE TABLE account;
ROLLBACK TO a;# 回滚到保存点
```





## 视图

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

2. 查询各部门的平均工资级别

   ```mysql
   CREATE VIEW myv2
   AS
   SELECT AVG(salary) ag,department_id
   FROM employees
   GROUP BY department_id;
   
   SELECT myv2.ag,g.grade_level
   FROM myv2
   JOIN job_grades g
   ON myv2.ag BETWEEN g.lowest_sal AND g.highest_sal
   ORDER BY ag DESC;
   ```

3. 查询平均工资最低的部门信息

   ```mysql
   CREATE VIEW myv3
   AS
   SELECT *
   FROM myv2
   ORDER BY ag LIMIT 1;
   ```

4. 查询平均工资最低的部门名和工资

   ```mysql
   SELECT d.*,ag
   FROM myv3
   JOIN departments d 
   ON myv3.department_id=d.department_id;
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
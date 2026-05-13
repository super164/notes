# Mysql数据库---数据操作语言

数据操作语言：

插入：insert

修改：update

删除：delete



## 插入语句

### 方式一：

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

   

### 方式二：

语法：

insert into 表名

set 列名=值,列名=值,...



```mysql
INSERT INTO beauty 
SET id=14,NAME='刘涛',phone='12134556677';
```



### 两种方式比较

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

   

## 修改语句

### 修改单标的记录

语法：

update 表名

set 列=新值,列=新值,...

where 筛选条件



案例1：修改beauty表中姓唐的女神的电话为13123143534

```mysql
UPDATE beauty SET phone = '13123143534'
WHERE NAME LIKE '唐%';
```



案例2：修改boys表中的id号为2的名称为张飞，魅力值为10

```mysql
UPDATE boys SET boyName = '张飞',userCP=10
WHERE id =2;
```





### 修改多表记录

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



案例2：修改没有男朋友的女神的男朋友编号都为2号

```mysql
UPDATE boys bo
RIGHT JOIN beauty b
ON bo.id=b.boyfriend_id
SET b.boyfriend_id=2
WHERE bo.id IS NULL;
```





## 删除语句

### 方式一：delect

#### 单表的删除：

语法：

delect from 表名 where 筛选条件



案例1：删除手机号以9结尾的女神信息

```mysql
DELECT FROM beauty WHERE phone LIKE '%9';
```





#### 多表的删除

sql92语法：

```mysql
delect 表1的别名，表2的别名
from 表1 别名，表2，别名
where 连接条件
and 筛选条件;
```

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



案例2：删除黄晓明的信息以及他的女朋友的信息

```mysql
DELECT b,bo
FROM beauty b
INNER JOIN boys bo
ON b.boyfriend_id=bo.id
WHERE bo.boyName='黄晓明';
```





### 方式二：truncate

语法：truncate table 表名

该语句中不能添加筛选语句

案例：将魅力值>100的男神信息删除

```mysql
TRUNCATE TABLE boys #直接删除该表中全部数据
```





### delect 和truncate的比较

1. delect 可以家WHERE 条件，TRUNCATE不能加

2. truncate删除，效率高一点

3. 加入要删除的表中有自增长列，

   如果delect删除后，在插入数据，自增长列的值从断点开始，

   而truncate删除后，在插入数据，自增长列的值从1开始

4. truncate删除没有返回值，delect删除有返回值

5. truncate删除不能回滚，delect删除可以回滚
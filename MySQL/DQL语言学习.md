# MySQl基础学习--查询语句

## 基础查询

```
select 查询列表 from 表名;
1.查询列表可以是，表中的字段、常量值、表达式、函数
2.查询的结果是一个虚拟的表格
```

1. **查询表中的单个字段:**

​	例如查询employees表中的last_name这个字段：

​	SELECT last_name FROM employees;

2. **查询多个字段:**

   SELECT last_name,salary,email FROM employees;

3. **查询表中的所有字段：**

   SELECT * FROM employees; 

4. **查询常量值：**

   SELECT 100;

   SELECT 'john';

5. **查询表达式：**

   SELECT 100%98;

6. **查询函数:**

   SELECT VERSION();

7. **起别名：**

   便于理解的同时，如果要查询的字段有重名的情况，使用别名可以区分开来

   方式一：

   SELECT 100%98 AS 结果;

   AS后面跟的就是要起的别名

   方式二：

   SELECT last_name 姓，first_name 名 FROM employees;

   直接写在字段后面，即为其别名

8. **去重：（DISTINCT）**

   SELECT DISTINCT department_id FROM employees;

   distinct

9. **"+"号的作用：**

   ```
   仅作为运算符，不能作为连接符
   
   1. 两个数都为数值型，则做加法运算
   
   2. 若一方为字符型，则试图将字符型数值转换成数值型，
   
   转换成功则做加法，转换失败则将字符型数值转换成0
   
   3. 只要是一方为null,则结果必定为null
   ```

   

10. **连接符(concat)**

    SELECT CONCAT('a','b'); 将列表内容拼接起来输出为"ab"

    利用IFNULL(字段名，0)判断是否存在空值

11. **if null函数**

    判断某字段或表达式是否为null,如果为Null,返回指定的值，否则返回原本的值

    ```mysql
    select ifnull(commission_pct,0) from employees
    ```

    

## 条件查询

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

查询员工名中的第二个字符为_的员工名

```mysql
SELECT * FROM employees WHERE last_name LIKE '_\_%'; 
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



## 排序查询

语法：

```
select 查询列表 
from 表 [where 筛选条件] 
order by 排序列表 [asc(升)|desc(降)]

1. 如果不写asc/desc 默认为升序
2. order by子句中可以支持单个字段，多个字段，表达式，函数，别名
3. order by 子句一般放在查询语句的最后面，limit子句除外
```



案例1：查询员工信息，要求工资从高到低排序

```mysql
SELECT * FROM employees ORDER BY desc;
```

案例2:查询部门编号>=90的员工信息，按入职时间的先后排序

```mysql
select * 
from employees 
WHERE department_id>=90 order by hiredate asc;
```

案例3：查询员工信息和年薪，要求年薪的高低排序(按表达式进行排序)

```mysql
SELECT *,salary*12*(1+IFNULL(commission_pct,0)) 年薪
FROM employees
ORDER BY salary*12*(1+IFNULL(commission_pct,0)) DESC;
# order by 后面支持填写别名
```

案例4:按姓名长度显示员工的姓名和工资（按函数进行排序）

```mysql
SELECT LENGTH(last_name) 字符长度,last_name,salary
FROM employees
ORDER BY 字符长度 desc;
```

案例5:查询员工信息，要求先按工资升序，再按员工编号降序（按多个字段排序）

```mysql
SELECT * FROM employees 
ORDER BY salary asc,employee_id desc;
```



## 分组查询

GROUP BY 子句语法

```mysql
SELECT column,group_function(column)
FROM table
[WHERE condition]
[GROUP BY group_by_expression]
[ORDER BY column];
```

**简单的分组查询：**

案例1：查询每个工种的最高工资

```mysql
SELECT MAX(salary),job_id
FROM employess
GROUP BY job_id;
```

案例2：查询每个位置上的部门个数

```mysql
SELECT COUNT(*),location_id
FROM departments 
GROUP BY location_id;
```



**添加筛选条件：**

案例1:查询邮箱中包含a字符的，每个部门的平均工资

```mysql
SELECT AVG(salary) 平均工资,department_id
FROM employees
WHERE email LIKE '%a%'
GROUP BY department_id;
```



案例2:

查询有奖进的每个领导手下员工的最高工资

```mysql
SELECT MAX(salary),manager_id
FROM employees
WHERE commission_pct IS NOT NULL
GROUP BY manager_id;
```



**添加复杂的筛选条件：**

案例1：查询哪个部门的员工个数大于二

```mysql
SELECT COUNT(*),department_id
FROM employees
GROUP BY department_id
HAVING COUNT(*)>2;
```

使用了一个HAVING这个关键字，用户分组后的每一行进行筛选的



案例2：查询每个公众版有奖金的员工的最高工资>12000的工种编号和最高工资

```mysql
SELECT MAX(salary),job_id
FROM employees
WHERE commission_pct IS NOT NULL
GROUP BY job_id
HAVING MAX(salary)>12000;
```



案例3:查询领导编号>102的每个领导手下最低工资>5000的领导编号

```mysql
SELECT MIN(salary),manager_id
FROM employees
WHERE manager_id>102
GROUP BY manager_id
HAVING MIN(salary)>5000;
```



**按表达式或函数分组：**

案例1：按员工姓名的长度分组，查询每一组的员工个数，筛选员工个数>5的有哪些

```mysql
SELECT COUNT(*),LENGTH(last_name) 姓名长度
FROM employees
GROUP BY LENGTH(last_name)
HAVING COUNT(*)>5;
```



**按多个字段分组：**

案例1:查询每个部门每个工种的员工的平均工资

```mysql
SELECT AVG(salary),department_id,job_id
FROM employees
GROUP BY department_id,job_id;
```



**添加排序：**

案例1：查询每个部门每个工种的员工的平均工资，并且按平均工资的高低显示

```mysql
SELECT AVG(salary),department_id,job_id
FROM employees
GROUP BY department_id,job_id
ORDER BY AVG(salary) desc;
```



## 连接查询

连接查询又称多表查询，当查询的字段来自于多个表的时候，就会用到连接查询

语法：select name,boyName from beauty,boys;

这样写会造成笛卡尔集错误的情况，没有有效的连接条件

需加上有效的连接条件进行连接查询23

```mysql
select name,boyName from beauty,boys
WHERE beauty.boyfriend_id=boys.id;
```

一般按照sql99标准：支持内连接+外连接(左外和右外)+交叉连接



### **按标准分类（sql92）：**

**内连接：**

等值连接、非等值链接、自连接

**外连接：**

左外连接、右外连接、全外连接

**交叉连接：**



#### 等值连接：

案例1:查询女的名字和对应的男的名字

```mysql
SELECT name,boyName
FROM boys,beauty
WHERE beauty.boyfriend_id=boys.id;
```



案例2:查询员工名、工种号、工种名

```mysql
SELECT last_name,employees.job_id,job_title
FROM employees,jobs
WHERE employees.job_id=jobs.job_id;
```



##### 加筛选的：

案例1：查询有奖金的员工名、部门名

```mysql
SELECT last_name,department_name
FROM employees,departments
WHERE employees.department_id=departments.department_id
AND salary IS NOT NULL;
```



案例2:查询城市名中第二个字符为o的部门

```mysql
SELECT department_name,city
FROM departments d,locations l
WHERE d.location_id=l.location_id
AND city LIKE '_o%';
```



##### 加分组的：

案例1：查询每个城市的部门个数

```mysql
SELECT COUNT(*) 个数,city
FROM departments d,locations l
WHERE d.location_id=l.location_id
GROUP BY city;
```



案例2：查询有奖金的每个部门的部门名和部门的领导编号和该部门的最低工资

```mysql
SELECT department_name,d.manager_id,MIN(salary)
FROM departments d,employees e
WHERE d.department_id=e.department_id
AND commission_pct IS NOT NULL
GROUP BY department_name,manager_id;
```



##### 加排序的：

案例1：查询每个工种的工种名和员工的个数，并且按员工的个数降序

```mysql
SELECT COUNT(*),job_title
FROM employees e,jobs j
WHERE e.job_id=j.job_id
GROUP BY job_title
ORDER BY COUNT(*) desc;
```



##### 实现三表连接：

案例1：查询员工名、部门名和所在的城市

```mysql
SELECT last_name,department_name,city
FROM employees e,departments d,locations l
WHERE e.department_id=d.department_id
AND l.location_id=d.location_id;
```



#### 非等值连接

案例1:查询员工的工资和工资级别

```mysql
SELECT salary,grade_level
FROM employees e,job_grades j
WHERE salary BETWEEN j.lowest_sal
AND j.highest_sal;
```



#### 自连接

案例1:查询员工名和上级的名称

```mysql
SELECT e.last_name,e.employee_id,m.last_name,m.employee_id
FROM employees e,employees m
WHERE m.employee_id=e.manager_id;
```



### 按标准分类（sql99）：

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

交叉连接:cross



#### 等值连接：

案例1:查询员工名、部门名

```mysql
SELECT last_name,department_name
FROM employees e
INNER JOIN departments d
ON e.department_id=d.department_id;
```



案例2:查询名字中包含e 的员工名和工种名

```mysql
SELECT last_name,job_title
FROM employees e
INNER JOIN jobs j
ON e.job_id=j.job_id
WHERE last_name LIKE '%e%';
```



案例3：查询部门个数>3的城市名和部门个数

```mysql
SELECT city,COUNT(*) 部门个数
FROM departments d
INNER JOIN locations l
ON d.location_id=l.location_id
GROUP BY city
HAVING COUNT(*)>3;

```



案例4：查询哪个部门的部门员工个数>3的部门名和员工个数，并按个数排序

```mysql
SELECT department_name,COUNT(*) 员工个数
FROM employees e
INNER JOIN departments d
ON e.department_id=d.department_id
GROUP BY department_name
HAVING COUNT(*)>3
ORDER BY COUNT(*) desc;
```



案例5:查询员工名、部门名、工种名，并按部门名降序

```mysql
SELECT last_name,department_name,job_title
FROM employees e
INNER JOIN departments d ON e.department_id =d.department_id
INNER JOIN jobs j ON e.job_id= j.job_id
ORDER BY department_name desc;
```



#### 非等值连接：

案例1:查询员工的工资级别

```mysql
SELECT grade_level,salary
FROM employees e
INNER JOIN job_grades j
ON e.salary BETWEEN j.lowest_sal AND j.highest_sal;
```



案例2：查询工资级别的个数>20的个数，并且按工资级别降序

```mysql
SELECT COUNT(*),grade_level
FROM employees e
INNER JOIN job_grades j
ON e.salary BETWEEN j.lowest_sal AND j.highest_sal
GROUP BY grade_level
HAVING COUNT(*)>20
ORDER BY grade_level desc;
```



#### 自连接：

案例1:查询员工的名字、上级的名字

```mysql
SELECT e.last_name,m.last_name
FROM employees e
INNER JOIN employees m
ON e.manager_id=m.employee_id;
```



外连接

应用场景：用于查询一个表中有，另一个表没有的记录

特点：

1. 外连接的查询结果为主表中的所有记

   如果从表中有和它匹配的，则显示匹配的值

   如果从表中没有和它匹配的，则显示null

   外连接查询结果=内连接结果+主表中有而从表中没有的记录

2. 左外连接:left join左边的是主表

3. 右外连接:right join右边的是主表

4. 左外和右外交换两个表的顺序，可以实现同样的效果

案例：查询男朋友不在男神表的女神名

```mysql
SELECT b.name
FROM beauty b
LEFT JOIN boys bo
ON b.boyfriend_id=bo.id
WHERE bo.id IS NULL;
```



#### 左右连接：

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





## 子查询：

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



### where或having后面

标量子查询（单行子查询,标量子查询）

列子查询（多行子查询）

行子查询（多列多行）



特点：

1. 子查询放在小括号内
2. 子查询一般放在条件的右侧
3. 表两字查询，一般搭配着单行操作符使用：>< >= = <>
4. 子查询执行优先于主查询执行的，主要查询条件用到了子查询的结果

列子查询，一般搭配着多行操作符使用：IN，ANY/SOME，ALL



#### 标量子查询

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



案例2：返回job_id与141员工相同，salary比员工多的员工 姓名，job_id和工资

```mysql
#1.查询141号员工的job_id
SELECT job_id
FROM employees
WHERE employee_id=141;

#2.查询143员工的salary
SELECT salary
FROM employees
WHERE employee_id=143;

#3.查询员工的姓名，job_id和工资，要求job_id=1的结果并salary>2的结果
SELECT last_name,job_id,salary
FROM employees
WHERE job_id=(
	SELECT job_id
	FROM employees
	WHERE employee_id=141
)AND salary>(
	SELECT salary
	FROM employees
	WHERE employee_id=143
);
```



案例3:返回公司工资最少的员工的last_name,job_id和salary

```mysql
#1.查询公司最低工资
SELECT MIN(salary)
FROM employees;

#2.查询last_name,job_id和salary,要求salary=1的结果
SELECT last_name,job_id,salary
FROM employees
WHERE salary = (
	SELECT MIN(salary)
	FROM employees
);
```



案例4:查询最低的工资大于50号部门最低工资的部门id和其最低工资

```mysql
SELECT department_id,MIN(salary)
FROM employees
GROUP BY department_id
HAVING MIN(salary) > (
	SELECT MIN(salary)
	FROM employees
	WHERE department_id=50
);
```





#### 列子查询：

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

案例2：返回其他工种中比job_id为IT_PROG工种任一工资低的员工的：工号、姓名、job_id、以及salary

```mysql
SELECT employee_id,last_name,job_id,salary
FROM employees
WHERE salary < ANY(
	SELECT salary
	FROM employees
	WHERE job_id='IT_PROG'
) AND job_id <> 'IT_PROG';
```



#### 行子查询

（结果集一行多列或多行多列）

案例：查询员工编号最小并且工资最高的员工信息

```mysql
SELECT *
FROM employees
WHERE employee_id=(
	SELECT MIN(employee_id)
	FROM employees
) AND (
	SELECT MAX(salary)
	FROM employees
);

# 行子查询：
SELECT *
FROM employees
WHERE (employee_id,salary)=(
	SELECT MIN(employee_id),MAX(salary)
	FROM employees
);
```

### select后面

案例：查询每个部门的员工个数

```mysql
SELECT d.*,(
	SELECT COUNT(*)
	FROM employees e
	WHERE e.department_id=d.department_id
)
FROM departments d;
```





### from后面

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



### exists后面

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



案例2：查询没有女朋友的男神信息

```mysql
SELECT bo.*
FROM boys bo
WHERE NOT EXISTS(
	SELECT *
	FROM beauty b
	WHERE bo.id=b.boyfriend_id
);
```



## 分页查询

应用场景：当同要显示的数据，一页显示不全，需要分页提交sql请求



语法：

```mysql
select 查询列表
from 表
筛选条件
limit offset,size;

offset要显示条目的起始索引（起始索引从0开始）
size显示条目个数
```

特点：

1. limit语句放在查询语句的最后，也是最后执行

2. 公式：

   要显示的页数page,每条的条数目size

   ```mysql
   select 查询列表
   from 表
   limit (page-1)*size,size
   ```

   

案例1：查询前五条员工信息

```mysql
SELECT *
FROM employees
LIMIT 0,5;
```



案例2：有奖金的员工信息，并且工资较高的前10名显示出来

```mysql
SELECT *
FROM employees
WHERE commission_pct IS NOT NULL
ORDER BY salary desc
LIMIT 10;
```



## 联合查询

union 联合 合并：将多条查询语句的结果合并成一个结果

引入：查询部门编号>90或邮箱包含a的员工信息

```mysql
SELECT * FROM employees WHERE email LIKE '%a%' OR department_id>90;

# 联合查询使用
SELECT * 
FROM employees 
WHERE email LIKE '%a%'
UNION
SELECT * 
FROM employees 
WHERE department_id>90;
```



语法： 

查询语句1

UNION

查询语句2

UNION

...

应用场景：

当要查询的内容涉及多个表，并且没有直接的连接关系，但是查询的信息一致时，可以使用联合查询



特点：

1. 要求查询语句的查询列表是一致的
2. 要求多条查询语句的查询的每一列的类型和顺序最好是一致的
3. union关键字默认是去重的，如果使用union all可以包含重复项




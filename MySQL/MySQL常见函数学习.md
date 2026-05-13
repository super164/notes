# MySQl基础学习--常见函数

调用:select 函数名(实参列表) [from 表]

函数里面调用了表中的字段则加上from表，没用的话则不用from表

函数名、函数功能

分类：

## 单行函数：

concat 、length、ifnull等

### 字符函数：

1. length:获取参数值的字节

   ```mysql
   SELECT LENGTH('john');  #4
   SELECT LENGTH('张三丰hahaha'); # 15
   ```

2. concat:拼接字符串

   ```mysql
   SELECT CONCAT(last_name,'_',frist_name) 姓名 FROM employees;
   ```

   

3. upper、lower:

   ```mysql
   SELECT UPPER('john'); # 变成大写
   SELECT LOWER('JOhN'); # 变成小写
   
   # 姓名变大写，名变小写，然后拼接
   SELECT CONCAT(UPPER(last_name,LOWER(first_name))) 姓名 
   FROM employees;
   ```

   

4. sbustr、substring

   截取字符，索引从1开始

   ```mysql
   # 截取从指定索引处后面所有字符
   SELECT SUBSTR('李莫愁爱上陆展元',7) out_put;
   
   #两个参数是截取从指定索引处，指定长度的字符
   SELECT SUBSTR('李莫愁爱上陆展元',1,,3) out_put; 
   
   #姓名中首字符大写，其他字符小写然后用_拼接，显示出来
   SELECT CONCAT(UPPER(SUBSTR(last_name,1,1)),'_',LOWER(SUBSTR(last_name,2))) 姓名 FROM employees;
   ```

5. instr:

   返回子串第一次出现的索引，如果找不到则返回0

   ```mysql
   SELECT INSTR('杨不悔爱上了殷六侠','殷六侠') out_put
   ```

   

6. trim:

   去掉前后空格或者指定字符

   ```mysql
   #去掉前后空格
   SELECT TRIM('     张翠山     ') out_put;
   
   #去掉前后指定的字符
   SELECT TRIM('a' FROM 'aaaaaaa张aaaaaaaa翠山aaaaaaaaaaaa') out_put
   
   ```

   

7. lpad:用指定的字符实现左填充指定长度

   ```mysql
   #如果长度参数小于左侧的字段，则会进行截断从左往右截断
   SELECT LPAD('殷素素',10,'*') out_put
   
   ```

   

8. rpad:用指定的字符实现右填充指定长度

   ```mysql
   #如果长度参数小于左侧的字段，则会进行截断从右往左截断
   SELECT RPAD('殷素素',10,'*') out_put
   
   ```

   

9. replace:替换

   ```mysql
   SELECT REPLACE('张无忌爱上了周芷若','周芷若','赵敏') out_put;
   ```

   

   

### 数学函数：

1. round:四舍五入

   ```mysql
   SELECT ROUND(1,65) out_put;
   
   SELECT ROUND(1.567,2) out_put; #输出1.57,小数点保留两位
   ```

   

2. ceil：向上取整，返回>=该参数的最小整数

   ```mysql
   SELECT CEIL(1.01); # 返回的是2
   ```

   

3. floor：向下取整，返回<=该参数的最大整数

   ```mysql
   SELECT FLOOR(9.99); #返回的是9
   ```

   

4. truncate:截断

   ```mysql
   # 不管后面有多少，直接截断到需要的长度
   SELECT TRUNCATE(1.699999,1);
   ```

   

5. mod:取模

   ```mysql
   SELECT MOD(10,3);
   ```

   

### 日期函数：

1. now:返回当前系统日期+时间

   ```mysql
   SELECT NOW();
   ```

   

2. curdate:返回当前系统日期，不包含时间

   ```mysql
   SELECT CURDATE();
   ```

   

3. curtime:返回当前时间，不包含日期

   ```mysql
   SELECT CURTIME();
   ```

   

4. 获取指定的部分，年、月、日、小时、分钟、秒

   ```mysql
   SELECT YEAR(NOW()) 年;
   SELECT YEAR('1998-1-1') 年;
   
   SELECT YEAR(hiredate) 年 FROM employees;
   
   SELECT MONTH(NOW()) 月; # 输出数字
   SELECT MONTHNAME(NOW()) 月； # 输出的是英文
   ```

   

5. str_to_date:将日期格式的字符转换成指定格式的日期

   ```mysql
   SELECT STR_TO_DATE('9-13-1999','%m-%d-%Y'); 1999-09-13
   ```

   案例:

   查询入职日期为1992-4-3的员工信息

   ```mysql
   SELECT * 
   FROM employees 
   WHERE hiredate = STR_TO_DATE('4-3-1992','%m-%d-%Y');
   ```

   

6. date_format:将日期转换成字符

   ```mysql
   SELECT DATE_FORMAT('2018/6/6','%Y年%m月%d日'); 2018年06月06日
   ```

案例：查询有奖金的员工名和入职日期(xx月/xx日 xx年)

```mysql
SELECT last_name,DATE_FORMAT(hiredate,'%m月%d日%Y年') 入职日期
FROM employees
WHERE commission_pct IS NOT NULL;
```



7. DATEDIFF

   ```mysql
   SELECT DATEDIFF(日期1,日期2);
   ```

   计算两个日期之间的差值

#### 格式符：

| 格式符 | 功能描述              |
| :----- | :-------------------- |
| %Y     | 四位的年份            |
| %y     | 2位的年份             |
| %m     | 月份（01,02...11,12） |
| %c     | 月份（1,2,...11,12）  |
| %d     | 日（01,02,...）       |
| %H     | 小时（24小时制）      |
| %h     | 小时（12小时制）      |
| %i     | 分钟（00,01...59）    |
| %s     | 秒（00,01,...59）     |

### 其他函数：

```mysql
#查看版本号
SELECT VERSION();
# 查看当前的库
SELECT DATABASE();
#当前的用户
SELECT USER();
```



### 流程控制函数：

1. if函数:if else 的效果

   ```mysql
   SELECT IF(10>5,'大','小');
   
   SELECT last_name,commission_pct,IF(commission_pct IS NULL,'没奖金','有奖金') 有无奖金
   FROM employees;
   ```

   

2. case函数的使用一:switch case的效果

   ```mysql
   case 要判断的字段或表达式
   when 常量1 then 要显示的值1或语句1;
   when 常量2 then 要显示的值2或语句2;
   when 常量3 then 要显示的值3或语句3;
   ···
   else 要显示的值n或语句n
   end
   ```

   案例：查询员工的工资，要求

   部门号=30,显示的工资为1.1倍

   部门号=40,显示的工资为1.2倍

   部门号=50,显示的工资为1.3倍

   其他部门显示工资为原工资

   ```mysql
   SELECT salary 原始工资,department_id,
   CASE department_id
   	WHEN 30 THEN
   		salary*1.1
   	WHEN 40 THEN
   		salary*1.2
   	WHEN 50 THEN
   		salary*1.3
   	ELSE
   		salary
   END 新工资
   FROM employees;
   ```

   

3. case函数的使用一:多重if的效果

   ```mysql
   case
   when 条件1 then 要显示的值1或语句1
   when 条件2 then 要显示的值2或语句2
   ···
   else 要显示的值n或语句n
   end
   ```

   案例：查询员工的工资情况

   如果工资>20000,显示A级别

   如果工资>15000,显示B级别

   如果工资>10000,显示C级别

   否则,显示D级别

   ```mysql
   SELECT last_name,salary,
   CASE 
   	WHEN salary>20000 THEN
   		'A'
   	WHEN salary>15000 THEN
   		'B'
   	WHEN salary>10000 THEN
   		'C'
   	ELSE
   		'D'
   END 工资级别
   FROM employees;
   ```

   

## 分组函数:

**功能：**做统计使用，又称统计函数、聚合函数、组函数

**分类:**

sum 求和，avg 平均值，max 最大值，min 最小值，count 计算个数



**特点：**

1. SUM和AVG,一般适合处理数值型的数

   MAX，MIN，COUNT什么类型的都适合

   COUNT,计算不为空的个数

3. 上面几个函数均忽略null值

   

4. 可以和distinct搭配实现去重的运算

   



### 1.简单的使用

```mysql
SELECT SUM(salary) FROM employees;
SELECT AVG(salary) FROM employees;
SELECT MIN(salary) FROM employees;
SELECT MAX(salary) FROM employees;
SELECT COUNT(salary) FROM employees;
```



### 2.count函数的单独介绍：

```mysql

SELECT COUNT(salary) FROM employees;

#统计个数
SELECT COUNT(*) FROM employees; //一般使用这个
SELECT COUNT(1) FROM employees;
```


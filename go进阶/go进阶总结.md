# Golang进阶总结

## 目录

[1.方法](#1.方法)

[2.面向对象](#2.面向对象)

[3.接口interface](#3.接口interface)

[4.单元测试](#4.单元测试)

[5.文件操作](#5.文件操作)

[6.JSON](#6.JSON)

[7.goroutine协程](#7.goroutine协程)

[8.channel管道](#8.channel管道)

[9.反射](#9.反射)

[10.context](#10.context)

[11.网络编程](#11.网络编程)

## 1.方法

[目录](#目录)	

​	方法可以实现结构体一些行为，Golang中方法是作用在指定的数据类型上的，因此自定义类型，都可以有方法  

方法的定义：

```go
func (recevier type) methodName(参数列表)(返回值列表){
    方法体
    return 返回值
}
```

1) 参数列表:表示方法输入 
2) recevier type :表示这个方法和type这个类型进行绑定，或者说该方法作用于type类型
3) receiver type:type可以是结构体，也可以其它的自定义类型
4) receiver :就是type类型的一个变量(实例)，比如: Person结构体的一个变量(实例) 
5) 返回值列表:表示返回的值,可以多个
6) 方法主体：表示为了实现某一功能代码块
7) return语句不是必须的。

方法声明和调用：

```go
type A struct{
    Num int
}
func (a A) test(){
    fmt.Println(a.Num)
}
```



方法和函数的区别:

1. 调用方式不一样 
   函数的调用方式： 函数名(实参列表) 
   方法的调用方式:  变量.方法名(实参列表) 

2. 对于普通函数,接收者为值类型时,不能将指针类型的数据直接传递,反之亦 然 

3. 对于方法(如struct的方法) ,接收者为值类型时,可以直接用指针类型的变量 调用方法，反过来同样也可以



## 2.面向对象

[目录](#目录)

go语言中的面向对象依然还是封装、继承、多态，但是和其他语言比如说java等还是有些许的不同，接下来展开的说一下：

### 封装

​	封装就是把抽象出来的字段和对字段的操作封装在一起，数据被保护在内部，程序的其他包只有通过被授权的操作（方法），才能对字段进行操作。

​	封装主要体现在对于结构体的属性进行封装，以及通过方法，包实现封装。

​	例如实现对于一个学生的结构体，将他们所具有的共同属性进行抽象封装成一个结构体

```go
type Student struct{
    Name string
    Age int
    id int
}
```



### 继承

继承的基本语法：

```go	
type Goods struct {
    Name string
    Price int
}
type Book struct {
    Goods //这里就是嵌套了匿名结构体Goods
    Writer string
}
```

​	对于go语言中的继承，不同于java的继承，它没有那么明确的父类子类的概念，通过结构体的嵌套实现，一个对象继承了另一个对象的属性。具体介绍如下：

1. 当多个结构体存在相同的属性(字段)和方法时,可以从这些结构体中抽象出结构体(比如刚才的Student)，在该结构体中定义这些相同的属性和方法。

2. 其它的结构体不需要重新定义这些属性和方法,只需嵌套一个Student匿名结构体即可。
3. 也就是说:在Golang中,如果一个struct嵌套了另一个匿名结构体,那么这个结构体可以直接访问匿名结构体的字段和方法，从而实现了继承特性。



### 多态

对于go语言中的多态，主要通过接口来实现，放在后面的接口部分进行讲解

[接口体现多态](#接口体现多态)



## 3.接口interface

[目录](#目录)

### 接口的介绍

1. 接口的基本概念：

   1. 基本介绍：

      interface类型可以定义一组方法，但是这些不需要实现。并且interface不能包含任何变量，到定义某个自定义类型要使用的时候，再根据具体情况吧方法写出来。

    2. 基本语法：

       ```go
       type 接口名 interface {
           method1(参数列表)返回值列表
           method2(参数列表)返回值列表
           ...
       }
       ```

       实现接口的所有方法：

       ```go
       func (t 自定义类型)method1(参数列表)返回值 {
           //方法实现
       }
       func (t 自定义类型)method2(参数列表)返回值 {
           //方法实现
       }
       ...
       ```

​	接口中的所有方法都没有方法体，即接口的方法都是没有实现的方法，接口体现了程序设计的多态和高内聚和低耦合的思想。

​	这Golang的接口，不需要显示的实现。只需要一个变量，含接口类型中的所有方法，那么这个变量就实现了这个接口。因此，Golang中没有implement这样的关键字

### 接口体现多态

对于接口体现出接口，通过一个具体的案例来说明：

```go
package main
import "fmt"
type Usb interface {
	Start()
	Stop()
}

type Camera struct {
}
type Phone struct {
}

type Computer struct {
}

func (c Camera) Start() {
	fmt.Println("相机开始工作")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}
func (p Phone) Start() {
	fmt.Println("手机开始工作")
}
func (p Phone) Stop() {
	fmt.Println("手机开始工作")
}
func (computer Computer) Working(u Usb) {
	u.Start()
	u.Stop()
}
func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	computer.Working(phone)
	computer.Working(camera)
}
```

对于以上的代码中，USB可以是Phone类型也可以是Camera类型，这就是多态的体现



### 接口和继承的区别

1. 接口和继承解决的解决的问题不同 
   1. 继承的价值主要在于:解决代码的复用性和可维护性。 
   2. 接口的价值主要在于：设计，设计好各种规范(方法)，让其它自定义类型去实现这些方法。 

2. 接口比继承更加灵活 
   1. 接口比继承更加灵活，继承是满足 is - a的关系，而接口只需满足 like- a的关系。 

3. 接口在一定程度上实现代码解耦



## 4.单元测试

[目录](#目录)

基本测试结构：
Go的测试文件以_test.go结尾，测试函数以Test开头，接受*testing.T参数。

下面一个案例：

`cal.go`文件中写入要测试的函数

```go
package main
func AddUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}
func getSub(n1, n2 int) int {
	return n1 - n2
}
```

分别创建`cal_test.go`和`sub_test.go`来进行测试

`cal_test.go`

```go
package main
import (
	"fmt"
	"testing"
)
func TestAddUpper(t *testing.T) {
	res := AddUpper(10)
	if res != 55 {
		t.Fatalf("addUpper执行错误，期望值=%v,实际值=%v\n", 55, res)
	}
	t.Logf("addUpper执行正确")
}
func TestHello(t *testing.T) {
	fmt.Println("TestHello被调用")
}
```

`sub_test.go`

```go
package main
import (
	"testing"
)
func TestGetSub(t *testing.T) {
	res := getSub(15, 10)
	if res != 5 {
		t.Fatalf("getSub执行错误，期望值=%v,实际值=%v\n", 5, res)
	}
	t.Logf("getSub执行正确")
}
```



## 5.文件操作

[目录](#目录)

### 文件夹的创建

在Go语言中，你可以使用标准库中的`os`包来创建文件夹（目录）。以下是几种常见的创建文件夹的方法：

#### 1. 使用`os.Mkdir`创建单个文件夹

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // 创建单个文件夹
    err := os.Mkdir("testdir", 0755)
    if err != nil {
        fmt.Println("创建文件夹失败:", err)
        return
    }
    fmt.Println("文件夹创建成功")
}
```

- `0755`是权限模式，表示：
  - 所有者：读、写、执行权限 (7)
  - 组用户：读、执行权限 (5)
  - 其他用户：读、执行权限 (5)

#### 2. 使用`os.MkdirAll`创建多级文件夹

如果需要创建多层嵌套的目录结构，可以使用`os.MkdirAll`：

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // 创建多级文件夹
    err := os.MkdirAll("parent/child/grandchild", 0755)
    if err != nil {
        fmt.Println("创建文件夹失败:", err)
        return
    }
    fmt.Println("多级文件夹创建成功")
}
```



### 文件的打开和关闭

1. 打开一个文件进行读操作：

   ```go
   os.Open(name string)(*File,error)
   ```

2. 关闭一个文件:

   ```go
   func (f *File) Close() error
   ```



### 文件的创建和写入

文件语法：

```go
func OpenFile(name string,flag int,perm FileMpde)(file *File,err error)
```

说明: `os.OpenFile`是一个更一般性的文件打开函数,它会使用指定的选项(如 O_RDONLY等)、指定的模式(如0666等)打开指定名称的文件。如果操作成功, 回的文件对象可用于1/0,如果出错,错误底层类型是*PathError.

其中第二个参数是设置文件的打开模式

```go
const (
    O_RDONLY int = syscall.O_RDONLY  // 只读模式打开文件
    O_WRONLY int = syscall.O_WRONLY  // 只写模式打开文件
    O_RDWR   int = syscall.O_RDWR    // 读写模式打开文件
    O_APPEND int = syscall.O_APPEND  // 写操作时将数据附加到文件尾部
    O_CREAT  int = syscall.O_CREAT   // 如果不存在将创建一个新文件
    O_EXCL   int = syscall.O_EXCL    // 与O_CREAT配合使用：文件必须不存在
    O_SYNC   int = syscall.O_SYNC    // 打开文件用于同步I/O
    O_TRUNC  int = syscall.O_TRUNC   // 如果可能：打开时清空文件
)
```

第三个参数：权限控制（在Linux中使用），暂时用不到随便填个数就行



**举个实例：创建一个新文件写入内容：5句"Hello World"**

```go
package main
import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	filePath := "e:/test01.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0)
	if err != nil {
		fmt.Println("打开文件错误")
		return
	}
	defer file.Close()

	//准备写入
	str := "Hello Golang\n"
	//写入时，使用带缓存的*writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//writer是带缓存的，因此在调用WriterString方法时
	//内容是先写入缓存的，所以要用Flush()将缓存的数据写入文件中
	err1 := writer.Flush()
	if err1 != nil {
		return
	}
	fmt.Println("文件写入成功")
}
```



### 拷贝文件

拷贝文件主要是使用io包中的copy()函数，分别获取到被复制文件内容以及要复制到的文件，一个是reader，一个是wirter,传入函数中实现拷贝

`func Copy(dst Writer,src Reader)(written int64,err error)`

下面是一个图片的复制案例：

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 编写一个函数，接受两个文件路径srcFileName dstFileName
func CopyFile(dstFileName, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("opne file err=%v", err)
	}
	defer srcFile.Close()
	//通过srcFile获取到Reader
	reader := bufio.NewReader(srcFile)
	//打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open flie error= %v", err)
		return
	}
	wirter := bufio.NewWriter(dstFile)
	defer dstFile.Close()

	return io.Copy(wirter, reader)
}
func main() {
	srcFile := "e:/qwe.jpg"
	dstFile := "f:/zxc.jpg"
	_, err := CopyFile(dstFile, srcFile)
	if err != nil {
		fmt.Printf("拷贝失败")
	} else {
		fmt.Printf("拷贝成功")
	}

}
```

在这段代码中实际上是调用io这个包中的Copy这个函数，这里编写了一个函数来分别获取wirter和reader从而调用该函数并返回结果，实现拷贝成功，该段程序中不止能够拷贝照片，还可以拷贝视频。



## 6.JSON

[目录](#目录)	

​	在JS语言中，一切都是对象，因此，任何的数据类型都可以通过JSON来表示，例如字符串、数字、对象、数组、map、结构体等

​	JSON键值对是用来保存，数据一种方式，键、值对组合中的键名写在前面并用双引号`""`包裹，使用冒号`:`分割，然后紧接着值

> https://www.json.cn/这个网站可以验证JSON数据是否正确

格式类似于这样

```json
{"k1":"v1","k2":"v2","k3":"v3","k4":"[v4,v5]"}

[{"k1":"v1","k2":"v2","k3":"v3","k4":"[v4,v5]"},
{"k1":"v1","k2":"v2","k3":"v3","k4":"[v4,v5]"}]
```

#### JSON序列化

**json的序列化是指，将有key-value结构的数据类型（比如结构体，map,切片）序列化成json字符串的操作**

主要用到的函数是`data,err :=json.Marshal()`返回值是一个切片,对于函数中传入的参数，可以使结构体、map以及切片

下面列举一个关于结构体序列化的例子：

```go
func testStruct() {
	monster := Monster{
		Name:     "牛魔王",
		Age:      1000,
		Birthday: "23-6-30",
		Sal:      10000.0,
		Skill:    "蛮牛冲撞",
	}
	//将monster序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化失败，err=%v", err)
	}
	//输出序列化后的结果
	fmt.Printf("monster序列化后=%v", string(data))
}
```



#### JSON的反序列化

**`json`反序列化是指，将`json`字符串反序列化成对应的数据类型（如结构体、map、切片）的操作**

主要用到的函数是`err :=json.Unmarshal([]byte(str),&monster)`对于函数中传入的参数为要转化的字符串，第二个是你要转换成类型的对象，下面我将以反序列化成结构体为例

```go
type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func unmarshalStruct() {
	//真正的是在项目中，是通过网络传输获取到的
	str := "{\"Name\":\"牛魔王\",\"Age\":1000,\"Birthday\":\"23-6-30\",\"Sal\":10000,\"Skill\":\"蛮牛冲撞\"}"
	//定义Monster
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("反序列化失败,err=%v", err)
	}
	fmt.Println("反序列化后monster= ", monster)
}
```

1. 在反序列化一个`json`字符串时，要确保反序列化后的数据类型和原来序列化数据类型一致
2. 如果`json`字符串是通过程序获取到的，则不需要在对双引号进行转义。



## 7.goroutine协程

[目录](#目录)

#### goroutine基本介绍

##### 1. 进程和线程的说明

1. 进程就是程序程序在操作系统中的一次执行过程，是系统进行资源分配和调度的基本单位 
2. 线程是进程的一个执行实例,是程序执行的最小单元，它是比进程更小的能独立运行的基本单位。
3. 一个进程可以创建核销毁多个线程,同一个进程中的多个线程可以并发执行。
4. 一个程序至少有一个进程,一个进程至少有一个线程

##### 2. 并行和并发的说明

1. 多线程程序在单核上运行,就是并发

2. 多线程程序在多核上运行,就是并行 

##### 3. go协程和go主线程

1. Go主线程(有程序员直接称为线程/也可以理解成进程):一个Go线程上,可以起多个协程,你可以这样理解,协程是轻量级的线程。

2. Go协程的特点:

   - 有独立的栈空间

   - 共享程序堆空间 

   - 调度由用户控制 

   - 协程是轻量级的线程

#### 协程的基本语法

在go程序中启动一个协程只需要使用一个go关键字

```go
//启动一个协程，go后面加上要执行功能的函数
go func(){}
```

开启协程之后，如果主线程退出了，即使协程还没有执行完毕，也会退出



#### 协程运行引发的资源竞争问题

比如，对于同一个map,同时启动两个协程对它进行操作，此时可能引发资源竞争问题引发死锁问题，而此时需要进行一些处理来解决此问题，一般有两种解决方案一是引用锁来对map进行保护，第二种是将数据存入管道

互斥锁：

```go
var lock sync.Mutex
//进行加锁
lock.Lock()
逻辑操作。。。。。
//释放锁
lock.Unlock()
```

管道后面再讲



#### MPG模式

Go语言的并发模型被称为**MPG模型**，这是Go语言高效并发能力的核心设计。MPG代表三个关键组件：**Machine (M)**、**Processor (P)** 和 **Goroutine (G)**。

##### 1 Goroutine (G)

- **轻量级线程**：比系统线程更轻量（初始栈约2KB，可动态扩容）
- **用户态调度**：由Go运行时管理，不直接依赖操作系统线程
- **低成本创建**：可轻松创建数十万个Goroutine

##### 1.2 Machine (M)

- **系统线程**：代表真正的操作系统线程
- **执行载体**：负责与操作系统的交互
- **绑定关系**：一个M必须持有一个P才能执行G

##### 1.3 Processor (P)

- **逻辑处理器**：Go运行时的资源调度单元
- **本地队列**：每个P维护一个本地Goroutine运行队列
- **数量控制**：默认等于CPU核心数（可通过`GOMAXPROCS`调整）



##### **MPG之间的协作关系**：

1. **M与P绑定**：每个工作线程(M)需要获取一个P才能执行G
2. **P管理G队列**：P维护本地运行队列和可运行的Goroutine
3. **G执行流程**：
   - 新建G进入P的本地队列
   - M获取P后从队列取出G执行
   - 当G阻塞时，M会释放P让其他M使用

##### 调度器的机制

##### 1. 工作窃取(Work Stealing)

当P的本地队列为空时，会尝试：

1. 从全局队列获取G
2. 从其他P的本地队列"窃取"一半的G

##### 2. 系统调用处理

- **阻塞型系统调用**：M会释放P进入阻塞状态，P被其他M获取
- **非阻塞型系统调用**：M会继续执行当前G

##### 3. 抢占式调度

Go 1.14+实现了基于信号的抢占：

- 防止长时间运行的Goroutine独占P
- 时间片默认10ms



## 8.channel管道

[目录](#目录)
对于前面提到的用锁来解决资源竞争的问题，虽然能解决，也并不利用多个协程对全局变量的读写操作。这就需要一个新的通讯机制-channel机制。

### channel的介绍

1. channle 本质就是一个数据结构 - 队列
2. 数据是先进先出
3. 线程安全，多 goroutine 访问时，不需要加锁，就是说 channel 本身就是线程安全的
4. channel 时有类型的，一个 string 的 channel 只能存放 string 类型数据。



### channel的基本用法

#### **管道的声明:**

`var 变量名 chan 数据类型`

1. channel是引用类型
2. channel必须初始化才能够写入数据，即需要make后才能使用
3. 管道是有类型的，声明一个`var intChan chan int `只能写入int 类型的数据



#### **管道的初始化:**

```go
var intChan chan int
intChan = make(chan int,10)
```

上面就是声明一个chan int类型的管道，对其进行初始化，容量为10



#### 管道数据的写入和取出

```go
//将10这个数写入管道
intChan <- 10
//取出管道数据
sum := <-intChan
```



### 管道的关闭和遍历

#### 管道的关闭

可以使用内置函数close()可以关闭channel,关闭channel之后就不能够再写入数据了，但是仍然能够从channel中取出数据。

```go
package main
import "fmt"
func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan)
	//intChan <- 300 此时会报错
	//读数据
	m1 := <-intChan
	fmt.Println(m1)
}
```

以上代码，在关闭管道后，即使还有空间但也无法继续添加数据，还是可以进行读数据



#### 管道的遍历

channel的遍历要使用for-ranger的方式进行遍历：

1. 在遍历的时候，如果channel没有关闭，则会出现deadlock，发生死锁
2. 在遍历的时候，如果channel已经关闭，则会正常遍历数据，遍历完之后，就会退出遍历

```go
//遍历管道
intChan2 := make(chan int, 100)
for i := 0; i < 100; i++ {
    intChan2 <- i * 2
}
close(intChan2)
for v := range intChan2 {
    fmt.Println("v=", v)
}
```

遍历管道之前必须要关闭管道



### channel使用细节和注意事项

#### channel可以声明为只读和只写

设置管道只读或者只写：

```go
var chan1 chan<- int
chan1 = make(chan int,2)
chan1<-1
//num := <-chan1 这时候会报错

var chan2 <-chan int
num1 := <-chan2
// chan2<- 3 此时会报错
```

这个在需要进行特殊操作的时候可以这样定义channel



#### 使用selcet可以解决从管道取数据的阻塞问题

`select`是Go语言中处理channel通信的核心控制结构，它允许goroutine同时等待多个通信操作（接收或发送），并执行其中一个准备就绪的操作。

##### 1. 基本语法和用法

```go
select {
case <-ch1:
    // 从ch1接收到数据时执行
case data := <-ch2:
    // 从ch2接收到数据时执行，数据存入data变量
case ch3 <- value:
    // 向ch3发送数据成功时执行
default:
    // 当没有任何case准备就绪时执行
}
```

##### 2.多路复用

`select`可以同时监听多个channel操作：

```go
func worker(ch1, ch2 <-chan int) {
    for {
        select {
        case x := <-ch1:
            fmt.Println("Received from ch1:", x)
        case y := <-ch2:
            fmt.Println("Received from ch2:", y)
        }
    }
}
```

##### 2. 随机选择

当多个case同时就绪时，`select`会随机选择一个执行：

```go
ch1 := make(chan int, 1)
ch2 := make(chan int, 1)
ch1 <- 1
ch2 <- 2
select {
case <-ch1:
    fmt.Println("Received from ch1")
case <-ch2:
    fmt.Println("Received from ch2")
}
// 输出可能是ch1或ch2，随机选择
```

##### 3. 阻塞与非阻塞

- **阻塞模式**：没有default子句时，select会阻塞直到某个case就绪
- **非阻塞模式**：有default子句时，select不会阻塞

##### 4.注意事项

1. **nil channel**：对nil channel的操作会永远阻塞
2. **已关闭的channel**：从已关闭的channel接收会立即返回零值
3. **内存泄漏**：长时间运行的select可能导致goroutine泄漏
4. **公平性**：select不保证公平性，长时间运行的case可能饿死其他case



## 9.反射

[目录](#目录)

### 1. 反射的基本介绍

1. 反射可以在运行时动态获取变量的各种信息，比如变量的类型，类别
2. 如果是反射结构体变量，还可以获取到结构体本身的信息（包括结构体的字段、方法）
3. 通过反射，可以修改变量的值，可以调用关联的方法
4. 使用反射，需要`import("reflect")`



### 2. 获取类型信息

#### 2.1 获取Type

```go
t := reflect.TypeOf(42)  // 获取int的类型信息
fmt.Println(t)          // 输出: int
```

#### 2.2 类型种类(Kind)

```go
var x float64 = 3.4
v := reflect.ValueOf(x)
fmt.Println("type:", v.Type())        // float64
fmt.Println("kind:", v.Kind())        // float64
fmt.Println("value:", v.Float())      // 3.4
```

#### 2.3 Kind与Type的区别

- Kind表示基础类型（int, float, struct等）
- Type包含更详细的类型信息（如结构体字段信息）

### 3. 获取和修改值

#### 3.1 获取值

```go
var x float64 = 3.4
v := reflect.ValueOf(x)
fmt.Println("value:", v.Float())  // 3.4
```

#### 3.2 修改值

```go
var x float64 = 3.4
v := reflect.ValueOf(&x).Elem()  // 必须获取可寻址的Value
v.SetFloat(7.1)
fmt.Println(x)  // 7.1
```

#### 3.3 可设置性

- `CanSet()`方法检查值是否可修改
- 只有可寻址的值才能被修改

### 4. 结构体反射

#### 4.1 获取结构体字段

```go
type User struct {
    Name string
    Age  int
}

u := User{"Alice", 25}
t := reflect.TypeOf(u)
for i := 0; i < t.NumField(); i++ {
    field := t.Field(i)
    fmt.Printf("%s: %v\n", field.Name, field.Type)
}
```

#### 4.2 获取字段值

```go
v := reflect.ValueOf(u)
for i := 0; i < v.NumField(); i++ {
    fmt.Printf("%v\n", v.Field(i).Interface())
}
```

#### 4.3 调用方法

```go
type Calculator struct{}

func (c Calculator) Add(a, b int) int {
    return a + b
}

c := Calculator{}
v := reflect.ValueOf(c)
method := v.MethodByName("Add")
args := []reflect.Value{reflect.ValueOf(2), reflect.ValueOf(3)}
result := method.Call(args)
fmt.Println(result[0].Int())  // 5
```

### 5. 反射与接口

#### 5.1 接口值到反射对象

```go
var x interface{} = "hello"
v := reflect.ValueOf(x)
t := reflect.TypeOf(x)
```

#### 5.2 反射对象到接口值

```go
v := reflect.ValueOf(42)
i := v.Interface()  // 转回interface{}
num := i.(int)      // 类型断言
```



对于反射中涉及到的大部分内容几乎都是reflect包中一些函数的使用

### 常用反射方法总结

| 方法/函数             | 用途                   |
| :-------------------- | :--------------------- |
| `reflect.TypeOf()`    | 获取类型信息           |
| `reflect.ValueOf()`   | 获取值信息             |
| `Value.Interface()`   | 将Value转回interface{} |
| `Value.Kind()`        | 获取基础类型           |
| `Value.SetXxx()`      | 设置值                 |
| `Value.Call()`        | 调用方法               |
| `Value.Field()`       | 获取结构体字段         |
| `Value.Elem()`        | 获取指针指向的值       |
| `reflect.New()`       | 创建新值               |
| `reflect.MakeSlice()` | 创建切片               |



## 10.context

[目录](#目录)

### 1. context基础概念

#### 1.1 context是什么

- 上下文管理工具，用于跨API边界传递截止时间、取消信号和其他请求范围的值
- 解决goroutine的超时控制、取消操作和值传递问题
- 线程安全，可以在多个goroutine间安全传递

#### 1.2 核心接口

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

### 2. 创建context的四种方式

#### 2.1 基础context

```go
ctx := context.Background()  // 空context，通常作为根context
ctx = context.TODO()         // 不确定使用何种context时使用
```

#### 2.2 派生context

```go
// 带取消功能的context
ctx, cancel := context.WithCancel(parentContext)
defer cancel() // 建议总是调用cancel释放资源

// 带超时的context
ctx, cancel := context.WithTimeout(parentContext, 2*time.Second)
defer cancel()

// 带截止时间的context
deadline := time.Now().Add(2 * time.Second)
ctx, cancel := context.WithDeadline(parentContext, deadline)
defer cancel()

// 带值的context
ctx := context.WithValue(parentContext, key, value)
```

### 3. context的核心功能

#### 3.1 取消传播

```go
func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("收到取消信号，停止工作")
            return
        default:
            // 正常工作
            fmt.Println("working...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}

// 调用示例
ctx, cancel := context.WithCancel(context.Background())
go worker(ctx)
time.Sleep(2 * time.Second)
cancel() // 取消所有派生出的操作
```

#### 3.2 超时控制

```go
func queryDB(ctx context.Context, query string) (string, error) {
    // 模拟数据库查询
    select {
    case <-time.After(5 * time.Second):
        return "result", nil
    case <-ctx.Done():
        return "", ctx.Err() // 返回取消原因
    }
}

// 调用示例
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
result, err := queryDB(ctx, "SELECT * FROM users")
```

#### 3.3 值传递

```go
type contextKey string

func main() {
    key := contextKey("userID")
    ctx := context.WithValue(context.Background(), key, "12345")
    
    // 在函数中获取值
    userID := ctx.Value(key).(string) // 类型断言
    fmt.Println(userID) // 输出: 12345
}
```

### 4. context使用最佳实践

#### 4.1 参数传递规则

- 将context作为函数的**第一个参数**传递
- 命名建议使用`ctx`而不是`context`

```go
func DoSomething(ctx context.Context, arg ArgType) error
```

#### 4.2 取消函数调用

- 总是**defer cancel()**，即使操作正常完成
- 确保资源及时释放，避免内存泄漏

#### 4.3 值传递注意事项

- 使用自定义类型作为key，避免字符串冲突
- 只传递请求范围的数据，不滥用context传值
- 值应该是不可变的

### 5. context的常见错误

#### 5.1 不检查Done通道

```go
// 错误示范
func badExample(ctx context.Context) {
    time.Sleep(3 * time.Second) // 可能永远不会被取消
    // ...
}

// 正确做法
func goodExample(ctx context.Context) {
    select {
    case <-time.After(3 * time.Second):
        // 工作完成
    case <-ctx.Done():
        return // 被取消
    }
}
```

#### 5.2 不调用cancel函数

```go
// 错误示范（可能导致内存泄漏）
ctx, cancel := context.WithCancel(context.Background())
go func() { /* 使用ctx */ }()
// 忘记调用cancel()

// 正确做法
ctx, cancel := context.WithCancel(context.Background())
defer cancel() // 确保cancel被调用
go func() { /* 使用ctx */ }()
```

### 6. context的高级用法

#### 6.1 组合多个context

```go
// 同时满足超时和手动取消
timeoutCtx, cancelTimeout := context.WithTimeout(ctx, 2*time.Second)
defer cancelTimeout()

cancelCtx, cancelFunc := context.WithCancel(timeoutCtx)
defer cancelFunc()

// 可以手动调用cancelFunc()提前取消
// 或者等待2秒后自动取消
```

#### 6.2 自定义context

```go
type customCtx struct {
    context.Context
    customValue string
}

func (c *customCtx) Value(key interface{}) interface{} {
    if key == "custom" {
        return c.customValue
    }
    return c.Context.Value(key)
}

// 使用示例
ctx := &customCtx{
    Context:     context.Background(),
    customValue: "hello",
}
fmt.Println(ctx.Value("custom")) // 输出: hello
```

### 7. context与标准库集成

#### 7.1 net/http

```go
// 服务器端
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    // 使用ctx进行超时控制
}

// 客户端
req, _ := http.NewRequest("GET", "http://example.com", nil)
ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
defer cancel()
req = req.WithContext(ctx)
resp, err := http.DefaultClient.Do(req)
```

#### 7.2 database/sql

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

// 带context的查询
rows, err := db.QueryContext(ctx, "SELECT * FROM users")
if err != nil {
    // 可能是context超时导致的错误
}
defer rows.Close()
```

### 8. context使用场景总结

| 场景       | 适用方法     | 说明                  |
| :--------- | :----------- | :-------------------- |
| 取消操作   | WithCancel   | 手动取消一组goroutine |
| 超时控制   | WithTimeout  | 设置绝对超时时间      |
| 截止时间   | WithDeadline | 设置具体截止时间点    |
| 值传递     | WithValue    | 跨API传递请求范围的值 |
| HTTP请求   | WithContext  | 控制请求超时          |
| 数据库查询 | XXXContext   | 控制查询超时          |

context是Go并发编程的重要工具，合理使用可以构建更健壮、更易维护的并发程序。



## 11.网络编程

[目录](#目录)

### 1. TCP基础概念

#### 1.1 TCP协议特点

- 面向连接的可靠传输协议
- 全双工通信
- 基于字节流而非消息边界
- 内置流量控制和拥塞控制

### 2. 服务端编程

#### 2.1 基本服务端流程

```go
package main

import (
    "net"
    "log"
)

func main() {
    // 1. 创建监听
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatal(err)
    }
    defer listener.Close()
    
    for {
        // 2. 接受连接
        conn, err := listener.Accept()
        if err != nil {
            log.Println("accept error:", err)
            continue
        }
        
        // 3. 处理连接(通常启动goroutine)
        go handleConn(conn)
    }
}

func handleConn(conn net.Conn) {
    defer conn.Close()
    
    // 4. 读写数据
    buf := make([]byte, 1024)
    for {
        n, err := conn.Read(buf)
        if err != nil {
            log.Println("read error:", err)
            return
        }
        log.Printf("received: %s", buf[:n])
        
        // 回显数据
        _, err = conn.Write(buf[:n])
        if err != nil {
            log.Println("write error:", err)
            return
        }
    }
}
```

#### 2.2 关键方法说明

- `net.Listen(network, address string)`：创建监听器
- `Listener.Accept()`：接受新连接
- `Conn.Read(b []byte)`：读取数据
- `Conn.Write(b []byte)`：写入数据
- `Conn.Close()`：关闭连接

### 3. 客户端编程

#### 3.1 基本客户端流程

```go
package main

import (
    "net"
    "log"
    "time"
)

func main() {
    // 1. 建立连接
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    
    // 2. 发送数据
    message := "hello, server"
    _, err = conn.Write([]byte(message))
    if err != nil {
        log.Println("write error:", err)
        return
    }
    
    // 3. 接收响应
    buf := make([]byte, 1024)
    n, err := conn.Read(buf)
    if err != nil {
        log.Println("read error:", err)
        return
    }
    log.Printf("received reply: %s", buf[:n])
}
```

#### 3.2 带超时的客户端

```go
// 设置连接超时
conn, err := net.DialTimeout("tcp", "localhost:8080", 3*time.Second)

// 设置读写超时
err = conn.SetDeadline(time.Now().Add(5 * time.Second))
```

### 4. 高级特性

#### 4.1 保持长连接

```go
// 服务端设置KeepAlive
conn.(*net.TCPConn).SetKeepAlive(true)
conn.(*net.TCPConn).SetKeepAlivePeriod(30 * time.Second)
```

#### 4.2 处理粘包问题

```go
// 使用长度前缀解决粘包
func sendMessage(conn net.Conn, message string) error {
    // 先发送长度(4字节)
    length := uint32(len(message))
    err := binary.Write(conn, binary.BigEndian, length)
    if err != nil {
        return err
    }
    
    // 再发送消息内容
    _, err = conn.Write([]byte(message))
    return err
}

func readMessage(conn net.Conn) (string, error) {
    // 先读取长度
    var length uint32
    err := binary.Read(conn, binary.BigEndian, &length)
    if err != nil {
        return "", err
    }
    
    // 根据长度读取内容
    buf := make([]byte, length)
    _, err = io.ReadFull(conn, buf)
    return string(buf), err
}
```

#### 4.3 并发连接管理

```go
type Server struct {
    connections map[net.Conn]struct{}
    mu          sync.Mutex
}

func (s *Server) addConn(conn net.Conn) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.connections[conn] = struct{}{}
}

func (s *Server) removeConn(conn net.Conn) {
    s.mu.Lock()
    defer s.mu.Unlock()
    delete(s.connections, conn)
    conn.Close()
}

func (s *Server) broadcast(message string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    for conn := range s.connections {
        _, err := conn.Write([]byte(message))
        if err != nil {
            log.Println("broadcast error:", err)
        }
    }
}
```

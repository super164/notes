# Socket简介

Socket在TCP/IP网络分层中并不存在，是对TCP或UDP封装

通俗点说Socket是实现网络上双向通讯连接的一套API,常称为套接字

## Socket分类：

按照连接时间：短连接 长连接（HTTP 1.1开始也支持长连接，Socket替换方案）

按照客户端和服务端数量：点对点 点对多 多对多

​	网络通信的内容都是包含客户端和服务端，服务端运行在服务器中，二客户端裕兴在客户端中，可以是浏览器，可以是桌面程序，也可以是手机app,客户端和服务端进行数据库交互遵守特定的协议。





# Go对Socket的支持

TCPAddr 结构体表示服务器 IP 和端口

- IP 是`type IP []byte`
- Port 是服务器监听的接口

```go
// TCPAddr represents the address of a TCP end point.
type TCPAddr struct {
    IP   IP //主机id
    Port int  //端口
    Zone string // IPv6 scoped addressing zone
}
```

- TCPConn 结构体表示连接，封装了数据读写操作

```go
// TCPConn is an implementation of the Conn interface for TCP network
// connections.
type TCPConn struct {
    conn
}
```

- TCPListener 负责监听服务器特定端口

```go
// TCPListener is a TCP network listener. Clients should typically
// use variables of type Listener instead of assuming TCP.
type TCPListener struct {
    fd *netFD
}
```



# 客户端向服务端发送消息

服务端：

```go
package main

import (
    "net"
    "fmt"
)

func main() {
    //创建TCPAddress变量,指定协议tcp4,监听本机8899端口
    addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")

    //监听TCPAddress设定的地址
    lis, _ := net.ListenTCP("tcp4", addr)

    fmt.Println("服务器已启动")

    //阻塞式等待客户端消息,返回连接对象,用于接收客户端消息或向客户端发送消息
    conn, _ := lis.Accept()

    //把数据读取到切片中
    b := make([]byte, 256)
    fmt.Println("read之前")
    //客户端没有发送数据且客户端对象没有关闭,Read()将会阻塞,一旦接收到数据就不阻塞
    count, _ := conn.Read(b)
    fmt.Println("接收到的数据:", string(b[:count]))
    //关闭连接
    conn.Close()
    fmt.Println("服务器结束")
}
```

客户端：

```go
package main

import (
    "net"
    "fmt"
)

func main() {
    //服务器端ip和端口
    addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")
    //申请连接客户端
    //第二个参数:本地地址 第三个参数:远程地址
    conn, _ := net.DialTCP("tcp4", nil, addr)
    //向服务端发送数据
    count, _ := conn.Write([]byte("客户端传递的数据"))
    fmt.Println("客户端向服务端发送的数据量为:", count)
    //通过休眠测试客户端对象不关闭,服务器是否能接收到对象
    //time.Sleep(10 * time.Second)
    ///关闭连接
    conn.Close()
    //fmt.Println("客户端结束")
}
```


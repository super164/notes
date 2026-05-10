# 网络聊天室 (Network Chat)



一个基于 Golang 实现的简单 TCP 网络聊天室，支持多客户端连接、公聊、私聊和在线用户列表查询功能。

## 功能特点



- 支持多客户端同时连接服务器（默认最大连接数 20）
- 用户名注册（确保用户名唯一，避免重复）
- 公聊功能：发送消息可被所有在线用户接收
- 私聊功能：通过特定格式向指定用户发送私密消息
- 在线用户列表查询：实时获取当前所有在线用户
- 优雅的连接管理：包含连接关闭、异常处理和资源释放机制

## 技术栈



- 编程语言：Golang
- 网络协议：TCP 协议
- 并发模型：Goroutine 实现并发处理
- 数据结构：Map 存储客户端连接信息
- 标准库：`net`（网络通信）、`bufio`（输入输出处理）等

## 项目结构



```plaintext
network_chat/
├── client/
│   └── client.go       # 客户端实现，负责连接服务器和消息交互
├── server/
│   └── server.go       # 服务器实现，负责监听连接和协程管理
├── common/
│   └── common.go       # 公共结构体(Client)、常量和工具函数
└── message/
    └── handler.go      # 消息处理逻辑（注册/公聊/私聊/列表查询）
```



## 使用方法



### 启动服务器



1. 进入服务器目录：

```
cd network_chat/server
```



1. 运行服务器程序：

```
go run server.go
```



服务器启动后将监听 `127.0.0.1:8080` 地址，等待客户端连接

### 启动客户端



1. 打开新终端，进入客户端目录：

```
cd network_chat/client
```



1. 运行客户端程序：

```
go run client.go
```



1. 按照提示输入用户名（若用户名已被占用，会提示重新输入）

### 聊天室操作命令



- 输入 `list` 并回车：查看当前所有在线用户列表
- 私聊格式：`>目标用户名:消息内容`（例如：`>Alice:你好，这是私聊`）
- 输入 `exit` 并回车：退出聊天室
- 直接输入消息内容：发送公聊消息（所有在线用户可见）

## 核心实现说明



1. 服务器架构

   ：

   - 通过 `net.Listen` 监听指定端口，使用循环接受客户端连接
   - 每个客户端连接由独立 Goroutine 处理（`process` 函数）
   - 全局 `ClientMap` 存储所有在线客户端信息（连接和用户名）

2. 客户端流程

   ：

   - 连接服务器后首先完成用户名注册（与服务器交互验证唯一性）
   - 启动独立 Goroutine 接收服务器消息（`receiveMsg` 函数）
   - 主线程处理用户输入并发送到服务器

3. 消息处理机制

   ：

   - 注册验证：`CheckNameExist` 函数检查用户名唯一性
   - 列表查询：`HandleListUser` 函数收集并返回在线用户
   - 私聊处理：`HandlePrivateChat` 函数解析格式并定向发送
   - 公聊处理：`HandleBroadcast` 函数向所有在线用户转发消息

## 注意事项



- 消息长度限制为 1024 字节，超过部分会被截断
- 未实现消息加密和身份认证，仅用于学习演示
- 最大连接数可通过修改 `common.ClientMap` 初始化参数调整（默认 20）
- 异常断开时服务器会自动清理用户连接信息 使用全局 Map 存储所有在线客户端连接信息，便于消息转发和用户列表查询 注意事项 本项目仅为学习演示用，未实现消息加密和身份验证 消息长度限制为 1024 字节 最大支持 20 个客户端同时连接（可通过修改 common 包中的 ClientMap 初始化参数调整）

# 代码部分

代码实现主要分为四个部分：server.go、client.go、handler.go、common.go

## common.go

对于项目中的一些公共的结构的定义

```go
package common

import "net"

// Client 客户端结构体定义
type Client struct {
	Conn net.Conn
	Name string
}

// ClientMap 全局客户端映射
var ClientMap = make(map[net.Conn]*Client, 20)

// 常量定义
const (
	CmdList    = "list" // 查看用户列表命令
	CmdPrivate = ">"    // 私聊前缀（如>用户名:消息）
)

// CloseConn 安全关闭连接的工具函数
func CloseConn(conn net.Conn) {
	if conn != nil {
		_ = conn.Close()
	}
}

```



## handler.go

对于消息的处理，判断消息的用法

```go
package message

import (
	"awesomeProject/network_chat/common"
	"fmt"
	"net"
	"strings"
)

// CheckNameExist 检查用户名是否存在
func CheckNameExist(conn net.Conn, client *common.Client) bool {
	name := client.Name
	for _, c := range common.ClientMap {
		if c.Name == name {
			return true
		}
	}
	return false
}

// HandleListUser 处理用户列表查询
func HandleListUser(conn net.Conn) bool {
	var userList string
	for _, c := range common.ClientMap {
		userList += c.Name + "\n"
	}
	msg := "当前在线用户：\n" + userList
	_, _ = conn.Write([]byte(msg))
	return true
}

// HandlePrivateChat 处理私聊消息
func HandlePrivateChat(senderConn net.Conn, msg string) bool {
	if !strings.HasPrefix(msg, common.CmdPrivate) {
		return false
	}
	// 解析格式：@用户名:消息内容
	parts := strings.SplitN(strings.TrimPrefix(msg, common.CmdPrivate), ":", 2)
	if len(parts) != 2 {
		_, _ = senderConn.Write([]byte("私聊格式错误，请使用：>用户名:消息内容\n"))
		return true
	}
	targetName, content := parts[0], parts[1]

	// 查找目标用户并发送消息
	senderName := common.ClientMap[senderConn].Name
	for conn, c := range common.ClientMap {
		if c.Name == targetName {
			privateMsg := fmt.Sprintf("[私聊]%s: %s\n", senderName, content)
			_, _ = conn.Write([]byte(privateMsg))
			return true
		}
	}
	_, _ = senderConn.Write([]byte("用户不存在\n"))
	return true
}

// HandleBroadcast 处理广播消息
func HandleBroadcast(senderConn net.Conn, msg string) {
	senderName := common.ClientMap[senderConn].Name
	broadcastMsg := fmt.Sprintf("%s: %s", senderName, msg)
	// 向所有其他客户端广播
	for conn := range common.ClientMap {
		if conn != senderConn {
			_, _ = conn.Write([]byte(broadcastMsg))
		}
	}
}

```



## server.go

服务端代码的管理

```go
package main

import (
	"awesomeProject/network_chat/common"
	"awesomeProject/network_chat/message"
	"fmt"
	"io"
	"net"
)

// 处理单个客户端连接
func process(conn net.Conn) {
	defer common.CloseConn(conn)
	client := &common.Client{Conn: conn}

	// 处理用户名注册
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取用户名失败：", err)
			return
		}
		client.Name = string(buf[:n])
		if message.CheckNameExist(conn, client) {
			_, _ = conn.Write([]byte("no")) // 用户名已存在
		} else {
			common.ClientMap[conn] = client
			_, _ = conn.Write([]byte("ok")) // 用户名可用
			break
		}
	}

	fmt.Printf("%s加入聊天室\n", client.Name)

	// 消息循环
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("用户%s退出聊天室\n", client.Name)
			} else {
				fmt.Printf("用户%s异常退出\n", client.Name)
			}
			delete(common.ClientMap, conn)
			return
		}

		msg := string(buf[:n])
		// 处理用户列表命令
		if msg == "list" {
			message.HandleListUser(conn)
			continue
		}
		// 处理私聊
		if message.HandlePrivateChat(conn, msg) {
			continue
		}
		// 广播消息
		message.HandleBroadcast(conn, msg)
		fmt.Print(common.ClientMap[conn].Name + ":" + msg)
	}
}

func main() {
	fmt.Println("聊天室创建成功")
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("监听失败：", err)
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("关闭监听连接失败")
		}
	}(listener)

	// 接受客户端连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接受连接失败：", err)
			continue
		}
		go process(conn)
	}
}

```



## client.go

用户端代码的管理

```go
package main

import (
	"awesomeProject/network_chat/common"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func menu() {
	fmt.Println("------------------------")
	fmt.Println("欢迎加入聊天室：")
	fmt.Println("输入list查看在线用户")
	fmt.Println("私聊格式>用户:消息")
	fmt.Println("输入exit退出聊天室")
	fmt.Println("------------------------")
}

// 处理用户名输入并注册
func registerName(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入网名：")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		if name == "" {
			fmt.Println("网名不能为空，请重新输入")
			continue
		}
		// 发送用户名到服务端
		_, err := conn.Write([]byte(name))
		if err != nil {
			fmt.Println("发送失败：", err)
			continue
		}
		// 接收服务端响应
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("接收响应失败：", err)
			continue
		}
		if string(buf[:n]) == "ok" {
			fmt.Println("注册成功，加入聊天室！")
			break
		} else {
			fmt.Println("网名已存在，请重新输入")
			continue
		}
	}
}

// 接收服务端消息并展示
func receiveMsg(conn net.Conn, done chan struct{}) {
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("\n连接已断开")
			close(done)
			return
		}
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("连接服务器失败：", err)
		return
	}
	defer common.CloseConn(conn)

	// 注册用户名
	registerName(conn)
	menu()
	// 启动消息接收协程
	done := make(chan struct{})
	go receiveMsg(conn, done)

	// 处理用户输入
	reader := bufio.NewReader(os.Stdin)
	for {
		select {
		case <-done:
			return
		default:
			//fmt.Print("请输入消息（输入exit退出，输入LIST查看用户，私聊格式@用户名:消息）：")
			line, _ := reader.ReadString('\n')
			line = strings.TrimSpace(line)

			if line == "" {
				fmt.Println("消息不能为空")
				continue
			}
			if strings.ToLower(line) == "list" {
				// 不添加换行符，直接发送
				_, err := conn.Write([]byte(line))
				if err != nil {
					fmt.Println("发送命令失败：", err)
					return
				}
				continue
			}
			if line == "exit" {
				fmt.Println("退出聊天室")
				return
			}
			// 发送消息到服务端
			_, err := conn.Write([]byte(line + "\n"))
			if err != nil {
				fmt.Println("发送消息失败：", err)
				return
			}
		}
	}
}
```


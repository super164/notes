# Gin + WebSocket 连接池

## 一、先搞懂 3 个基础概念

### 1. WebSocket 是什么？

- 普通 HTTP：**请求一次、响应一次**，就像发短信，一问一答。
- WebSocket：**长连接**，客户端和服务器**一直连着**，服务器可以主动给客户端发消息（像微信聊天、实时通知、弹幕）。
- Gin 中使用：官方推荐 `gorilla/websocket` 库，是 Golang 最稳定的 WebSocket 库。

### 2. 连接池是什么？

你可以把它理解成一个 **「在线用户花名册 + 消息中转站」**：

- 不是数据库那种连接池！

- 作用：**管理所有在线的 WebSocket 连接**。

- 能干什么：

  1. 知道谁在线
  2. 给某个用户发消息
  3. 给所有人广播消息
  4. 用户断开连接自动清理

  

### 3. 为什么 Gin + WebSocket 必须用连接池？

- 不用连接池：你只能处理**单个连接**，无法群发、无法管理在线用户，项目根本没法用。
- 用连接池：**统一管理所有客户端**，这是 WebSocket 项目的**标准必备结构**。

------

## 二、核心：WebSocket 连接池到底长啥样？

### 本质就是 3 样东西

1. **一个存储连接的容器**（map / 列表）
2. **注册机制**（用户连上来，存进池子）
3. **注销机制**（用户断开，从池子删除）
4. **消息广播方法**（给所有人发消息）

### 最简单的连接池结构（Go 代码）

```go
// 连接池：管理所有在线客户端
type Hub struct {
	// 在线客户端列表：key 是连接唯一标识，value 是连接对象
	clients map[*Client]bool

	// 广播消息通道：发消息就往这里扔
	broadcast chan []byte

	// 注册通道：客户端连上来，加入池子
	register chan *Client

	// 注销通道：客户端断开，移出池子
	unregister chan *Client
}

// 客户端：每个 WebSocket 连接就是一个 Client
type Client struct {
	hub  *Hub            // 所属连接池
	conn *websocket.Conn // WebSocket 原始连接
	send chan []byte     // 发送消息缓冲
}
```

------

## 三、连接池的工作流程

1. 启动项目，**连接池 Hub 先跑起来**（后台一直运行）
2. 浏览器 / APP 发起 WebSocket 连接
3. Gin 接口升级协议，创建 `Client` 对象
4. `Client` 发送到 `register` 通道 → **加入连接池**
5. 服务器要发消息：扔到 `broadcast` 通道
6. 连接池遍历所有在线客户端，把消息发出去
7. 用户断开连接 → 发送到 `unregister` 通道 → **从池子删除**

------

## 四、Gin 集成 WebSocket 连接池 完整代码

### 1. 先安装依赖

```bash
go get github.com/gin-gonic/gin
go get github.com/gorilla/websocket
```

### 2. 完整可直接运行代码（带注释）

```go
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 一、定义 连接池 Hub
type Hub struct {
	// 在线客户端
	clients map[*Client]bool

	// 广播消息
	broadcast chan []byte

	// 注册客户端
	register chan *Client

	// 注销客户端
	unregister chan *Client
}

// 创建连接池
func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

// 连接池核心逻辑：后台一直运行
func (h *Hub) Run() {
	for {
		select {
		// 客户端注册：加入池子
		case client := <-h.register:
			h.clients[client] = true

		// 客户端注销：删除并关闭通道
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}

		// 广播消息：发给所有在线客户端
		case message := <-h.broadcast:
			for client := range h.clients {
				client.send <- message
			}
		}
	}
}

// 二、定义 客户端 Client
type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

// WebSocket 升级器：把 HTTP 升级成 WebSocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 允许跨域（开发用）
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 三、Gin 接口：WebSocket 连接入口
func WsHandler(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 升级协议
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("升级失败：", err)
			return
		}

		// 2. 创建客户端
		client := &Client{
			hub:  hub,
			conn: conn,
			send: make(chan []byte, 256),
		}

		// 3. 注册到连接池
		client.hub.register <- client

		// 4. 启动读写协程
		go client.WritePump()
		go client.ReadPump()
	}
}

// 读消息：客户端 → 服务器
func (c *Client) ReadPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		// 收到消息后，直接广播给所有人
		c.hub.broadcast <- message
	}
}

// 写消息：服务器 → 客户端
func (c *Client) WritePump() {
	defer func() {
		c.conn.Close()
	}()

	for message := range c.send {
		_ = c.conn.WriteMessage(websocket.TextMessage, message)
	}
}

// 测试接口：主动给所有客户端发消息
func BroadcastMsg(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		msg := c.Query("msg")
		hub.broadcast <- []byte(msg)
		c.JSON(200, "发送成功")
	}
}

// 主函数
func main() {
	// 1. 创建并启动连接池
	hub := NewHub()
	go hub.Run()

	// 2. Gin 路由
	r := gin.Default()

	// WebSocket 连接地址
	r.GET("/ws", WsHandler(hub))

	// 测试广播消息
	r.GET("/send", BroadcastMsg(hub))

	// 3. 启动服务
	r.Run(":8080")
}
```

------

## 五、怎么测试？

1. 运行代码
2. 打开 **在线 WebSocket 测试工具**
3. 连接地址：`ws://127.0.0.1:8080/ws`
4. 打开多个网页，就是**多个连接**
5. 访问：`http://127.0.0.1:8080/send?msg=hello`
6. 所有连接都会收到消息 

------

## 六、你必须记住的 5 个关键点

1. **WebSocket 连接池 = 在线客户端管理中心**
2. 核心结构：`Hub`（池子）+ `Client`（每个连接）
3. 三大通道：`register`/`unregister`/`broadcast`
4. Gin 中必须用 `gorilla/websocket` 升级协议
5. 连接池要在 `go hub.Run()` 后台独立运行

------

## 七、直接用的扩展功能

1. **给指定用户发消息**：给 Client 加 UserID，用 map 存 `userID -> *Client`
2. **统计在线人数**：`len(hub.clients)`
3. **心跳检测**：客户端定时发心跳，超时断开
4. **房间功能**：创建多个 Hub，每个 Hub 是一个房间

------

### 总结

1. WebSocket 是长连接，支持服务器主动推消息
2. 连接池就是**管理所有在线连接**的工具
3. 结构 = `Hub`（池子）+ `Client`（连接）
## 一、什么是跨域？为什么要跨域？

### 1. 同源策略

浏览器的**同源策略**（Same-Origin Policy）是一个重要的安全机制，它限制了一个源（origin）的文档或脚本如何与另一个源的资源进行交互。

**源（origin）** 由三部分组成：

- 协议（http / https）
- 域名（[example.com](https://example.com/)）
- 端口（80 / 443 / 8080）

如果两个 URL 的协议、域名、端口都相同，则它们**同源**，否则就是**跨域**。

例如：

- `http://localhost:8080/api` 与 `http://localhost:8080/user` 同源（协议、域名、端口相同）
- `http://localhost:8080/api` 与 `http://localhost:3000/api` 跨域（端口不同）
- `http://example.com/api` 与 `https://example.com/api` 跨域（协议不同）

同源策略限制了以下行为：

- Cookie、LocalStorage 等无法跨域读取
- DOM 无法跨域访问
- AJAX 请求无法跨域发送（会被浏览器拦截）

### 2. 跨域需求

在实际开发中，前后端分离很常见：前端运行在 `http://localhost:3000`，后端运行在 `http://localhost:8080`。这时前端发起的 AJAX 请求就属于跨域请求，会被浏览器拦截。为了让前后端能够正常通信，就需要**跨域资源共享**（CORS）。

------

## 二、CORS 是什么？

**CORS（Cross-Origin Resource Sharing，跨域资源共享）** 是一种机制，它使用额外的 HTTP 头来告诉浏览器，允许某个源的网页访问另一个源的资源。

### 1. 基本流程

浏览器在发起跨域请求时，会携带 `Origin` 头，服务器在响应中携带 `Access-Control-Allow-Origin` 等头，浏览器根据这些头决定是否允许跨域访问。

### 2. 简单请求与非简单请求

#### 简单请求

满足以下所有条件：

- 方法为 `GET`、`HEAD`、`POST` 之一
- 头部只包含 `Accept`、`Accept-Language`、`Content-Language`、`Content-Type`（且值为 `application/x-www-form-urlencoded`、`multipart/form-data`、`text/plain`）
- 没有自定义头部

简单请求浏览器直接发送请求，响应中若 `Access-Control-Allow-Origin` 匹配当前源，则允许访问。

#### 非简单请求（预检请求）

不符合简单请求条件的请求（如使用 `PUT`、`DELETE`、自定义头、`Content-Type: application/json` 等），浏览器会先发送一个 **OPTIONS** 请求（预检请求），询问服务器是否允许实际请求。服务器需响应正确的 CORS 头，浏览器再根据结果决定是否发送实际请求。

------

## 三、Gin 框架中如何配置 CORS

在 Gin 中，通常使用中间件来处理 CORS。最常用的是官方维护的 `github.com/gin-contrib/cors` 包。

### 1. 安装

```bash
go get github.com/gin-contrib/cors
```



### 2. 基本用法

```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    r := gin.Default()

    // 使用默认 CORS 中间件
    r.Use(cors.Default())

    r.GET("/api/data", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello CORS"})
    })

    r.Run(":8080")
}
```

`cors.Default()` 的默认配置：

- 允许所有源 `*`
- 允许方法：`GET`、`POST`、`PUT`、`PATCH`、`DELETE`、`HEAD`、`OPTIONS`
- 允许头：`Origin`、`Content-Length`、`Content-Type`
- 不允许携带凭证（Credentials）

### 3. 自定义配置

更常见的做法是根据项目需求自定义 CORS 配置。

```go
package main

import (
    "time"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    r := gin.Default()

    // 自定义 CORS 配置
    r.Use(cors.New(cors.Config{
        // 允许的源，可以设置多个，或者使用 "*" 允许所有
        AllowOrigins:     []string{"http://localhost:3000", "https://example.com"},
        // 允许的方法
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        // 允许的头
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        // 暴露的头（前端可读取）
        ExposeHeaders:    []string{"Content-Length"},
        // 是否允许携带凭证（Cookie、HTTP认证等）
        AllowCredentials: true,
        // 预检请求缓存时间
        MaxAge:           12 * time.Hour,
    }))

    r.GET("/api/data", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello CORS"})
    })

    r.Run(":8080")
}
```



### 4. 使用通配符

如果需要允许所有源，可以将 `AllowOrigins` 设为 `[]string{"*"}`，但此时 `AllowCredentials` 必须为 `false`，因为浏览器不允许带凭证的通配符跨域。

如果既要允许所有源又要支持凭证，可以使用 `AllowOriginFunc` 动态判断：

```go
r.Use(cors.New(cors.Config{
    AllowOriginFunc: func(origin string) bool {
        // 可以在这里写自定义逻辑，比如检查 origin 是否在白名单中
        return true // 允许所有
    },
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    AllowCredentials: true,
}))
```



### 5. 预检请求处理

Gin 的 CORS 中间件会自动处理 `OPTIONS` 请求，无需手动编写。但如果你自己写中间件，需要确保对 `OPTIONS` 方法返回正确的响应头并停止后续处理。

------

## 四、手动实现 CORS 中间件（可选）

为了更深入理解，可以自己写一个简单的 CORS 中间件：

```go
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204) // 预检请求直接返回
            return
        }

        c.Next()
    }
}
```



然后在路由中调用：

```go
r.Use(CORSMiddleware())
```



------

## 五、注意事项

1. **预检请求 OPTIONS**
   如果前端发起的是非简单请求，会先发 OPTIONS。确保你的路由或中间件能正确处理 OPTIONS 方法，否则浏览器会报错。
2. **AllowCredentials 与通配符**
   `AllowCredentials: true` 时，`AllowOrigins` 不能为 `*`，必须指定具体源，或者使用 `AllowOriginFunc` 动态返回允许。
3. **安全考虑**
   生产环境中不要滥用 `*` 允许所有源，应只配置实际需要的前端域名，避免 CSRF 等风险。
4. **多个中间件顺序**
   CORS 中间件应放在最前面，确保在所有其他路由处理之前就设置好响应头，特别是 `OPTIONS` 请求能被正确拦截。
5. **自定义头**
   如果你的前端在请求中添加了自定义头（如 `X-Token`），需要在 `AllowHeaders` 中明确列出，否则浏览器会认为请求不合法。
6. **缓存预检结果**
   通过设置 `MaxAge` 可以减少不必要的 OPTIONS 请求，提升性能。

------

## 六、完整示例（带日志）

```go
package main

import (
    "time"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    r := gin.Default()

    // 配置 CORS
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    // 测试路由
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.Run(":8080")
}
```



前端访问 `http://localhost:3000` 下的页面，通过 AJAX 请求 `http://localhost:8080/ping`，即可成功跨域通信。

------

## 总结

- **同源策略** 是浏览器安全基础，CORS 是突破同源限制的标准方案。
- **CORS 通过 HTTP 头实现**，浏览器和服务器协商是否允许跨域。
- **Gin 框架中通常使用 `gin-contrib/cors` 中间件** 快速配置，也可手动实现。
- 配置时注意**预检请求**、**凭证与通配符的限制**，以及**生产环境的安全性**。
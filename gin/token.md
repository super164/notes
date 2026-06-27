# 全面解析 Token：从入门到 JWT 实战

在现代 Web 开发中，Token 已经成为身份验证和授权的重要手段。尤其在分布式系统、前后端分离和移动端应用场景下，理解 Token 的本质非常重要。本文将从零开始，帮你彻底理解 Token，并结合 Go 实战讲解 JWT 的使用。

------

## 一、为什么需要 Token？

### 1. 传统登录方式

传统的登录流程通常使用 **Session** 保存用户状态：

```
用户登录 → 服务端验证账号密码 → 创建 Session → 返回 SessionID → 浏览器存 Cookie → 后续请求带 Cookie → 服务端根据 SessionID 验证
```

问题：

- **服务器压力大**：大量用户登录需要存储 Session。
- **分布式麻烦**：多台服务器需要共享 Session（如通过 Redis），增加复杂度。
- **跨端不便**：移动端、API 调用等场景不好统一管理 Session。

------

### 2. Token 的出现

**Token 的核心思想**：

> 登录状态不再存储在服务器，而是存储在客户端，服务器只负责签发和验证。

形象比喻：

- Session → 前台登记入住，拿房卡，每次都要问前台
- Token → 酒店门卡，拿到后自己刷门，前台不用管你状态

Token 就是互联网世界的门卡。

------

## 二、Token 是什么？

**定义**：

> Token 是服务器签发给客户端的“身份凭证”，携带用户信息，并且经过签名防篡改。

特点：

1. 可以验证用户身份
2. 无需服务器存储状态
3. 适合分布式系统和跨端使用

------

## 三、Token 的工作流程

### 1. 登录阶段

1. 用户输入账号密码
2. 服务端验证成功
3. 生成 Token 并返回给客户端
4. 前端保存 Token（如 localStorage、Cookie 或内存）

### 2. 请求接口阶段

1. 前端在请求中带上 Token（通常在 Header）
2. 服务端验证 Token 的合法性
3. 验证成功 → 放行
4. 验证失败 → 拒绝访问

```
[登录]
用户 -> 服务端：账号密码
服务端验证
生成 Token
返回前端保存

[访问接口]
前端请求接口
Header: Authorization: Bearer Token
服务端验证签名
合法 -> 返回数据
非法 -> 拒绝
```

------

## 四、Token 的常见形式：JWT

JWT（JSON Web Token）是最常用的 Token 方案之一，格式如下：

```
xxxxx.yyyyy.zzzzz
```

三部分组成：

1. **Header（头部）**：算法信息
2. **Payload（载荷）**：用户信息 + 过期时间等
3. **Signature（签名）**：防篡改

------

### 1. Header

```
{
  "alg": "HS256",
  "typ": "JWT"
}
```

- `alg`：签名算法（HMAC-SHA256）
- `typ`：Token 类型（JWT）

------

### 2. Payload

存储用户信息或自定义数据：

```
{
  "userId": 1,
  "username": "admin",
  "role": "admin",
  "exp": 1715000000
}
```

- `exp`：过期时间
- `iat`：签发时间
- `iss`：签发者

注意：Payload 是可见的，不要存储密码等敏感信息。

------

### 3. Signature

使用 Header + Payload + Secret 生成：

```
HMACSHA256(base64(header) + "." + base64(payload), secret)
```

作用：保证 Token 不被篡改。

------

## 五、Token 的优势

| 优点           | 说明                     |
| -------------- | ------------------------ |
| 无状态         | 不需要服务器存储登录信息 |
| 适合分布式     | 多台服务器都能验证 Token |
| 跨端方便       | Web / App / 小程序统一   |
| 性能高         | 减少对数据库/Redis的访问 |
| 可携带业务信息 | Payload 可自定义字段     |

------

## 六、Token 的不足

1. 无法主动失效
   - Token 一旦签发直到过期都有效
   - 解决方案：使用黑名单（Redis 保存失效 Token）
2. 泄露风险
   - 谁拿到 Token 谁就可以使用
   - 注意保管，不要在 URL 或不安全环境传递

------

## 七、Access Token 和 Refresh Token

### 1. Access Token

- 短期有效（如 15 分钟）
- 用于访问受保护接口

### 2. Refresh Token

- 长期有效（如 7 天）
- 用于刷新 Access Token
- 可以降低 Access Token 泄露风险

------

## 八、Go 中的 JWT 实战

### 1. 生成 Token

```go
claims := jwt.MapClaims{
    "userId": user.ID,
    "role": user.Role,
    "exp": time.Now().Add(15*time.Minute).Unix(),
}

token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
tokenString, _ := token.SignedString([]byte(secret))
```

jwt.SigningMethod____:这是生成token的时候决定用什么算法来签发和验证

| 类型            | 加密方式                              | 适用范围                            |
| --------------- | ------------------------------------- | ----------------------------------- |
| HS256（最常用） | 对称加密：签发和验证都用同一个 Secret | 单体项目、中小型系统、内网服务      |
| RS256           | 非对称加密：私钥签发、公钥验证        | 微服务、OAuth / SSO、第三方登录平台 |



### 2. 验证 Token（Gin 中间件示例）

```go
func JWTAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(secret), nil
        })

        if err != nil || !token.Valid {
            c.AbortWithStatus(401)
            return
        }

        c.Next()
    }
}
```

------

## 九、Token 与 Session 的区别

| 对比项   | Session  | Token    |
| -------- | -------- | -------- |
| 存储位置 | 服务器   | 客户端   |
| 状态     | 有状态   | 无状态   |
| 分布式   | 需要共享 | 天然适合 |
| 跨端     | 麻烦     | 简单     |

------

## 十、总结

Token 是现代 Web 的身份凭证，尤其是 JWT，具备以下特点：

1. **无状态**：无需服务器存储
2. **防篡改**：签名保证安全
3. **跨端**：适合前后端分离、多终端
4. **扩展灵活**：Payload 可携带业务信息
5. **结合 Refresh Token 更安全**

一句话理解：

> Token 不是加密身份证，而是“服务器签名过的身份声明”。
>  客户端拿着它可以访问接口，服务端通过验签确认身份。
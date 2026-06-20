# AJAX 请求理解

## 一、什么是 AJAX？

**AJAX（Asynchronous JavaScript and XML）**
中文：异步 JavaScript 和 XML

### 本质

> 在**不刷新页面的情况下，与服务器进行数据交互**

### 作用

- 局部更新页面（提升用户体验）
- 提高性能（减少整页刷新）
- 实现前后端分离

------

## 二、AJAX 的核心特点

| 特点         | 说明                    |
| ------------ | ----------------------- |
| 异步         | 请求不阻塞页面运行      |
| 局部刷新     | 只更新部分 DOM          |
| 前后端通信   | 使用 HTTP 请求          |
| 数据格式灵活 | JSON（主流）、XML、Text |

------

## 三、AJAX 工作原理

 流程：

1. 用户触发事件（点击按钮等）
2. JavaScript 创建请求
3. 发送 HTTP 请求到服务器
4. 服务器返回数据
5. JS 接收数据并更新页面

简化理解：

```
浏览器(JS)  ←→  服务器(API)
```

------

## 四、AJAX 实现方式（重点）

### 原生 XMLHttpRequest（了解）

```js
var xhr = new XMLHttpRequest();

xhr.open("GET", "/api/user", true);

xhr.onreadystatechange = function () {
    if (xhr.readyState === 4 && xhr.status === 200) {
        console.log(xhr.responseText);
    }
};

xhr.send();
```

#### readyState 状态

| 值   | 含义       |
| ---- | ---------- |
| 0    | 未初始化   |
| 1    | 已建立连接 |
| 2    | 已发送请求 |
| 3    | 接收中     |
| 4    | 完成       |

------

### Fetch（主流 ）

```js
fetch("/api/user")
  .then(res => res.json())
  .then(data => {
      console.log(data);
  })
  .catch(err => {
      console.error(err);
  });
```

#### 优点

- 语法简洁
- 基于 Promise
- 更现代

------

### async / await（推荐 ）

```js
async function getUser() {
    try {
        const res = await fetch("/api/user");
        const data = await res.json();
        console.log(data);
    } catch (err) {
        console.error(err);
    }
}
```

 最接近同步写法，**面试加分项**

------

## 五、常见请求方式（HTTP）

| 方法   | 用途     |
| ------ | -------- |
| GET    | 获取数据 |
| POST   | 提交数据 |
| PUT    | 更新数据 |
| DELETE | 删除数据 |

------

## 六、发送 POST 请求示例

```js
fetch("/api/login", {
    method: "POST",
    headers: {
        "Content-Type": "application/json"
    },
    body: JSON.stringify({
        username: "admin",
        password: "123456"
    })
})
.then(res => res.json())
.then(data => console.log(data));
```

------

## 七、AJAX 常见数据格式

### JSON（最重要）

```js
{
  "username": "admin",
  "age": 20
}
```

JS 转换：

```
JSON.parse()   // 字符串 → 对象
JSON.stringify() // 对象 → 字符串
```

------

## 八、AJAX 优缺点

### 优点

- 用户体验好（不卡顿）
- 减少服务器压力
- 支持动态页面

### 缺点

- SEO 不友好
- 依赖 JS
- 存在跨域问题

------

## 九、跨域问题（重点 ）

###  什么是跨域

> 浏览器限制不同源之间的请求

#### 不同源指：

- 协议不同（http / https）
- 域名不同
- 端口不同

------

### 解决方案

#### CORS（最常用 ）

服务器返回：

```
Access-Control-Allow-Origin: *
```

------

#### JSONP（已过时）

只支持 GET

------

####  代理服务器（开发常用）

------

## 十、AJAX 在项目中的应用

结合你现在做的**用户管理系统**，AJAX可以用在：

###  示例

#### 1. 登录

```js
fetch("/api/login", { method: "POST" })
```

#### 2. 获取用户列表

```js
fetch("/api/users")
```

#### 3. 删除用户

```js
fetch("/api/user/1", {
    method: "DELETE"
})
```

#### 4. 修改用户信息

```js
fetch("/api/user/1", {
    method: "PUT"
})
```

------

## 十一、AJAX 面试常问问题

###  AJAX 和 Fetch 区别

- AJAX 是概念
- Fetch 是实现方式

------

###  AJAX 是同步还是异步？

 默认异步（可以同步但不推荐）

------

###  为什么推荐 async/await？

- 可读性强
- 避免回调地狱

------

### AJAX 和 axios 区别？

- axios 是库（封装更强）
- fetch 是原生 API

------

## 十二、总结（核心记忆）

 一句话理解：

> AJAX = 用 JS 在后台偷偷请求数据，不刷新页面更新内容
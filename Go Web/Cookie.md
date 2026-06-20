# 一、Cookie简介

- Cookei就是客户端存储技术，意见支队的形式存在
- 在B/S架构中，服务器端产生Cookie相应给客户端，浏览器接受后把Cookie存在在特定的文件夹中，以后每次请求浏览器会把Cookie内容放入到请求中



# 二、Go语言对Cookie的支持

在`net/http`包下提供了 Cookie 结构体

- `Name`：设置 Cookie 的名称
- `Value`：表示 Cookie 的值
- `Path`：有效范围
- `Domain`：可访问 Cookie 的域
- `Expires`：过期时间
- `MaxAge`：最大存活时间，单位秒
- `HttpOnly`：是否可以通过脚本访问

```go
type Cookie struct {
    Name  string
    Value string

    Path       string    // optional
    Domain     string    // optional
    Expires    time.Time // optional
    RawExpires string    // for reading cookies only

    // MaxAge=0 means no 'Max-Age' attribute specified.
    // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
    // MaxAge>0 means Max-Age attribute present and given in seconds
    MaxAge   int
    Secure   bool
    HttpOnly bool
    Raw      string
    Unparsed []string // Raw text of unparsed attribute-value pairs
}
```

# 三、代码演示

- 默认显示`index.html`页面，显示该页面时没有 Cookie，点击超链接请求服务器后，服务端把 Cookie 响应给客户端，通过开发者工具 (F12) 观察整个过程。

- html代码

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <script type="text/javascript" src="/static/js/index.js"></script>
</head>
<body>
<a href="setCookie">产生Cookie</a>
<a href="getCookie">获取Cookie</a>
<br>
{{.}}
</body>
</html>
```



- Go代码

```go
package main

import (
	"html/template"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, nil)
}
func setCookie(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{Name: "mykey", Value: "myValue"}
	http.SetCookie(w, &c)
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, nil)
}
func getCookie(w http.ResponseWriter, r *http.Request) {
	//根据key取cookie
	//c1, _ := r.Cookie("mykey")
	//取出全部的Cookie
	c2 := r.Cookies()

	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, c2)
}
func main() {
	server := http.Server{Addr: ":8090"}
	http.HandleFunc("/", welcome)
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookie", getCookie)
	server.ListenAndServe()
}

```



# 常用属性

## 一、HttpOnly

- `HttpOnly`：控制 Cookie 的内容是否可以被 JavaScript 访问到。通过设置`HttpOnly=true`是防止 XSS 攻击的防御手段之一
- 默认`HttpOnly=false`，表示客户端可以通过 js 获取
- 在项目中导入`jquery.cookie.js`库，使用 jquery 获取客户端 Cookie 内容

- HTML 代码如下

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <script src="/static/js/jquery-1.7.2.js"></script>
    <script src="/static/js/jquery.cookie.js"></script>
    <script type="text/javascript">
        $(function () {
            $("button").click(function () {
                var value = $.cookie("mykey")
                alert(value)
            })
        })
    </script>
</head>
<body>
<a href="setCookie">产生Cookie</a>
<button>获取cookie</button>
</body>
</html>
```

- go代码

```go
package main

import (
	"html/template"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, nil)
}
func doCookie(w http.ResponseWriter, r *http.Request) {
	//验证httpOnly
	c := http.Cookie{Name: "mykey", Value: "myValue", HttpOnly: true}
	http.SetCookie(w, &c)
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, nil)
}

func main() {
	server := http.Server{Addr: ":8090"}
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", welcome)
	http.HandleFunc("/doCookie", doCookie)

	server.ListenAndServe()
}
```



## Path

- `Path`属性设置 Cookie 的访问范围
- 默认为`"/"`表示当前项目下所有都可以访问
- `Path`设置路径及子路径内容都可以访问
- 首先访问`index.html`，点击超链接产生 cookie，在浏览器地址栏输入`localhost:8090/abc/mypath`后发现可以访问 cookie
- html 代码没有变化，只需要修改服务器端代码如下

```go
func doCookie(w http.ResponseWriter, r *http.Request) {
	//验证httpOnly
	//c := http.Cookie{Name: "mykey", Value: "myValue", HttpOnly: true}
	c := http.Cookie{Name: "mykey", Value: "myValue", Path: "/abc"}
	http.SetCookie(w, &c)
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, nil)
}
```



## Expires

------

- Cookie 默认存活时间是浏览器不关闭，当浏览器关闭后，Cookie 失效
- 可以通过`Expires`设置具体什么时候过期，Cookie 失效。也可以通过`MaxAge`设置 Cookie 多长时间后失效
- IE6,7,8 和很多浏览器不支持`MaxAge`，建议使用`Expires`
- `Expires`是`time.Time`类型，所以设置时需要明确设置过期时间
- 修改服务器端代码如下，只需要修改创建 Cookie 的代码，其他位置不变

```go
func doCookie(w http.ResponseWriter, r *http.Request) {
	//验证httpOnly
	//c := http.Cookie{Name: "mykey", Value: "myValue", HttpOnly: true}
	//c := http.Cookie{Name: "mykey", Value: "myValue", Path: "/abc"}
    //time.NOW().Add()
    c := http.Cookie{Name: "mykey", Value: "myValue", Expires: time.Date(2026, 3, 1, 1, 1, 1, 0, time.Local)}
	http.SetCookie(w, &c)
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, nil)
}
```


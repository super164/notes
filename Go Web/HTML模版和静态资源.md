# 一、项目模版

- go语言中web项目标准结构如下

```go
--项目名
	--src
	--static
		--css
		--images
		--js
	--view
		--index.html
	--main.go
```

- Go语言标准库中html/template包提供了html模版支持，把HTML当做模版可以在访问控制器时显示HTML模版信息

# 二、HTML模版显示

- 使用emplate.ParseFiles()可以解析多个模版文件

```go
func ParseFiles(filenames ...string)(*Template,error){
    return parseFiles(nil,filenames...)
}
```

- 把模版信息相应写入到输出流中

```go
func (t *Template) Execute(wr io.Writer,data interface{} error){
    if err:=t.escape();err!=nil{
        return err
    }
    return t.text.Execute(wr,data)
}
```

- 代码演示显示index.html信息

```go
package main

import (
    "net/http"
    "html/template"
)

func welcome(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("view/index.html")
    t.Execute(w, nil) //第二个参数表示向模板传递的数据
}

func main() {
    server := http.Server{Addr: ":8090"}
    http.HandleFunc("/", welcome)
    server.ListenAndServe()
}
```

# 三、应用静态文件

- 把静态问价放入到特定的文件夹中，使用Go原因的文件服务就可以进行加载
- 项目结构

```go
--项目
	--static
		--js
			--index.js
	--view
		--index.html
	--main.go
```

index.html文件

```html	
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <script type="text/javascript" src="/static/js/index.js"></script>
</head>
<body>
这是显示的html文件，被显示了vhbjkjvhbgvhbj
<button onclick="myclick()">按钮</button>
</body>
</html>
```

index.js代码如下

```js
function myclick(){
    alert("你点击了按钮")
}
```

代码示例

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
func main() {
	server := http.Server{Addr: ":8090"}
	//访问url以/static/开头，就会把访问信息映射到指定的目录中

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", welcome)
	server.ListenAndServe()
}

```


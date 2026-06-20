# 一、Action

- "Arguments"和"pipelines"











# 二、if

- if写在模版中和写在go文件中功能是相同的，区别是语法
- 布尔函数会将任何类型的零值视为假，其余视为真
- if后面的表达是中包含逻辑控制符在模版中试线的全局函数

- 二院比较运算的集合

```
运算符	说明
eq	如果 arg1 = arg2 则返回真
ne	如果 arg1 != arg2 则返回真
lt	如果 arg1 < arg2 则返回真
le	如果 arg1 <= arg2 则返回真
gt	如果 arg1 > arg2 则返回真
ge	如果 arg1 >= arg2 则返回真
```



- 简单if示例-go文件

```go
package main

import (
    "net/http"
    "html/template"
)

func test(rw http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("template/html/if.html")
    //第二个参数传递类型默认值:nil,"",0,false都会导致if不成立
    t.Execute(rw, "")
}

func main() {
    //创建server服务
    server := http.Server{Addr: ":8090"}

    //设置处理器函数
    http.HandleFunc("/test", test)

    //监听和开始服务
    server.ListenAndServe()
}
```



- 简单if示例-html文件

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>if测试</title>
</head>
<body>
测试if是否执行<br
{{if .}}
if成立这个位置输出
{{end}}
</body>
</html>
```



- 使用else...if用法

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>if测试</title>
</head>
<body>
{{$n:=123}}
{{if ne $n 123}}
if成立这个位置输出
{{else}}
这是else功能
{{end}}
</body>
</html>
```



- if..else if...else用法

```html
<body>
{{$n:=124}}
{{if eq $n 123}}
123
{{else if eq $n 124}}
124
{{else if eq $n 125}}
125
{{else}}
else
{{end}}
</body>
```





# 三、range使用

- range遍历数组或切片或map或channel是，在range内容中{{.}}表示获取迭代变量

```html
<body>
    {{range .}}
    	{{.}}{{/*dot为迭代变量*/}}
    {{end}}
    
    {{.}}{{/*获取还是传递给模版的切片*/}}
</body>
```


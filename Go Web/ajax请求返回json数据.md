# 一、JSON简介

- 轻量级数据传输格式、总体上分为两种

  - 一种是JSONObject对象

    ```go
    {"key":"value","key":"value"}
    ```

  - 一种是JSONArrayP(json数据)，包含多个JSONObject

    ```go
    [{"key":"value"},{"key":"value"}]
    ```

    

- `key` 是 string 类型，`value` 可以是 string 类型（值被双引号包含），也可以是数值或布尔类型等，也可以是 JSONObject 类型或 JSONArray 类型

- 可以使用 Go 语言标准库中 encoding/json包下的 Marshal() 或 Unmarshal() 把结构体对象转换成 []byte

   或把 []byte中信息写入到结构体对象中

  - 在转换过程中结构体属性 `tag` 中定义了 json 中的 `key`，属性的值就是 json 中的 `value`
  - 如果属性没有配置 `tag`，属性就是 json 中的 `key`

- 属性的 `tag` 可以进行下面配置

------

### 代码示例

```go
// 字段被本忽略
Field int `json:"-"`

// 字段在json里的键为"myName"
Field int `json:"myName"`

// 字段在json里的键为"myName"且如果字段为空值将在对象中省略掉
Field int `json:"myName,omitempty"`

// 字段在json里的键为"Field"（默认值），但如果字段为空值会跳过；注意前导的逗号
Field int `json:",omitempty"`
```



## 二、代码示例

- 结构体和[]byte进行转换代码比较简单
  - 只要满足键值对形式的类型都可以转换成标准的 json 格式

```go
package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    Name string
    Age  int
}

func main() {
    user := User{"张三", 12}
    // user := map[string]interface{}{"Name": "张三", "Age": 18}
    // 把结构体转换为[]byte
    b, _ := json.Marshal(user)
    fmt.Println(string(b))
    // 把[]byte转为json
    u2 := new(User)
    json.Unmarshal(b, u2)
    fmt.Println(u2)
}
```

# 三、Ajax 访问返回 json 数据

- 使用 jQuery 封装的`$.post()`进行 ajax 请求
- HTML 页面发送 ajax 请求，请求数据

```html
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
    "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
    <title>Title</title>
    <script type="text/javascript" src="/static/js/jquery-1.7.2.js"></script>
    <script type="text/javascript">
        $(function () {
            $("button").click(function () {
                $.post("getUser", function (data) {
                    var result = "";
                    for (var i = 0; i < data.length; i++) {
                        result += "<tr>";
                        result += "<td>";
                        result += data[i].Name;
                        result += "</td>";
                        result += "<td>";
                        result += data[i].Age;
                        result += "</td>";
                        result += "</tr>";
                    }
                    $("#t_tbody").html(result)
                })
            })
        })
    </script>
</head>
<body>
<button>加载数据到表格</button>
<table border="1">
    <tr>
        <th>姓名</th>
        <th>年龄</th>
    </tr>
    <tbody id="t_tbody">
    </tbody>
</table>
</body>
</html>
```
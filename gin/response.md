gin框架响应一些结构的数据

```go
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// 响应字符串
func _string(context *gin.Context) {
    context.String(200, "你好")
}

// 响应json
func _json(c *gin.Context) {
    type Userinfo struct {
       UserName string `json:"username""`
       Age      int    `json:"age""`
       PassWord string `json:"-"`
    }
    user := Userinfo{"杨帆", 18, "123456"}
    c.JSON(200, user)
    c.JSON(200, gin.H{"username": "李想", "age": 19})
}

// 相应xml
func _xml(c *gin.Context) {
    c.XML(200, gin.H{"user": "hr", "message": "hey", "status": http.StatusOK, "data": gin.H{"user": "futao"}})
}

// 响应yaml
func _yaml(c *gin.Context) {
    c.YAML(200, gin.H{"user": "hr", "message": "hey", "status": http.StatusOK, "data": gin.H{"user": "futao"}})
}

// 响应html
func _html(c *gin.Context) {
    type Userinfo struct {
       UserName string `json:"username"`
       Age      int    `json:"age"`
       PassWord string `json:"-"`
    }
    user := Userinfo{
       UserName: "wuhan",
       Age:      9178,
       PassWord: "123456",
    }
    //c.HTML(200, "index.html", gin.H{"username": "futao"})
    c.HTML(200, "index.html", gin.H{"username": user.UserName, "age": user.Age, "password": user.PassWord})
}

// 重定向
func _redirect(c *gin.Context) {
    c.Redirect(302, "https://www.fengfengzhidao.com")
}
func main() {
    route := gin.Default()
    route.LoadHTMLGlob("templates/*")
    //在go中，没有相对文件的路径，只有相对项目的路径
    //访问路径 文件路径
    route.StaticFile("/go.jpg", "static/go.jpg")
    //网页请求静态目录的前缀，第二个参数是一个目录，前缀不能重复

    route.StaticFS("/static", http.Dir("./static/static"))
    route.GET("/", _string)
    route.GET("/json", _json)
    route.GET("/xml", _xml)
    route.GET("/yaml", _yaml)
    route.GET("/html", _html)
    route.GET("/baidu", _redirect)
    route.Run(":8080")
}
```
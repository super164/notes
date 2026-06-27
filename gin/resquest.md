# 请求参数

```go
func _query(c *gin.Context) {
    user := c.Query("user")
    fmt.Println(user)
    //判断是否传输
    fmt.Println(c.GetQuery("user"))
    //拿到多个相同的查询参数
    fmt.Println(c.QueryArray("user"))
    //拿到多个相同的查询参数
    userMap := c.QueryMap("user")
    fmt.Println("QueryMap 返回的完整 map：", userMap)
}
http://localhost:8080/query?user=付涛&user=付涛
```

# 动态参数

```go
func _param(c *gin.Context) {
    fmt.Println(c.Param("user_id"))
    fmt.Println(c.Param("book_id"))
}

//路径
router.GET("/param/:user_id", _param)
	router.GET("/param/:user_id/:book_id", _param)
```





# 表单参数

可以接受multipart/form-data和application/x-www-form-urlencoded

```go
// 表单参数POSTFORM
func _form(c *gin.Context) {
    fmt.Println(c.PostForm("name"))
    //接受多个值
    fmt.Println(c.PostFormArray("name"))

    //如果用户没有传值就使用默认值
    fmt.Println(c.DefaultPostForm("addr", "河南省"))
    //接受所有的form参数，包括文件
    forms, err := c.MultipartForm()
    fmt.Println(forms, err)
}

router.POST("/form", _form)
```





# 原始参数

form-data

```go
----------------------------659138997098209924904601
Content-Disposition: form-data; name="name"

abc
----------------------------659138997098209924904601
Content-Disposition: form-data; name="addr"

北京市
----------------------------659138997098209924904601--
```



x-www-form-unlencoded

```go
name=abc&addr=%E5%8C%97%E4%BA%AC%E5%B8%82
```



json

```json
{
    "name":"futao",
    "age":18
}
```



```go
// 原始参数GetRawData
func _raw(c *gin.Context) {
    data, _ := c.GetRawData()
    contentType := c.GetHeader("Content-Type")
    switch contentType {
    case "application/json":
       //json解析到结构体
       type User struct {
          Name string `json:"name"`
          Age  int    `json:"age"`
       }
       var user User
       //fmt.Println(string(data))
       err := json.Unmarshal(data, &user)
       if err != nil {
          fmt.Println(err.Error())
       }
       fmt.Println(user)
    }

}
```

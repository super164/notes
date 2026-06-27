# Gin验证器

### Gin binding 验证标签速查表

| 类别           | 标签                   | 示例 (`binding:"..."`)          | 说明                                                         |
| :------------- | :--------------------- | :------------------------------ | :----------------------------------------------------------- |
| **必填与忽略** | `required`             | `binding:"required"`            | 字段必须存在且不能是其类型的空值。                           |
|                | `-`                    | `binding:"-"`                   | 忽略该字段，Gin不会对其进行任何绑定和验证。                  |
|                | `omitempty`            | `binding:"omitempty,min=1"`     | 如果字段为空值（如其类型的零值），则跳过后续的验证规则。     |
| **比较与范围** | `len`                  | `binding:"len=6"`               | 验证字符串、数组或切片的长度（字符数/元素数）。              |
|                | `min` / `max`          | `binding:"min=2,max=20"`        | 对数值限制大小，对字符串限制长度。                           |
|                | `gt` / `lt`            | `binding:"gt=0,lt=10"`          | **严格**大于或小于，与`gte`/`lte`不同。                      |
|                | `gte` / `lte`          | `binding:"gte=18,lte=60"`       | **大于等于**或**小于等于**。                                 |
|                | `ne`                   | `binding:"ne=0"`                | 值不等于指定值。                                             |
|                | `oneof`                | `binding:"oneof=admin user"`    | 值必须是枚举列表中的一个。                                   |
| **字符与内容** | `contains`             | `binding:"contains=@gmail.com"` | 字符串必须包含子串。                                         |
|                | `excludes`             | `binding:"excludes=admin"`      | 字符串不能包含子串。                                         |
|                | `alpha`                | `binding:"alpha"`               | 字符串只能包含字母字符。                                     |
|                | `alphanum`             | `binding:"alphanum"`            | 字符串只能包含字母和数字字符。                               |
|                | `numeric`              | `binding:"numeric"`             | 字符串是一个数值。                                           |
| **格式与类型** | `email`                | `binding:"email"`               | 验证是否为有效的电子邮件格式。                               |
|                | `url`                  | `binding:"url"`                 | 验证是否为有效的URL地址。                                    |
|                | `ip`                   | `binding:"ip"`                  | 验证是否为有效的IP地址（v4或v6）。                           |
|                | `ipv4` / `ipv6`        | `binding:"ipv4"`                | 验证是否为有效的IPv4或IPv6地址。                             |
|                | `uuid`                 | `binding:"uuid"`                | 验证是否为有效的UUID。                                       |
|                | `json`                 | `binding:"json"`                | 验证字符串是否为有效的JSON。                                 |
|                | `boolean`              | `binding:"boolean"`             | 验证字符串是否为有效的布尔值（"1", "true", "0", "false"等）。 |
| **字段关系**   | `eqfield`              | `binding:"eqfield=Password"`    | 必须等于同一结构体中的另一个字段。                           |
|                | `nefield`              | `binding:"nefield=Username"`    | 不能等于同一结构体中的另一个字段。                           |
|                | `ltfield` / `ltefield` | `binding:"ltfield=EndTime"`     | 验证当前字段的值小于/小于等于同一结构体中另一个字段的值。    |
|                | `gtfield` / `gtefield` | `binding:"gtfield=StartTime"`   | 验证当前字段的值大于/大于等于同一结构体中另一个字段的值。    |

> **提示**：标签可以组合使用，用英文逗号`,`分隔即可，例如 `binding:"required,min=2,max=20,alpha"`。

## 常用验证器（在binding标签中设置）

required：必填字段，如：binding:"required" 不能不传，不能为空

针对字符串的长度：

min 最小长度，如：binding:"min=5" 

max 最大长度，如：binding:"max=10" 

len 长度，如：binding:"len=6"

针对数字的大小：

eq 等于，如：binding:"eq=3"

ne 不等于，如：binding:"ne=12"

gt 大于，如：binding:"gt=10"

gte 大于等于，如：binding:"gte=10"

lt 小于，如：binding:"lt=10"

lte 小于等于，如：binding:"lte=10"

//针对同级字段的

eqfield 等于其他字段的值，如：Password string `binding:"eqfield=ConfirmPassword"`

nefield 不等于其他字段的值

- 忽略字段，如：binding:"-"

```go
type SignUserInfo struct {
    Name       string `json:"name" binding:"min=4,max=6"`              //用户名
    Age        int    `json:"age" binding:"lt=30,gt=18"`               //年龄
    Password   string `json:"password" `                               //密码
    RePassword string `json:"re_password" binding:"eqfield=Password""` //确认密码

}
```





## Gin内置验证器

```go
//枚举
oneof=red green
Sex        string `json:"sex" binding:"oneof=man woman"` //限制性别只能是男或女

//字符串
contains=fengfeng //包含fengfeng的字符串
excludes //不包含
startswith //字符串前缀
endswith //字符串后缀

//数组
dive //dive后面的数据就是针对数组中的每一个元素
LikeList   []string `json:"like_list" binding:"required,dive,startswith=like"`

//网络验证
//uri是统一资源标识符，可以唯一表示一个资源
//url是统一资源定位符，提供找到该资源的确切路径
ip ipv4 ipv6 uri url


//日期验证
datetime=2006-01-02 15:04:05
//1月2号下午3点4分5秒在2006年
```

示例

```go
type SignUserInfo struct {
    Name       string   `json:"name" binding:"endswith=f"`               //用户名
    Age        int      `json:"age" binding:"lt=30,gt=18"`               //年龄
    Password   string   `json:"password" `                               //密码
    RePassword string   `json:"re_password" binding:"eqfield=Password""` //确认密码
    Sex        string   `json:"sex" binding:"oneof=man woman"`
    LikeList   []string `json:"like_list" binding:"required,dive,startswith=like"`
    Ip         string   `json:"ip" binding:"ip"`
    Url        string   `json:"url" binding:"url"`
    Uri        string   `json:"uri" binding:"uri"`
    Date       string   `json:"date" binding:"datetime=2006-01-02 15:04:05"`
}
```





## 自定义错误信息





```go
package main

import (
    "fmt"
    "reflect"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
)

// 获取结构体中msg参数

func GetValidMsg(err error, obj any) string {
    //err断言为具体类型
    getObj := reflect.TypeOf(obj)
    if errs, ok := err.(validator.ValidationErrors); ok {
       //断言成功
       for _, e := range errs {
          //循环错误信息
          //根据报错字段获取结构体的具体字段
          if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
             msg := f.Tag.Get("msg")
             return msg
          }
       }

    }
    return err.Error()
}

func main() {
    router := gin.Default()

    router.POST("/", func(c *gin.Context) {
       type User struct {
          Name string `json:"name" binding:"required" msg:"用户名校验失败"`
          Age  int    `json:"age" binding:"required" msg:"年龄不能为空"`
       }
       var user User
       err := c.ShouldBindJSON(&user)
       if err != nil {

          fmt.Println(err)
          c.JSON(200, gin.H{"msg": GetValidMsg(err, &user)})
          return
       }

       c.JSON(200, gin.H{"data": user})
    })

    router.Run(":8080")
}
```





## 自定义验证器

1. 注册验证器函数

   ```go
   //gin框架版本必须是V10的
   if v,ok:=binding.Validator.Engine().(*validator.Validate):ok{
       v.RegisterValidation("sign",signValid)
   }
   ```

2. 编写函数

   ```go
   //如果用户名不等于fengfeng就校验失败
   func signValid(fl validator.FieldLevel)bool{
       name:=fl.Field()Interface().(string)
       if name!="fengfeng"{
           return false
       }
       return true
   }
   
   ```

3. 使用

   ```go
   type UserInfo struct{
       Name string `json:"name" binding:"sign" msg:"用户名错误"`
       Age int `json:"age" binding:""`
   }
   ```

   
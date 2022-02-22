package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

/*绑定form表单数据
type User struct {
	Username string   `form:"username"` //对应的表单名字
	Password string   `from:"password"`
	Hobby    []string `form:"hobby"`
	Gender   string   `form:"gender"`
	City     string   `form:"city"`
}

func Register(c *gin.Context) {
	var user User
	//根据Content-Type（内容类型）默认解析并绑定form格式
	//解析错误直接返回，至于要给客户端返回什么错误状态码有你决定
	c.ShouldBind(&user)
	c.String(200, "User:%s", user)
}

func GoRegister(c *gin.Context) {
	//get请求返回指定网页
	c.HTML(200, "register.html", nil)
}
func main() {
	//创建路由
	r := gin.Default()
	r.LoadHTMLGlob("5.gin数据绑定/templates/*")
	//绑定路由规则和执行函数
	r.GET("/register", GoRegister)
	r.POST("/register", Register)
	//监听端口，默认8080
	r.Run()
}
*/

/*路径请求参数绑定
type User struct {
	Username string `uri:"username"`
	Password string `uri:"password"`
}

func TestGetBind(c *gin.Context) {
	var user User
	//解析并绑定uri格式
	//解析错误直接返回，至于要给客户端返回什么错误状态码有你决定
	err := c.ShouldBindUri(&user)
	if err != nil {
		//相当于fmt.print后面跟着os.Exit(1)
		log.Fatal(err)
	}
	c.String(200, "User:%s", user)
}
func main() {
	//创建路由
	r := gin.Default()
	//绑定路由规则和执行函数
	r.GET("/testGetBind/:username/:password", TestGetBind)
	//监听端口，默认8080
	r.Run()
}
*/

/*
type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func TestGetBind(c *gin.Context) {
	var user User
	//根据Content-Type（内容类型）默认解析并绑定form格式
	//解析错误直接返回，至于要给客户端返回什么错误状态码有你决定
	err := c.ShouldBind(&user)
	if err != nil {
		//相当于fmt.print后面跟着os.Exit(1)
		log.Fatal(err)
	}
	//将给定的字符串写给响应体
	c.String(200, "User:%s", user)
}

func main() {
	//创建路由
	e := gin.Default()
	//绑定路由规则和处理函数
	e.GET("/testGetBind", TestGetBind)
	//监听端口，默认8080
	e.Run()
}
*/
//json数据解析和绑定
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestJson(c *gin.Context) {
	var user User
	//将request的body中的数据，自动按照json格式解析到结构体
	//解析错误直接返回，至于要给客户端返回什么错误状态码有你决定
	err := c.ShouldBindJSON(&user)
	if err != nil {
		//相当于fmt.print后面跟着os.Exit(1)
		log.Fatal(err)
	}
	c.String(200, "user:%s", user)
}
func main() {
	//创建路由
	r := gin.Default()
	//绑定路由规则和处理函数
	r.POST("testJson", TestJson)
	//监听端口，默认8080
	r.Run()
}

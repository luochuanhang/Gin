package main

import "github.com/gin-gonic/gin"

func MyHandl(c *gin.Context) {
	//将给定的数据以json响应
	c.JSON(200, gin.H{
		"hello": "helloword",
	})
}
func Login(c *gin.Context) {
	//将给定的html页面返回
	c.HTML(200, "login.html", nil)
}

//gin.Context封装了request和response
func DoLogin(c *gin.Context) {
	//PostForm返回表单参数
	username := c.PostForm("username")
	password := c.PostForm("password")
	//将给定的html响应，将数据根据map的值写入html
	c.HTML(200, "hello.html", gin.H{
		"username": username,
		"password": password,
	})
}
func main() {
	//创建路由
	r := gin.Default()
	//gin支持加载HTML模板, 然后根据模板参数进行配置并返回相应的数据
	//LoadHTMLGlob()方法加载模板文件
	r.LoadHTMLGlob("templates/*")
	//绑定路由规则和执行的函数
	r.GET("/login", Login)
	r.POST("/login", DoLogin)
	//监听端口，默认8080
	r.Run()
}

package main

import "github.com/gin-gonic/gin"

func Regsiter(c *gin.Context) {
	//postform返回给定参数的值
	username := c.PostForm("username")
	password := c.PostForm("password")
	//postformarray返回给定参数多选框的值，是一个字符串切片
	hobby := c.PostFormArray("hobby")
	gender := c.PostForm("gender")
	city := c.PostForm("city")
	c.String(200, "username:%s,Password:%s, hobby:%s, gender:%s, city:%s", username, password, hobby, gender, city)

}
func GoRegsiter(c *gin.Context) {
	//将给定的html响应
	c.HTML(200, "register.html", nil)
}
func main() {
	//创建路由
	r := gin.Default()
	//加载模板文件
	r.LoadHTMLGlob("4.ginform表单处理/temolates/*")
	//绑定路由规则和处理函数
	r.GET("/register", GoRegsiter)
	r.POST("/register", Regsiter)
	//监听端口，默认8080
	r.Run()
}

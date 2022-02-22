package main

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	//将给定的页面返回
	c.HTML(200, "login.html", nil)
}
func main() {
	//创建路由
	r := gin.Default()
	//加载模板文件
	r.LoadHTMLGlob("6.gin访问静态文件集成bootStrap框架/templates/*")
	//加载静态文件
	r.Static("/assets", "./assets")
	//绑定路由规则和处理函数
	r.GET("/login", Login)
	//监听端口，默认8080
	r.Run()
}

package main

import "github.com/gin-gonic/gin"

func F1(c *gin.Context) {}
func F2(c *gin.Context) {
	c.String(200, "blog")
}
func F3(c *gin.Context) {}
func F4(c *gin.Context) {
	c.String(200, "video")
}
func main() {
	//创建路由
	r := gin.Default()
	//创建路由分组
	v1 := r.Group("/blog")
	{ //绑定路由规则和处理函数
		v1.POST("/list", F1)
		v1.GET("/add", F2)
	}
	v2 := r.Group("/video")
	{
		v2.POST("list", F3)
		v2.POST("/add", F4)
	}
	//监听端口
	r.Run()
}

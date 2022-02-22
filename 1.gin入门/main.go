package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//gin.Context封装了request和response
func hello(c *gin.Context) {
	//将给定的字符串写入响应体
	c.String(http.StatusOK, "helloword")
}
func main() {
	//创建gin路由
	//默认创建了2个中间件Logger(),Recovery()
	r := gin.Default()
	//绑定路由规则，执行的函数
	r.GET("/", hello)
	//监听端口默认8080端口
	r.Run()
}

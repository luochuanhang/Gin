package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//自定义中间件
//gin中间件必须是一个 gin.HandlerFunc 类型
func Middle(c *gin.Context) {
	fmt.Printf("自定义中间件")
}

func TestMW(c *gin.Context) {
	//将给定的数据通过字符串返回
	c.Next()
	c.String(200, "测试")
}
func main() {
	//创建路由默认使用了2个中间件loogger和recovery
	//r := gin.Default()
	r := gin.New() //创建路由不使用默认中间件
	//全局中间件,所以请求都会通过此中间件
	r.Use(Middle)
	//绑定路由规则和处理函数
	r.GET("/testMW", TestMW)
	//监听端口默认8080
	r.Run()
}

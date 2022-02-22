package main

import "github.com/gin-gonic/gin"

func Handler(c *gin.Context) {
	//获得cookie
	s, err := c.Cookie("username")
	if err != nil {
		s = "lch"
		//给客户端设置cookie
		// maxAge int, 单位为秒  多少秒后失效
		// path,cookie所在目录
		// domain string,域名
		// secure 是否智能通过https访问
		// httpOnly bool  是否允许别人通过js获取自己的cookie
		c.SetCookie("username", s, 60*60, "/", "localhost", false, true)
	}
	c.String(200, "测试cookie")
}
func main() {
	//创建路由
	r := gin.Default()
	//绑定路由规则和函数
	r.GET("test", Handler)
	//监听端口
	r.Run()
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Middle() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取客户端cookie并校验
		//如果有数据
		if s, err := c.Cookie("luo"); err == nil {
			//判断密码对不对
			if s == "123" {
				//执行后续的函数
				c.Next()
				return
			}
		}
		//返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		//验证不通过，不在调用后续函数
		c.Abort()
		return
	}
}
func Home(c *gin.Context) {
	// maxAge int, 单位为秒  多少秒后失效
	// path,cookie所在目录
	// domain string,域名
	// secure 是否智能通过https访问
	// httpOnly bool  是否允许别人通过js获取自己的cookie
	c.JSON(200, gin.H{"data": "home"})
}
func Login(c *gin.Context) {
	c.SetCookie("luo", "123", 60, "/", "localhost", false, true)
	c.String(200, "login success!")
}
func main() {
	//创建路由
	r := gin.Default()
	//绑定路由和处理函数
	r.GET("home", Middle(), Home)
	r.GET("login", Login)
	//监听端口
	r.Run()
}

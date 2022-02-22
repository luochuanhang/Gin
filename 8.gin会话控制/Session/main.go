package main

import (
	//导入session包
	"github.com/gin-contrib/sessions"
	//导入session存储引擎
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	//初始化session对象
	session := sessions.Default(c)
	//通过session.Get读取session的值
	//session是键值对格式数据，因此需要通过key查询数据
	if session.Get("hello") != "world" {
		//设置session数据
		session.Set("hello", "world")
		//删除session数据
		session.Delete("tizi123")
		//保存session数据
		session.Save()
		//删除整个session
		//session.Clear
	}
	c.JSON(200, gin.H{"hello": session.Get("hello")})
}
func main() {
	//创建路由
	r := gin.Default()
	//创建基于cookie的存储引擎，secret参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret"))
	//设置session中间件，参数mysession指的是session的名字也是cookie的名字
	//store是在之前创建的存储引擎，我们也可以替换成其他引擎
	r.Use(sessions.Sessions("mysession", store))
	//绑定路由规则和处理函数
	r.GET("/hello", hello)
	//监听端口
	r.Run()
}

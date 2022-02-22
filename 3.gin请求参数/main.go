package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//表单参数
func DoLogin(c *gin.Context) {
	//表单参数可以通过postform()方法获得
	username := c.PostForm("username")
	//表单参数可以通过defaultpostform()方法获得,没有返回默认值
	password := c.DefaultPostForm("password", "123")
	c.HTML(200, "hello.html", gin.H{
		"username": username,
		"password": password,
	})
}

//URL参数
func TestQueryString(c *gin.Context) {
	//url参数可以通过Query()和defaultQuery()方法返回
	//Query()若参数不存在，返回空串
	username := c.Query("username")
	//DefaultQuery()若参数不存在，返回默认值
	site := c.DefaultQuery("site", "www.baidu.com")
	//将给定的字符串写入响应体
	c.String(http.StatusOK, "username:%s,site:%s", username, site)
}

//路径参数
func TestPathParam(c *gin.Context) {
	//返回url参数的值
	s := c.Param("username")
	c.String(200, "Username:%s", s)
}

func main() {
	//创建路由
	r := gin.Default()
	//绑定路由规则和执行函数
	r.GET("/testQueryString", TestQueryString)
	r.POST("/login", DoLogin)
	r.GET("/hello/:username", TestPathParam)
	//监听端口，默认8080
	r.Run()

}

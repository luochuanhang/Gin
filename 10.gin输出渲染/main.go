package main

import "github.com/gin-gonic/gin"

func TestJson(c *gin.Context) {
	//将JSON将给定的结构作为XML序列化到响应体中，它还将Content-Type设置为“application/json”。
	//引入 gin.H 这个东西可以简化生成 json 的方式
	c.JSON(200, gin.H{
		"name": "多课网",
		"site": "www.duoke360.com",
	})
}

func TestXML(c *gin.Context) {
	//XML将给定的结构作为XML序列化到响应体中。它还将Content-Type设置为“application/xml”。
	c.XML(200, gin.H{
		"name": "多课网",
		"site": "www.duoke360.com",
	})
}

func TestHtml(c *gin.Context) {
	//HTML呈现由文件名指定的HTTP模板。它还更新了HTTP代码，并将Content-Type设置为“text/html”。
	c.HTML(200, "login.html", nil)
}

func TestString(c *gin.Context) {
	//第一个参数http状态码
	//第二个参数返回结果，支持类似Sprintf函数一样的字符串格式定义
	//第三个参数任意个format参数定义的字符串格式参数
	//String将给定字符串写入 response body
	c.String(200, "多课网，老郭讲golang")
}

func main() {
	//创建路由
	e := gin.Default()
	//绑定路由规则和处理函数
	e.GET("/test_json", TestJson)
	e.GET("/test_xml", TestXML)
	e.LoadHTMLGlob("10.gin输出渲染/templates/*")
	e.GET("/test_html", TestHtml)
	e.GET("/test_string", TestString)
	//监听端口
	e.Run()
}

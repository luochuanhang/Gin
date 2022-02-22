package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	// 获取上传文件，返回的是multipart.FileHeader对象，代表一个文件，里面包含了文件名之类的详细信息
	// file是表单字段名字
	file, _ := c.FormFile("file")
	// 打印上传的文件名
	log.Println(file.Filename)

	// 上传文件到项目根目录，使用原文件名
	c.SaveUploadedFile(file, file.Filename)
	//将字符串响应给浏览器
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func GoUpload(c *gin.Context) {
	//将给定的html响应给浏览器
	c.HTML(200, "upload.html", nil)
}

func main() {
	//创建路由
	router := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	router.MaxMultipartMemory = 8 << 20
	//加载模板
	router.LoadHTMLGlob("11.Gin文件上传/templates/*")
	//绑定路由规则和处理函数
	router.GET("/upload", GoUpload)
	router.POST("/upload", Upload)
	//监听端口
	router.Run(":8080")
}

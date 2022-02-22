package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	//创建路由
	r := gin.Default()
	//路由组使用gin.BasicAuth中间件
	//创建一个新的路由组
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		//存储账号和密码
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	//绑定路由规则的处理函数
	authorized.GET("/secrets", func(c *gin.Context) {
		//获取用户，它是由basicAuth中间件设置的
		//MustGet返回给定键的值，如果它存在，否则它会恐慌。
		//AuthUserKey是基本认证中用户凭据的cookie名称。
		user := c.MustGet(gin.AuthUserKey).(string)
		fmt.Println(user)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})
	//监听端口，默认8080
	r.Run()

}

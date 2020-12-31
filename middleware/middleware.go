/*
* @Author: zh
* @Email: zhangheng9394@163.com
* @Date: 2020/12/8 15:35
 */
package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TokenMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		header := context.Request.Header
		fmt.Println("-------- middle ------- ", header)
	}
}

func LogMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("login middle ware")
		//log.Println(context.Request.Host, context.Request.URL)
		//log.Println("Host: ", context.Request.Host, "Api:", context.Request.URL, "User-Agent: ", context.Request.Header.Get("User-Agent"))
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

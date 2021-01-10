/*
 * @Author: Zh
 * @Email: zhangheng9394@163.com
 * @Project: gin_web
 * @Package: views
 * @File: view_gorm.go
 * @Date: 2020/12/24 15:53
 */
package views

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "github.com/swaggo/swag"
)

func TemIndex(c *gin.Context) {



	c.HTML(http.StatusOK,"index.html", gin.H{
		"status": 1,
		"msg":    "success",
	})

}

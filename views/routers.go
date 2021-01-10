/*
 * @Author: Zh
 * @Email: zhangheng9394@163.com
 * @Project: gin_web
 * @Package: views
 * @File: routers.go
 * @Date: 2020/12/23 17:07
 */
package views

import (
	"fmt"
	_ "gin_web/docs"
	"gin_web/middleware"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/url"
)

func Router() *gin.Engine {

	// default 的日志格式是： [GIN] 2020/12/09 - 16:51:08 | 200 |     67.0031ms |       127.0.0.1 | GET      "/api/middle"
	//router := gin.Default()

	router := gin.New()
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	// type LogFormatter func(params LogFormatterParams) string 这里的LogFormatterParams是一个格式化日志参数的结构体
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		enEscapeUrl, _ := url.QueryUnescape(param.Path)
		statusCode := param.StatusCode
		colorCode := new(int)
		switch statusCode {
		case 200:
			*colorCode = 32 // 深绿:36, 绿：32
		case 404:
			*colorCode = 33 // 黄
		case 500:
			*colorCode = 31 // 红
		}
		// 你的自定义格式
		// Zh 2020-12-09 18:10:51	|	GET|	127.0.0.1|	/api/middle|	HTTP/1.1|	200|	22.0018ms|	PostmanRuntime/7.26.8|
		return fmt.Sprintf("\u001B[%dmZh %s\t|\t%s|\t%s|\t%s|\t%s|\t%d|\t%s|\t%s|\t%s \u001B[0m\n",
			*colorCode,
			//时间格式
			param.TimeStamp.Format("2006-01-02 15:04:05"),

			//http请求方式 get post等
			param.Method,

			//客户端IP
			param.ClientIP,

			//客户端请求的路径
			enEscapeUrl,

			//http请求协议版本
			param.Request.Proto,

			//http请求状态码
			statusCode,
			//耗时
			param.Latency,

			//http请求代理头
			param.Request.UserAgent(),

			//处理请求错误时设置错误消息
			param.ErrorMessage,
		)
	}))

	router.Use(gin.Recovery())

	router.Use(middleware.LogMiddleWare())
	router.Use(middleware.Cors()) // 解决跨越问题

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiGroup := router.Group("/api/", middleware.TokenMiddleWare())
	apiGroup.GET("/middle", Middle)

	api := router.Group("/api")

	get := api.Group("/get")

	{
		// 从URL获取get请求拼接的参数
		get.GET("/test", GetTest)

		// 从 path中获取参数，格式是：/api/get/second:/name/get_order_list，在请求中通过key=name，获取url这个位置的参数
		get.GET("/second/:urlName/test", GetUrlTest)

		// /api/get/all/tom/test?a=1111111111&b=2
		get.GET("/all/:urlName/test", GetUrlAll)
	}

	post := api.Group("/post")

	{
		post.POST("start", PostStart)

		post.POST("second", PostSecond)
	}

	gorms := api.Group("/gorms")

	{
		gorms.GET("/people/", GetPeople)
		gorms.GET("/people/:id", GetPerson)        // 通过id精确查询
		gorms.GET("/get/people/", GetPersonByName) // 通过名字模糊查询
		gorms.POST("/people", CreatePerson)
		gorms.PUT("/people/:id", UpdatePerson)
		gorms.DELETE("/people/:id", DeletePerson)
	}


	// 加载静态模版
	router.LoadHTMLGlob("templates/*")
	// 渲染引擎模板
	tem := api.Group("tem")
	{
		tem.GET("/index", TemIndex)
	}

	// 加载静态资源，路由通过 /static 开始匹配，后面的事文件的路径
	router.Static("/static", "./static")

	// 单个路由指向特定的静态资源
	router.StaticFile("/index/img", "./static/image/1.jpg")

	// 网站图标
	router.StaticFile("/favicon.ico", "./static/image/favicon.ico")



	return router
}

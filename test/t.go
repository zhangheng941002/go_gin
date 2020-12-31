package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Name *string

func TestMain()  {
	//gin 框架
	engin := gin.Default()
	engin.GET("/set", func(context *gin.Context) {
		FirstName := "测试名称"
		Name = &FirstName
		fmt.Println("------------ 设置变量 ---------------", *Name)

		fmt.Println(context.Request.URL, context.Request.Host)
	})
	engin.GET("/get", func(context *gin.Context) {
		fmt.Println("------------ 获取变量 ---------------", *Name)

		fmt.Println(context.Request.URL, context.Request.Host)
	})


	engin2 := gin.Default()
	engin2.GET("/set_again", func(context *gin.Context) {
		NameAgain := "名称更新"
		Name = &NameAgain
		fmt.Println("------------ 再次设置变量 ---------------", *Name)

		fmt.Println(context.Request.URL, context.Request.Host)
	})



	////浏览器访问 http://localhost:8080/api
	//mux := http.NewServeMux()
	//mux.HandleFunc("/set", myHandler)
	//go http.ListenAndServe(":8080", mux)
	//
	////浏览器访问 http://localhost:8081/api
	//mux1 := http.NewServeMux()
	//mux1.HandleFunc("/get", myHandler1)
	//go http.ListenAndServe(":8081", mux1)

	//浏览器访问 http://localhost:8082/api
	go http.ListenAndServe(":8082", engin)

	go http.ListenAndServe(":8083", engin)

	go http.ListenAndServe(":8084", engin2)

	fmt.Println("启动成功")

	//阻塞程序
	select {}

}

//func myHandler(res http.ResponseWriter, req *http.Request)  {
//	Name := "第一个设置了Name"
//	fmt.Println("------------ 第一个 ---------------", Name)
//
//	fmt.Println(req.URL, req.Host)
//}
//
//func myHandler1(res http.ResponseWriter, req *http.Request)  {
//	fmt.Println("------------ 第二个 ---------------")
//
//	fmt.Println(Name,"-----------------")
//
//	fmt.Println(req.URL, req.Host)
//}




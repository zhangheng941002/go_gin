/*
* @Author: zh
* @Email: zhangheng9394@163.com
* @Date: 2020/12/3 13:56
 */
package views

import (
	"fmt"
	"gin_web/models/test"
	"gin_web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Middle(c *gin.Context) {
	fmt.Println("111111111111111111")
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"msg":    "success",
	})
}

func Start(c *gin.Context) {
	fmt.Println("--------------par---", c.Params)
	fmt.Println("--------------par---", c.Params)
	for _, d := range c.Params {
		fmt.Println(d)
		fmt.Println(d.Key)
		fmt.Println(d.Value)
	}
	fmt.Println("------------------------------------------")
	fmt.Println(c.DefaultQuery("a", "aaa")) // 请求加入的参数
	fmt.Println("------------------------------------------")
	fmt.Println("--------------req---", c.Request)
	fmt.Println("--------------req---", c.Request.Body)
	fmt.Println("--------------req---", c.Request.Host)
	fmt.Println("--------------req---", c.Request.Method)
	fmt.Println("--------------req---", c.Request.Header)
	fmt.Println("--------------req---", c.Request.Header.Get("Accept"))
	fmt.Println("--------------req---", c.Request.Form)
	fmt.Println("--------------req---", c.Param("name")) // url 中定义的字段
	fmt.Println("--------------req---", c.Param("names"))

	/*
		post 请求获取body
	*/

	var user test.User
	codeAddress := new(int)
	//err := c.Bind(&user)
	err := c.ShouldBind(&user)
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Println(err)
	fmt.Println(user)
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	if err != nil {
		*codeAddress = http.StatusInternalServerError
	} else {
		*codeAddress = http.StatusOK
	}
	//c.JSON(*codeAddress, gin.H{
	c.JSON(*codeAddress, gin.H{
		"message": "pong",
	})
}

func GetTest(c *gin.Context) {
	params := c.Params
	fmt.Println(params)

	// url get 取参数  ?name=张三&page=1

	// DefaultQuery 取参数，需要设置取不到字段的默认值
	name := c.DefaultQuery("name", "nil")
	age := c.DefaultQuery("age", "nil")

	fmt.Println("--- name ---", name)
	fmt.Println("--- age ---", age)

	// GetQuery 返回 两个字段，第一个是返回值，第二个是是否获取到，有该字段返回时 true，没有返回 false
	page, pageStatus := c.GetQuery("page")

	fmt.Println("--- page ---", page)
	fmt.Println("--- pageStatus ---", pageStatus)

	_name := c.Query("name")
	fmt.Println("---------- query ------------", _name)

	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"msg":    "success",
	})
}

func GetUrlTest(c *gin.Context) {

	// 从url中获取参数

	// 方法一：Params，得到的是[]
	params := c.Params
	fmt.Println(params)

	// 方法二：param,通过key，获取指定的参数
	urlName := c.Param("urlName")
	urlNames := c.Param("urlNames")
	fmt.Println(urlName)
	fmt.Println(urlNames)

	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"msg":    "success",
	})
}

// @Summary GET请求包好所有获取参数
// @Description GET ALL
// @Tags GET请求测试
// @Accept json
// @Produce json
// @Param urlName path string true "人名" default(tom)
// @Param  gender query string true  "性别" Enums(男, 女, 其他)
// @Param  age query string false "年龄"
// @Success 200 {object} models.ResData {"status": 1,"msg":"success"}
// @Failure 500 {object} models.ResData {"status": 0,"msg":"success"}
// @Router /get/all/{urlName}/test [get]
func GetUrlAll(c *gin.Context) {

	fmt.Println("-------------- path parma -----------", c.Param("urlName")) // url 中定义的字段
	fmt.Println("-------------- path parma -----------", c.Param("names"))

	fmt.Println("------------------- DefaultQuery -----------------------")
	fmt.Println(c.DefaultQuery("who", "aaa")) // 请求加入的参数
	fmt.Println("------------------ DefaultQuery -----------------------")

	fmt.Println("============== query =================")
	fmt.Println(c.Query("age"))
	fmt.Println("============== query =================")

	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"msg":    "success",
	})
}


// @Summary POST请求
// @Description post
// @Tags GET请求测试
// @Accept json
// @Produce json
// @Param user body test.UserInfo true "用户信息"
// @Success 200 {object} models.ResData {"status": 1,"msg":"success"}
// @Failure 500 {object} models.ResData {"status": 0,"msg":"success"}
// @Router /post/start [post]
func PostStart(c *gin.Context) {

	/*
		post 请求获取body
		接收所有的介绍的数据
	*/
	var json map[string]interface{}
	//if err := c.ShouldBind(&json); err != nil {
	//	c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
	//	return
	//}
	err := c.ShouldBind(&json)
	code := utils.ResponseHttpCode(err)
	fmt.Println(json)
	//c.JSON(code, gin.H{"msg": fmt.Sprintf("username is %s", json["username"])})
	c.JSON(code, json)

}


// @Summary POST请求struct
// @Description 通过声明的 struct 来接收绑定的json数据
// @Description | 项目 | 价格 | 数量 |
// @Description | :-------- | --------:| :--: |
// @Description | iPhone | 6000 元 | 5 |
// @Description | iPad | 3800 元 | 12 |
// @Description | iMac | 10000 元 | 234 |
// @Tags GET请求测试
// @Accept json
// @Produce json
// @Param user body test.User true "用户注册信息"
// @Success 200 {object} models.ResData "成功"
// @Failure 500 {object} models.ResData "失败"
// @Router /post/second [post]
func PostSecond(c *gin.Context) {

	/*
		post 请求获取body
		通过声明的 struct 来接收绑定的json数据
	*/

	var user test.User
	//err := c.Bind(&user)
	err := c.ShouldBind(&user)
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Println(err)
	fmt.Printf("%+v", user)
	fmt.Println()
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")

	code := utils.ResponseHttpCode(err)

	data := utils.ResponseData(err)

	//var result []map[string]interface{}
	//for _, _data := range user.Data{
	//	fmt.Println(">>>>>>>>>>>>>", _data)
	//result = append(result, map[string]interface{}{
	//	"Name": _data["Username"],
	//	"Age": _data["Age"],
	//
	//})
	//	}
	data.Result = user.Data
	//c.JSON(*codeAddress, gin.H{
	c.JSON(code, data)
}

/*
* @Author: zh
* @Email: zhangheng9394@163.com
* @Date: 2020/12/1 11:31
 */
package main

import (
	"fmt"
	"gin_web/models"
	_ "gin_web/models"
	"gin_web/utils"
	"gin_web/views"
	_ "gopkg.in/ini.v1"
	"net/http"
	"time"
)

func main() {

	//// 方法一
	//port := flag.String("port", "8080", "port")
	//host := flag.String("host", "127.0.0.1", "host")
	//flag.Parse()
	//address := *host + ":" + *port
	//_ = router.Run(address) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	//// 运行多个服务
	//go http.ListenAndServe("0.0.0.0:8080", router)
	//
	//go http.ListenAndServe(":8081", router)

	timeStr := time.Now().Format("2006-01-02 15:04:05")  // 时间格式化
	router := views.Router()
	port := utils.Config.WEB.PORT
	host := utils.Config.WEB.HOST
	for _, i := range port{
		addr := host+":"+i
		go http.ListenAndServe(addr, router)
		fmt.Println(timeStr, "服务启动成功, host:",host, " port:",i)
		//break

	}
	// 修改某个值然后进行保存
	utils.Config.ConfigInfo.ConfigRouter.Section("RUNTIME").Key("LastRunTime").SetValue(timeStr)
	_ = utils.Config.ConfigInfo.ConfigRouter.SaveTo(utils.Config.ConfigInfo.ConfigPath)

	utils.Config.DbRouter.AutoMigrate(&models.Person{}, &models.User{})
	//阻塞程序
	select {}
}

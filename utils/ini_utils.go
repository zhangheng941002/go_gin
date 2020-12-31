/*
 * @Author: Zh
 * @Email: zhangheng9394@163.com
 * @Project: gin_web
 * @Package: utils
 * @File: ini_utils.go
 * @Date: 2020/12/23 10:55
 */
package utils

import (
	"fmt"
	"gin_web/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/ini.v1"
	"strconv"
)


var _conf *ini.File
var _err error
var config models.ConfigInfo
// 初始化 ini对象

var Config models.ConfigMap

func init() {
	// 加载配置文件
	_configPath := "config/app.ini"
	_conf, _err = ini.Load(_configPath)
	//_conf, _err = ini.Load("config/app.ini")
	if _err != nil {
		panic("加载配置文件初始化失败，err=" + _err.Error())
	}

	config.ConfigPath = _configPath
	config.ConfigRouter = _conf
	Config.ConfigInfo = &config

	// 获取web信息，host和 port
	web := new(models.WebServer)
	_ = _conf.Section("web").MapTo(web)
	Config.WEB = web

	// 获取加载数据库信息
	db := new(models.DataBase)
	_ = _conf.Section("mysql").MapTo(db)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", db.USER, db.PASSWORD, db.HOST, db.PORT, db.DBNAME, db.TIMEOUT)
	Config.DbRouter, _err = gorm.Open("mysql", dsn)
	if _err != nil {
		panic("连接数据库失败, error=" + _err.Error() + "  dsn=" + dsn)
	}

	//设置数据库连接池参数
	SetMaxOpenCons, _ := strconv.Atoi(db.SetMaxOpenCons)
	SetMaxIdleCons, _ := strconv.Atoi(db.SetMaxIdleCons)
	Config.DbRouter.DB().SetMaxOpenConns(SetMaxOpenCons) //设置数据库连接池最大连接数, int
	Config.DbRouter.DB().SetMaxIdleConns(SetMaxIdleCons)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。 int

}


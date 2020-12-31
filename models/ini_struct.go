/*
 * @Author: Zh
 * @Email: zhangheng9394@163.com
 * @Project: gin_web
 * @Package: models
 * @File: ini_struct.go
 * @Date: 2020/12/22 16:12
 */
package models

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/ini.v1"
)

type DataBase struct {
	PORT           string
	HOST           string
	DBNAME         string
	USER           string
	PASSWORD       string
	TIMEOUT        string
	SetMaxOpenCons string
	SetMaxIdleCons string
}

type WebServer struct {
	HOST string
	PORT []string
}

type ConfigInfo struct {
	ConfigPath  string
	ConfigRouter *ini.File
}


type ConfigMap struct {
	DbRouter *gorm.DB
	WEB      *WebServer
	ConfigInfo *ConfigInfo

}
/*
 * @Author: Zh
 * @Email: zhangheng9394@163.com
 * @Project: gin_web
 * @Package: models
 * @File: gorm_model.go
 * @Date: 2020/12/24 16:01
 */
package models

import "time"

type InitModel struct {
	//CreateDate time.Time `gorm: default: "2020-12-25 06:06:06" json:"create_date"`
	//UpdateDate time.Time `gorm: default: "2020-12-25 06:06:06" json:"update_date"`
	Status     int       `gorm:"default:1 "json:"status"`
}

type User struct {
	Name     string `json:"name"`
	PassWord string `json:"password"`
	InitModel
}

// gorm.Model 的定义
type Model struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time `gorm: default: time.Now()`
	UpdatedAt time.Time  `gorm: default: time.Now()`
}

type Person struct {
	ID uint `gorm:"primaryKey" json:"id"`
	//gorm.Model  // 默认有四个字段
	UserInfo  JSON   `gorm:"type:json;comment:'个人信息'" json:"user_info"` // json字段，对应数据库里的字段类型也是json
	User      User   `gorm:"embedded" json:"user"`                      // 会将引用的struct中的字段拆分给继承过来，创建和查询又是根据struct来的
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Tel       string `json:"tel"`
	TelPhone  string `json:"tel_phone"`
	City      string `json:"city"`
}

/*
 * @Author: Zh
 * @Email: zhangheng9394@163.com
 * @Project: gin_web
 * @Package: test
 * @File: test.go
 * @Date: 2020/12/18 17:58
 */
package main

// only need mysql OR sqlite
// both are included here for reference
import (
	"bytes"
	"database/sql/driver"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type User struct {
	Name string `json:"name"`
	Sd   string `json:"sd"`
}

type Person struct {
	ID uint `json:"id"`
	//gorm.Model
	//User      JSON   `gorm:"type:json;comment:'个人信息'" json:"user"`  // json字段，对应数据库里的字段类型也是json
	User      User   `gorm:"embedded" json:"user"`  // 会将引用的struct中的字段拆分给继承过来，创建和查询又是根据struct来的
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name" `
	Tel       string `json:"tel"`
	TelPhone  string `json:"tel_phone" `
	City      string `json:"city" `
}

type JSON []byte

func (j JSON) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(j), nil
}
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("Invalid Scan Source")
	}
	*j = append((*j)[:], s...)
	return nil
}
func (m JSON) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}
func (m *JSON) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("null point exception")
	}
	*m = append((*m)[:], data...)
	return nil
}

func (j JSON) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}
func (j JSON) Equals(j1 JSON) bool {
	return bytes.Equal([]byte(j), []byte(j1))
}

func StructToJSON(value interface{}) JSON {
	bts, _ := jsoniter.Marshal(value)
	return bts
}

//// Value 实现方法
//func (p User) Value() (driver.Value, error) {
//	return json.Marshal(p)
//}
//
//// Scan 实现方法
//func (p User) Scan(input interface{}) error {
//	return json.Unmarshal(input.([]byte), &p)
//}

func main() {

	// NOTE: See we’re using = to assign the global var
	// instead of := which would assign it only in this function
	//db, err = gorm.Open("sqlite3", "./gorm.db")
	db, _ = gorm.Open("mysql", "root:Yunkuan1!@tcp(172.31.0.8:3306)/test_portal?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Person{}) // migrate 仅支持创建表、增加表中没有的字段和索引，它并不支持改变已有的字段类型或删除未被使用的字段

	r := gin.Default()
	r.GET("/people/", GetPeople)
	r.GET("/people/:id", GetPerson)
	r.POST("/people", CreatePerson)
	r.PUT("/people/:id", UpdatePerson)
	r.DELETE("/people/:id", DeletePerson)

	_ = r.Run(":8080")
}

func DeletePerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person
	d := db.Where("id = ?", id).Delete(&person)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func UpdatePerson(c *gin.Context) {

	var person Person
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	_ = c.BindJSON(&person)

	db.Save(&person)
	c.JSON(200, person)

}

func CreatePerson(c *gin.Context) {

	var person Person
	_ = c.Bind(&person)
	fmt.Println(person, "-----------------", person.FirstName, person.Tel)
	// 	db.Create(&Person{FirstName:"DDD",LastName:"sss"})
	db.Create(&person)
	c.JSON(200, person)
}

func GetPerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person []Person
	//if err := db.Where("id = ?", id).First(&person).Error; err != nil {
	//if err := db.Where("first_name = ?", id).First(&person).Error; err != nil {
	if err := db.Where("first_name = ?", id).Find(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, person)
	}
}

func GetPeople(c *gin.Context) {
	var people []Person
	if err := db.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, people)
	}

}

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
	"fmt"
	"gin_web/models"
	"gin_web/utils"
	"github.com/gin-gonic/gin"
)


func DeletePerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person models.Person
	d := utils.Config.DbRouter.Where("id = ?", id).Delete(&person)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func UpdatePerson(c *gin.Context) {

	var person models.Person
	id := c.Params.ByName("id")

	if err := utils.Config.DbRouter.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	_ = c.BindJSON(&person)

	utils.Config.DbRouter.Save(&person)
	c.JSON(200, person)

}

func CreatePerson(c *gin.Context) {

	var person models.Person
	_ = c.ShouldBindJSON(&person)
	fmt.Printf("%+v", person,"==================")
	// 	db.Create(&Person{FirstName:"DDD",LastName:"sss"})
	utils.Config.DbRouter.Create(&person)
	c.JSON(200, person)
}

func GetPerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person []models.Person
	//if err := db.Where("id = ?", id).First(&person).Error; err != nil {
	//if err := db.Where("first_name = ?", id).First(&person).Error; err != nil {
	if err := utils.Config.DbRouter.Where("id = ?", id).Find(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, person)
	}
}

// 模糊查询
func GetPersonByName(c *gin.Context) {
	_name := c.Query("name")
	fmt.Println("------------------", _name)
	queryName := "%"+ _name + "%"
	var person []models.Person

	if err := utils.Config.DbRouter.Where("first_name like ?", queryName).Find(&person).Error; err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, person)
	}
}

func GetPeople(c *gin.Context) {
	var people []models.Person
	if err := utils.Config.DbRouter.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, people)
	}

}


package controllers

import (
	"helloapp/internal/initializers"
	"helloapp/models"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	//request body data
	var body struct {
		Name     string
		Phone    string
		Password string
	}
	c.Bind(&body)
	//post data
	user := models.User{Name: body.Name, Phone: body.Phone, Password: body.Password}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"message": "user register success",
		"user":    user,
	})

}

// get the all user
func GetUser(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)

	c.JSON(200, gin.H{
		"message": "user are:",
		"users":   users,
	})

}

// get specific user
func GetSpecificUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	initializers.DB.First(&user, id)
	c.JSON(200, gin.H{
		"message": "user are:",
		"user":    user,
	})
}

// update user
func UpdateUser(c *gin.Context) {
	//get id from url
	id := c.Param("id")
	//get user from data base
	var user models.User
	initializers.DB.First(&user, id)

	//request body data
	var body struct {
		Name     string
		Phone    string
		Password string
	}
	c.Bind(&body)
	initializers.DB.Model(&user).Updates(models.User{
		Name:     body.Name,
		Phone:    body.Phone,
		Password: body.Password,
	})

	c.JSON(200, gin.H{
		"message": "updated user info is:",
		"user":    user,
	})

}

func Delete(c *gin.Context) {
	//get user from url
	id := c.Param("id")
	//get user from data base
	var user models.User
	initializers.DB.First(&user, id)
	//delete user from data base
	initializers.DB.Delete(&user)
	c.JSON(200, gin.H{
		"message": "user deleted",
	})
}

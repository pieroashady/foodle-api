package controllers

import (
	"foodle-api/initializers"
	"foodle-api/models"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func FindAllUser(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)
	c.JSON(200, gin.H{
		"data": users,
	})
}

func CreateUser(c *gin.Context) {
	user := models.User{Fullname: "Adit", Email: "adit", Password: "Adit"}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}

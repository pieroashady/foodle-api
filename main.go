package main

import (
	"foodle-api/controllers"
	"foodle-api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
    r := gin.Default()

	r.GET("/users", controllers.FindAllUser)
	r.POST("/users", controllers.CreateUser)

	r.Run() 
}
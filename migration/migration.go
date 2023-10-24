package main

import (
	"fmt"
	"foodle-api/initializers"
	"foodle-api/models"
	"os"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	dbSync := os.Getenv("DB_SYNC")
	
	if dbSync == "true" {
		fmt.Println("Syncing database")
		initializers.DB.AutoMigrate(&models.User{})
		fmt.Println("Done Syncing database")
	}
}
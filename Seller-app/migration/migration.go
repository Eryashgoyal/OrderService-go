package main

import (
	"log"
	"orderService/initializers"
	"orderService/models"
)

func main() {

	initializers.ConnectToDB()

	err := initializers.DB.AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatal("Error in migration")
		
	}

	log.Println("Migration done successfully")
}
package main

import (
	"log"
	"authentication/router"
	"authentication/utils/database"
)

// @title Stock Broker Application
// @description api for Stock Broker using gin and gorm
// @version 1.0
// @BasePath /
// @host localhost:8080
func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database:%v", err)
	}
	log.Println("Database connection is initialized, Application starting...")
	router := router.SetupRouter(database)
	router.Run(":8080")
}

package main

import (
	"authentication/router"
	db "authentication/utils/database"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

// @title Your API Title
// @version 1.0
// @description This is a sample API for demonstration purposes.
// @host localhost:8080
// @BasePath /

func main() {
	db, err := db.SetupGorm()
	if err != nil {
		log.Fatalf("Error initializing database:%v", err)
	}
	router := router.SetupRouter(db)
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}

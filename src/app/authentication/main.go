package main

import (
	_ "authentication/docs"
	"authentication/router"
	"fmt"
)

// @title Stock Broker Application
// @description   api for Stock Broker using gin and gorm
// @version 2.0
// @host localhost:8080
// @securityDefinitions.apiKey JWT
// @in header
// @name token
func main() {
	apiRouter := router.InitializeRouter() // router initialize on port 8080
	fmt.Println("The Connection is open at port-8080")
	apiRouter.Run(":8080")
}

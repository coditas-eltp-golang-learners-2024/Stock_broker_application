package database

import (
	"authentication/constants"
	"authentication/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// databases connection creation (mysql)
func ConnectionWithDb() (*gorm.DB, error) {
	// Parse Details from Yaml to struct
	sqlCredentials := utils.ParseYAML()
	DatabaseURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		sqlCredentials.DBUsername,
		sqlCredentials.DBPassword,
		sqlCredentials.DBHost,
		sqlCredentials.DBPort,
		sqlCredentials.DBName)
	ConnectToDatabase, err := gorm.Open(mysql.Open(DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("%s: %s", constants.ErrOpenDatabaseConnection.Error(), err.Error())
	}
	// Ping connection (Status check)
	sqlDB, _ := ConnectToDatabase.DB()
	if sqlDB.Ping() != nil {
		log.Fatalf("%s: %s", constants.ErrDatabasePing.Error(), err.Error())
	}
	return ConnectToDatabase, nil
}

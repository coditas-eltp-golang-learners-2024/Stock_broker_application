package db

import (
	"authentication/utils"
	"fmt"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB


func InitDB() {
    // Load configuration
    sqlConfig := utils.LoadConfig()

    // Database connection string
    connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
        sqlConfig.Username,
        sqlConfig.Password,
        sqlConfig.Host,
        sqlConfig.Port,
        sqlConfig.DBName,
    )

    var err error
    DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }

    // Test the connection
    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatalf("Error getting underlying DB: %v", err)
    }
   // defer sqlDB.Close()

    err = sqlDB.Ping()
    if err != nil {
        log.Fatalf("Error pinging database: %v", err)
    }

    log.Println("Connected to database")
}

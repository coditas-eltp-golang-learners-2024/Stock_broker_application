package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"authentication/utils"
)

// InitDB initializes and returns a database connection.
func InitDB() (*gorm.DB, error) {
	// Load database configuration
	config, err := utils.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading database configuration: %w", err)
	}

	// Construct Data Source Name (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		config.Username, config.Password, config.Host, config.Port, config.DBName)

	// Open a database connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	// Ping the database to verify the connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error getting DB instance: %w", err)
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	log.Println("Connected to database")
	return db, nil
}

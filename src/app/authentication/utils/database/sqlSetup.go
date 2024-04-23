package db

import (
	"authentication/utils"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupGorm() (*gorm.DB, error) {
	database := utils.DbConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", database.Username, database.Password, database.Host, database.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

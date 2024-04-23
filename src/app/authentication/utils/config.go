package utils

import (
	"authentication/constants"
	"authentication/models"

	"github.com/spf13/viper"
)

func DbConfig() *models.DBConfig {

	viper.SetConfigName("database")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(constants.DatabaseYamlFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var dbConfig models.DBConfig

	err = viper.UnmarshalKey("dbHostName", &dbConfig.Host)
	if err != nil {
		panic(err)
	}

	err = viper.UnmarshalKey("port", &dbConfig.Port)
	if err != nil {
		panic(err)
	}

	err = viper.UnmarshalKey("dbName", &dbConfig.DBName)
	if err != nil {
		panic(err)
	}

	err = viper.UnmarshalKey("userName", &dbConfig.Username)
	if err != nil {
		panic(err)
	}

	err = viper.UnmarshalKey("password", &dbConfig.Password)
	if err != nil {
		panic(err)
	}

	return &dbConfig

}

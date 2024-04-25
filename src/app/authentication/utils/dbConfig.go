package utils

import (
	"authentication/constants"
	"stock_broker_application/src/models"

	"github.com/spf13/viper"
)

func PostgresConfig() *models.PostgresConfig {

	viper.SetConfigName("database")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(constants.DatabaseYamlFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var postgresConfig models.PostgresConfig

	err = viper.UnmarshalKey("Host", &postgresConfig.Host)
	if err != nil {
		panic(err)
	}

	err = viper.UnmarshalKey("Port", &postgresConfig.Port)
	if err != nil {
		panic(err)
	}

	err = viper.UnmarshalKey("DBName", &postgresConfig.DBName)
	if err != nil {
		panic(err)
	}

	err = viper.UnmarshalKey("User", &postgresConfig.User)
	if err != nil {
		panic(err)
	}

	err = viper.UnmarshalKey("Password", &postgresConfig.Password)
	if err != nil {
		panic(err)
	}

	return &postgresConfig

}

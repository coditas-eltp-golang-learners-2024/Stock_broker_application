package utils

import (
	"github.com/spf13/viper"
	"log"
	"authentication/models"
)

func LoadConfig() (models.DatabaseConfig, error) {
	// Set the configuration file path
	viper.SetConfigFile("resources/database.yml")

	// Read the configuration file into Viper
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
		return models.DatabaseConfig{}, err
	}

	// Unmarshal the configuration into a struct
	var config models.DatabaseConfig
	if err := viper.UnmarshalKey("database", &config); err != nil {
		log.Fatalf("Unable to decode config into struct: %s", err)
		return models.DatabaseConfig{}, err
	}

	return config, nil
}

package utils

import (
	"authentication/constants"
	"authentication/models"
	"github.com/spf13/viper"
	"log"
)

// Takes yaml filePath and return struct
func ParseYAML() models.DbDetails {
	viper.SetConfigFile("resources/database.yml")

	// Read the configuration file into Viper
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("%s: %s", constants.ErrReadConfigFile.Error(), err.Error())
	}

	var dbDetails models.DbDetails // Unmarshal the configuration into a struct
	if err := viper.UnmarshalKey("database", &dbDetails); err != nil {
		log.Fatalf("%s: %s", constants.ErrDecodeConfig.Error(), err.Error())
	}
	return dbDetails // Return the database details

}

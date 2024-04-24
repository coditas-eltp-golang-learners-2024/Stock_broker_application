package configs

import (
	"context"
	"os"
	"stock_broker_application/src/constants"
	"stock_broker_application/src/models"
	"strconv"
)

var applicationConfig *models.ApplicationConfig

func InitApplicationConfigs(ctx context.Context) error {
	// Get a Apis config from yaml
	applicationViperConfig, err := Get(constants.ApplicationConfig)
	if err != nil {
		return err
	}

	applicationConfig = &models.ApplicationConfig{
		SwaggerConfig: models.SwaggerConfig{
			SwaggerHost: applicationViperConfig.GetString(constants.SwaggerHostKey),
		},
		Token: models.Token{
			AccessTokenSecretKey:     applicationViperConfig.GetString(constants.AccessTokenSecretKey),
			RefreshTokenSecretKey:    applicationViperConfig.GetString(constants.RefreshTokenSecretKey),
			AccessTokenExpiryInDays:  applicationViperConfig.GetInt(constants.AccessTokenExpiryInDaysKey),
			RefreshTokenExpiryInDays: applicationViperConfig.GetInt(constants.RefreshTokenExpiryInDaysKey),
			EnableTokenCompression:   applicationViperConfig.GetBool(constants.EnableTokenCompressionKey),
		},
	}
	return nil
}

func GetApplicationConfig() *models.ApplicationConfig {
	return applicationConfig
}

func getEnvInt(name string, defaultValue int) int {
	valueStr := os.Getenv(name)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvBool(name string, defaultValue bool) bool {
	valueStr := os.Getenv(name)
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultValue
}

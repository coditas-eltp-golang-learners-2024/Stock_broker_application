package models

type ApplicationConfig struct {
	SwaggerConfig SwaggerConfig
	Token         Token
}

type Token struct {
	AccessTokenSecretKey     string
	RefreshTokenSecretKey    string
	AccessTokenExpiryInDays  int
	RefreshTokenExpiryInDays int
	SecretKey                string
	EnableTokenCompression   bool
}

type SwaggerConfig struct {
	SwaggerHost string
}

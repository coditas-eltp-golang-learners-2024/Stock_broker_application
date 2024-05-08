package authorization

import (
	"stock_broker_application/src/constants"
	genericModel "stock_broker_application/src/models"
	"stock_broker_application/src/utils/configs"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWTToken(tokenData genericModel.TokenData) (string, error) {
	secretKey := configs.GetApplicationConfig().Token.SecretKey

	// Create a new JWT claims instance
	claims := jwt.MapClaims{}

	// Set the token payload fields
	claims[constants.TokenPayload] = tokenData

	// Set the token expiration time
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(configs.GetApplicationConfig().Token.AccessTokenExpiryInDays)).Unix()

	// Create a new JWT token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

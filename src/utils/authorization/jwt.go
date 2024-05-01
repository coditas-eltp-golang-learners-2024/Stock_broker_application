package authorization

import (
	"github.com/dgrijalva/jwt-go"
	"stock_broker_application/src/models"
	"stock_broker_application/src/utils/configs"
	"time"
)

func GenerateJWTToken(username string) (string, error) {
	tokenData := models.Tokens{Username: username}

	secretKey := configs.GetApplicationConfig().Token.SecretKey

	// Set up JWT claims.
	claims := jwt.MapClaims{}
	claims["username"] = tokenData.Username
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(configs.GetApplicationConfig().Token.AccessTokenExpiryInDays)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

package authorization

import (
	"github.com/dgrijalva/jwt-go"
	"stock_broker_application/src/utils/configs"
	"time"
	"stock_broker_application/src/constants"
	genericModel "stock_broker_application/src/models"
)

func GenerateJWTToken(tokenData genericModel.TokenData) (string, error) {

	secretKey := configs.GetApplicationConfig().Token.SecretKey

	// Set up JWT claims.
	claims := jwt.MapClaims{}
	claims[constants.TokenPayload] = tokenData
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(configs.GetApplicationConfig().Token.AccessTokenExpiryInDays)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

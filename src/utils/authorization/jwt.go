package authorization

import (
	jwt "github.com/dgrijalva/jwt-go"
	"stock_broker_application/src/constants"
	genericModel "stock_broker_application/src/models"
	"stock_broker_application/src/utils/configs"
	"time"
)

func GenerateJWTToken(tokenData genericModel.TokenData) (string, error) {
	secretKey := configs.GetApplicationConfig().Token.AccessTokenSecretKey
	claims := jwt.MapClaims{}
	claims[constants.TokenPayload] = tokenData
	claims[constants.TokenExpiration] = time.Now().Add(24 * time.Hour * time.Duration(configs.GetApplicationConfig().Token.AccessTokenExpiryInDays)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return constants.EmptySpace, err
	}
	return tokenString, nil
}

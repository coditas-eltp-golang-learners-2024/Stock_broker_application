package headerCheck

import (
	"net/http"
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/models"
	"stock_broker_application/src/utils/configs"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		secretKey := configs.GetApplicationConfig().Token.AccessTokenSecretKey
		tokenString := ctx.GetHeader(genericConstants.Authorization)
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{genericConstants.GenericJSONErrorMessage: genericConstants.AuthTokenMissing})
			ctx.Abort()
			return
		}
		token, err := jwt.ParseWithClaims(tokenString, &models.TokenData{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{genericConstants.GenericJSONErrorMessage: genericConstants.FailedJWTValidation})
			ctx.Abort()
			return
		}
		if claims, ok := token.Claims.(*models.TokenData); ok && token.Valid {
			ctx.Set(genericConstants.Id, claims.UserId)
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{genericConstants.GenericJSONErrorMessage: genericConstants.InvalidJWT})
			ctx.Abort()
			return
		}
	}
}

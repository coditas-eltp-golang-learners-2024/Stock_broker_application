package headerCheck

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/models"
	"strings"
)

var SecretKey = []byte("secret_key")

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader(genericConstants.Authorization)
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{genericConstants.GenericErrorMessage: genericConstants.AuthTokenMissing})
			ctx.Abort()
			return
		}
		const bearerPrefix = genericConstants.Bearer
		if strings.HasPrefix(tokenString, bearerPrefix) {
			tokenString = tokenString[len(bearerPrefix):]
		}
		tokenString = strings.TrimSpace(tokenString)
		token, err := jwt.ParseWithClaims(tokenString, &models.TokenModel{}, func(token *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{genericConstants.GenericErrorMessage: genericConstants.FailedJWTValidation})
			ctx.Abort()
			return
		}
		if claims, ok := token.Claims.(*models.TokenModel); ok && token.Valid {
			ctx.Set(genericConstants.EmailId, claims.Email)
			ctx.Next()
		}else {
			ctx.JSON(http.StatusUnauthorized, gin.H{genericConstants.GenericErrorMessage: genericConstants.InvalidJWT})
			ctx.Abort()
			return
		}
	}
}

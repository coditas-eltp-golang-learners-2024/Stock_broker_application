package headerCheck

import (
	"net/http"
	"stock_broker_application/src/constants"
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/models"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
		// Check if the token starts with "Bearer "
		const bearerPrefix = genericConstants.Bearer
		if strings.HasPrefix(tokenString, bearerPrefix) {
			// Remove the "Bearer " prefix
			tokenString = tokenString[len(bearerPrefix):]
		}
		// Trim any leading or trailing whitespace from the token string
		tokenString = strings.TrimSpace(tokenString)
		token, err := jwt.ParseWithClaims(tokenString, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{genericConstants.GenericErrorMessage: constants.FailedJWTValidation})
			ctx.Abort()
			return
		}
		if claims, ok := token.Claims.(*models.Claim); ok && token.Valid {
			ctx.Set("email", claims.Email)
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{genericConstants.GenericErrorMessage: constants.InvalidJWT})
			ctx.Abort()
			return
		}
	}
}

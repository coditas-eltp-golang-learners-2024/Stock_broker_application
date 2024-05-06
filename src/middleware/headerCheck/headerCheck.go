package headerCheck

import (
	"log"
	"net/http"
	genericConstants "stock_broker_application/src/constants"
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

		token, tokenErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if tokenErr != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{genericConstants.GenericJSONErrorMessage: genericConstants.FailedJWTValidation})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{genericConstants.GenericJSONErrorMessage: genericConstants.FailedJWTValidation})
			return

		}

		log.Print(claims)

		// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 	return []byte(secretKey), nil
		// })
		// log.Print(token)
		// if err != nil {
		// 	ctx.JSON(http.StatusUnauthorized, gin.H{genericConstants.GenericJSONErrorMessage: genericConstants.FailedJWTValidation})
		// 	ctx.Abort()
		// 	return
		// }
		// if claims, ok := token.Claims.(*models.TokenModel); ok && token.Valid {
		// 	ctx.Set(genericConstants.Id, claims.UserID)
		// 	ctx.Next()
		// } else {
		// 	ctx.JSON(http.StatusUnauthorized, gin.H{genericConstants.GenericJSONErrorMessage: genericConstants.InvalidJWT})
		// 	ctx.Abort()
		// 	return
		// }
	}
}

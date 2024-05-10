package headerCheck

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/utils/configs"
	"strings"
	"time"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		secretKey := configs.GetApplicationConfig().Token.AccessTokenSecretKey
		tokenString := ctx.GetHeader(genericConstants.Authorization)
		if tokenString == genericConstants.EmptySpace {
			ctx.JSON(http.StatusUnauthorized, gin.H{genericConstants.GenericJSONErrorMessage: genericConstants.AuthTokenMissing})
			ctx.Abort()
			return
		}
		if !strings.HasPrefix(tokenString, genericConstants.Bearer) {
			ctx.JSON(http.StatusUnauthorized, gin.H{genericConstants.GenericJSONErrorMessage: genericConstants.JWTTokenBearerMissingError})
			ctx.Abort()
			return
		} else {
			tokenString = tokenString[len(genericConstants.Bearer):]
		}
		tokenString = strings.TrimSpace(tokenString)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{genericConstants.GenericJSONErrorMessage: genericConstants.FailedJWTValidation})
			ctx.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			exp := int64(claims[genericConstants.TokenExpiration].(float64))
			if exp < time.Now().Unix() {
				ctx.JSON(http.StatusUnauthorized, gin.H{genericConstants.GenericJSONErrorMessage: genericConstants.JWTTokenExpiredError})
				ctx.Abort()
				return
			}
			id, ok := claims[genericConstants.TokenPayload].(map[string]interface{})[genericConstants.Id].(float64)
			if !ok {
				ctx.JSON(http.StatusUnauthorized, gin.H{genericConstants.GenericJSONErrorMessage: genericConstants.ExtractPayloadError})
				ctx.Abort()
				return
			}
			ctx.Set(genericConstants.Id, uint16(id))
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{genericConstants.GenericJSONErrorMessage: genericConstants.InvalidJWT})
			ctx.Abort()
			return
		}
	}
}

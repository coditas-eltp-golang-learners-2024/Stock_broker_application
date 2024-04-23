package handlers

import (
	"authentication/constants"
	"authentication/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

// @Summary Authenticate requests using JWT token
// @Description Authenticate requests using JWT token
// @ID auth-middleware
// @Produce json
// @Success 200 {object} string "Token is valid"
// @Failure 401 {object} string "Unauthorized: missing or invalid token"
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{constants.ErrGeneric.Error(): constants.ErrMissingToken.Error()})
			ctx.Abort()
			return
		}
		// Check if the token starts with "Bearer "
		const bearerPrefix = "Bearer "
		if strings.HasPrefix(tokenString, bearerPrefix) {
			// Remove the "Bearer " prefix
			tokenString = tokenString[len(bearerPrefix):]
		}
		// Trim any leading or trailing whitespace from the token string
		tokenString = strings.TrimSpace(tokenString)
		token, err := jwt.ParseWithClaims(tokenString, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
			return constants.SecretKey, nil
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{constants.ErrGeneric.Error(): constants.ErrPrasingToken.Error()})
			ctx.Abort()
			return
		}
		if claims, ok := token.Claims.(*models.Claim); ok && token.Valid {
			ctx.Set("email", claims.Email)
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{constants.ErrGeneric.Error(): constants.ErrInvalidToken.Error()})
			ctx.Abort()
			return
		}

	}
}

package authorization

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("secret_key")

type Token struct {
	Email     string    `json:"email"`
	ExpiresAt time.Time `json:"expires_at"`
}

func GenerateJWTToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      email,
		"expires_at": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, error := token.SignedString(jwtKey)
	if error != nil {
		return "", error
	}
	return tokenString, nil
}

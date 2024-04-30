package authorization

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("secret_key")

type Token struct {
	UserName  string    `json:"username"`
	ExpiresAt time.Time `json:"expires_at"`
}

func GenerateJWTToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":   username,
		"expires_at": time.Now().Add(time.Hour * 720).Unix(),
	})
	tokenString, error := token.SignedString(jwtKey)
	if error != nil {
		return "", error
	}
	return tokenString, nil
}

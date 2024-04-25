package authorization
import(
	"time"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key_here")

// GenerateJWTToken generates a JWT token for the given email
func GenerateJWTToken(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours

	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	 
	return tokenString, nil	
}

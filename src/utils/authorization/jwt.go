package authorization
import(
	"time"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

func GenerateJWTToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), 
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
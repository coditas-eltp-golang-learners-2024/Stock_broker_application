package business

import (
	"github.com/dgrijalva/jwt-go"
	"authentication/constants"
	"authentication/models"
	"authentication/repositories"
	"time"
)

type OTPService struct {
	UserRepository repositories.CustomerRepository
}

var jwtKey = []byte("your_secret_key_here")

func NewOTPService(userRepository repositories.CustomerRepository) *OTPService {
	return &OTPService{
		UserRepository: userRepository,
	}
}
func (otpService *OTPService) OtpVerification(otpData models.OTPValidationRequest) error {
	if !otpService.UserRepository.CheckOtp(otpData.Email, otpData.OTP) {
		return constants.ErrOtpVerification
	}
	return nil
}

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

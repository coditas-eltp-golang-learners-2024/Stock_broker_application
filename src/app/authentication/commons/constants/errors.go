package constants

import "errors"

var (
	ErrOtpVerification = errors.New("OTP verification failed")
	ErrGenToken        = errors.New("failed to generate token")
)

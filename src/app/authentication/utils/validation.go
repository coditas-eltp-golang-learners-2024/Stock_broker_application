package utils

import (
	"authentication/commons/constants"
	"errors"
	"regexp"
)

func ValidatePassword(password string) error {
	// Minimum length check
	if len(password) < 8 {
		return errors.New(constants.ErrPasswordTooShort)
	}

	// Uppercase, lowercase, digit, and special character check using regular expressions
	var (
		hasUpper   = regexp.MustCompile(`[A-Z]`).MatchString
		hasLower   = regexp.MustCompile(`[a-z]`).MatchString
		hasDigit   = regexp.MustCompile(`[0-9]`).MatchString
		hasSpecial = regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString
	)

	if !hasUpper(password) {
		return errors.New(constants.ErrPasswordNoUppercase)
	}

	if !hasLower(password) { 
		return errors.New(constants.ErrPasswordNoLowercase)
	}

	if !hasDigit(password) {
		return errors.New(constants.ErrPasswordNoDigit)
	}

	if !hasSpecial(password) {
		return errors.New(constants.ErrPasswordNoSpecialCharacter)
	}

	return nil
}

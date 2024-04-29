package utils

import (
	"authentication/commons/constants"
	"errors"
	"regexp"
)

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New(constants.ErrorPasswordTooShort)
	}
	var (
		hasValidFormat = regexp.MustCompile(constants.PasswordRegex).MatchString
	)
	if !hasValidFormat(password) {
		return errors.New(constants.ErrorInvalidPasswordFormat)
	}
	return nil
}

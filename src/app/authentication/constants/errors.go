package constants

import "errors"

var (
	ErrGeneric                = errors.New("error")
	ErrOpenDatabaseConnection = errors.New("error opening database connection")
	ErrReadConfigFile         = errors.New("error reading config file")
	ErrDecodeConfig           = errors.New("unable to decode config into struct")
	ErrDatabasePing           = errors.New("error pinging database")
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
	ErrFailedToSetNewPassword = errors.New("failed to set a new password")
	ErrMissingToken           = errors.New("missing token")
	ErrPrasingToken           = errors.New("error parsing token")
	ErrInvalidToken           = errors.New("invalid token")
	ErrConnectingDB           = errors.New("error while connecting db")
)

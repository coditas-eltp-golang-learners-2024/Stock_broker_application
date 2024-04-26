package constants

const (
	ErrOpenDatabaseConnection     = "error opening database connection"
	ErrReadConfigFile             = "error reading config file"
	ErrDecodeConfig               = "unable to decode config into struct"
	ErrDatabasePing               = "error pinging database"
	ErrInvalidEmailOrPassword     = "invalid email or password"
	ErrFailedToSetNewPassword     = "failed to set a new password"
	ErrConnectingDB               = "error while connecting db"
	ErrPasswordTooShort           = "password length must be at least 8 characters"
	ErrPasswordNoUppercase        = "password must contain at least one uppercase letter"
	ErrPasswordNoLowercase        = "password must contain at least one lowercase letter"
	ErrPasswordNoSpecialCharacter = "password must contain at least one special character"
	ErrPasswordNoDigit            = "password must contain at least one digit"
)

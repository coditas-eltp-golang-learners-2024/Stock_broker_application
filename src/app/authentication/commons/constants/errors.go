package constants

const (
	ErrorOpenDatabaseConnection = "error opening database connection"
	ErrorReadConfigFile         = "error reading config file"
	ErrorDecodeConfig           = "unable to decode config into struct"
	ErrorDatabasePing           = "error pinging database"
	ErrorInvalidEmailOrPassword = "invalid email or password"
	ErrorFailedToSetNewPassword = "failed to set a new password"
	ErrorConnectingDB           = "error while connecting db"
	ErrorPasswordTooShort       = "password length must be at least 8 characters"
	ErrorInvalidPasswordFormat  = "password must contain a combination of letters, numbers, and special characters"
)

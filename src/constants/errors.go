package constants

// Common Error Messages
const (
	ExternalServiceError            = "server failed to start"
	InternalServerError             = "internal server error"
	DatabaseQueryError              = "failed while querying from database: %w"
	UpdateDatabaseQueryError        = "failed to update the database: %w"
	APIYamlFileUnmarshalError       = "failed to unmarshal apis.yaml: %w"
	ResponseBodyReadError           = "failed to read response body: %w"
	HashPasswordError               = "error hashing password: %w"
	HitRestAPIError                 = "failed to hit the REST API: %w"
	WebSocketConnectionUpgradeError = "failed to upgrade connection to WS "
	DirectoryReadError              = "failed while reading directory"
	NoDataFoundError                = "no data found"
	BadRequestError                 = "bad request error"
	ValidatorError                  = "error while registering validator"
	RateLimitExceeded               = "server has reached maximum request limit"
	ErrorInitializeTraceProvider    = "failed to initialize open telemetry trace provider: %w"
	ErrorShoutDownTraceProvider     = "failed to shutdown open telemetry trace provider: %w"
	ErrorValidatePassword           = "invalid Password Format. Password should be alphanumeric"
)

// Initialization Error Messages
const (
	RedisInitializationError       = "failed to init redis"
	PostgresDBInitializationError  = "failed to init PostgresDb"
	PostgresDBModelsMigrationError = "failed to migrate PostgresDb models"
)

// Config Error Messages
const (
	ConfigParsingError            = "config %s error : %w"
	CreateCloudClusterConfigError = "failed to create cloud cluster config %+v"
	GetApiConfigError             = "failed to get API config: %w"
	GetApplicationConfigError     = "failed to get application config: %w"
	GetRedisConfigError           = "failed to get Redis config: %w"
	GetPostgresConfigError        = "failed to get Postgres config: %w"
	GetNestAPIConfigError         = "failed to get NEST API Mock config: %w"
	ConfigFilePathError           = "failed to get config file path: %w"
	ConfigFileReadError           = "failed to read config file: %w"
	ConfigFileUnmarshalError      = "failed to unmarshal config file: %w"
)

// Redis Error Messages
const (
	RedisInitializationFailedError     = "redis initiation failed. error %+v"
	KeyExistsError                     = "failed while checking whether key %s exists: %w"
	SettingExpirationForKeyExistsError = "failed while setting Expiration for key %s exists: %w"
	GetRedisHostsURLKeyError           = "failed to get redis hosts url key: %w"
	CloseRedisClientError              = "failed to close redis client"
	RedisPoolSizeKeyError              = "failed to get redis pool size key: %w"
	GetRedisMaxConnectionAgeKeyError   = "failed to get redis maxConnAge key: %w"
	GetRedisMinIdleConnectionKeyError  = "failed to get redis min idle conn key: %w"
	GetRedisPoolTimeoutKeyError        = "failed to get redis poolTimeout key: %w"
	GetRedisReadTimeoutKeyError        = "failed to get redis readTimeout key: %w"
	SetKeyError                        = "failed to set key %s: %w"
	GetKeyError                        = "failed to get key %s: %w"
	MSetError                          = "failed during MSet: %w"
	MgetError                          = "failed during Mget: %w"
	HSetError                          = "failed during HSet: %w"
	HGetError                          = "failed during HGet: %w"
	HMGetError                         = "failed during HMGet: %w"
	HGetAllError                       = "failed during HGetAll: %w"
)

// ScyllaDB Error Messages
const (
	ExecuteQueryError     = "failed to execute query: %+v"
	ConnectToClusterError = "failed to connect to cluster %+v"
)

// Postgres Error Messages
const (
	PostgresConnectionError     = "failed to connect to Postgres: %w"
	PostgresMockConnectionError = "failed to connect to Postgres mock: %w"
	ClosePostgresClientError    = "failed to close Postgres client"
)

// Binding Error Messages
const (
	JsonBindingFailedError   = "json binding fail"
	ConfigBindingFailedError = "config binding fail"
	JsonBindingFieldError    = "unexpected value for the field"
)

// Validation Error Message
const (
	DateFormatValidationError          = "give the correct format: YYYY/MM/DD"
	PANFormatValidationError           = "give the correct format"
	RetentionDateFormatValidationError = "give the correct format: DD/MM/YYYY or D/M/YYYY"
	ScannerTypeValueValidationError    = "give the correct value"
	EnumValidationError                = "invalid enum value"
	LtValidationError                  = "must be less than %v"
	GtValidationError                  = "must be greater than %v"
	LteValidationError                 = "must be less than or equal to %v"
	GteValidationError                 = "must be greater than or equal to %v"
	RequiredValidationError            = "this field is required for %s"
	RequiredWithoutValidationError     = "this field is required if %s is not given"
	RequiredWithoutAllValidationError  = "this field is required if %s are not given"
	LteFieldValidationError            = "must be less than or equal to %s"
	MinValidationError                 = "must be greater than %v"
	MaxValidationError                 = "must be less than or equal to %v"
	AlphaNumericValidationError        = "must be alphanumeric"
	BidLengthValidationError           = "please enter bid details"
	UpperCaseValidationError           = "must be uppercase"
	LowerCaseValidationError           = "must be lowercase"
)

// NEST API Error Messages
const (
	MockResponseNotFoundError  = "mock response not found"
	JDataPreparationError      = "failed to prepare jData: %w"
	RequestPayloadMarshalError = "failed to marshal request payload: %w"
	GetClientKeyPairError      = "failed to get client key pair: %w"
	SessionExpiredError        = "session expired"
)

// NEST Error Messages
const (
	NestApiCallError               = "failed to call NEST API: %w"
	NestInternalServerError        = "nest internal server error: %s #end#"
	NestApiResponseMarshalError    = "failed to unmarshal NEST API response: %w"
	RestAPINestCallError           = "failed to call REST API from Nest: %w"
	NestApiResponseDecryptionError = "failed to decrypt NEST API response: %w"
	NestApiCallFlowError           = "failed during calling nest api %s flow: %w"
	NestNoDataError                = "No Data"
)

// Public Key Error Messages
const (
	GetInitialServerPublicKeyError                   = "failed to get initial server public key: %w"
	ParseInitialServerPublicKeyError                 = "failed to parse initial server public key: %w"
	ClientKeyPairCreationError                       = "failed to create client key pair: %w"
	ParseClientPublicKeyError                        = "failed to parse client public key: %w"
	ParseClientPrivateKeyError                       = "failed to parse client private key: %w"
	GetPreAuthenticationServerPublicKeyError         = "failed to get pre-authentication server public key: %w"
	ParsePreAuthenticationServerPublicKeyError       = "failed to parse pre-authentication server public key: %w"
	HashPreAuthenticationServerPublicKeyError        = "failed to hash pre-authentication server public key: %w"
	ParsePreAuthenticationServerPublicRsaPemKeyError = "failed to parse pre-authentication server public key as RSA PEM key: %w"
	ParsePublicKeyRsaPemKeyError                     = "failed to parse public key as RSA PEM key: %w"
	FetchServerPublicKeyError                        = "failed to fetch server public key: %w"
	StoreClientPublicKeyError                        = "failed to store client public key to file: %w"
	StoreClientPrivateKeyError                       = "failed to store client private key to file: %w"
	StorePreAuthenticationServerPublicKeyError       = "failed to store preAuthServerPublicKey key to file: %w"
	StorePreAuthenticationServerKeyHashError         = "failed to store preAuthServerKey Hash to file: %w"
	CreatePublicKeyFileError                         = "failed to create public key file: %w"
	ReadPublicKeyError                               = "failed to read public key: %w"
	HashPublicKeyError                               = "failed to hash public key: %w"
	MarshalPrivateKeyError                           = "failed to marshal private key: %w"
	MarshalPublicKeyError                            = "failed to marshal public key: %w"
	ReadPrivateKeyFileError                          = "failed to read private key file: %w"
	MarshalPublicKeyFileError                        = "failed to create private key file: %w"
	ReadPublicKeyHashError                           = "failed to read public key hash: %w"
	ClientPublicKeyError                             = "failed to read client public key: %w"
	PublicKeyParsingError                            = "error parsing public key: "
	RSAPublicKeyCastingError                         = "error casting to RSA public key"
	ParseRSAPublicKeyError                           = "parsed key is not an RSA public key"
	PublicKeyFileOpeningError                        = "error opening public key file: "
	PublicKeyFileReadingError                        = "error reading public key file: %w"
)

// Private Key Error Messages
const (
	ParsePrivateKeyError       = "failed to parse private key: %w"
	PrivateKeyError            = "private key error"
	ParseRSAPrivateKeyError    = "parsed key is not an RSA private key"
	NotRSAPrivateKeyError      = "key is not an RSA private key"
	PrivateKeyParseError       = "failed to parse private key"
	GeneratePrivateKeyError    = "failed to generate private key: %w"
	MarshalPrivateKeyFileError = "failed to marshal private key file: %w"
)

// Encoding Error Messages
const (
	EncodePrivateKeyError     = "failed to encode private key: %w"
	EncodePublicKeyError      = "failed to encode public key file: %w"
	EncodePrivateKeyFileError = "failed to encode private key file: %w"
)

// Decoding Error Messages
const (
	DecodePEMError                     = "failed to decode PEM block"
	DecodePublicKeyError               = "failed to decode public key: %w"
	DecodeCipherTextError              = "failed to decode cipher text: %w"
	ServerPublicKeyResponseDecodeError = "failed while decoding server publicKey response: %w"
	HttpPostResponseDecodeError        = "failed while decoding http post response: %w"
	HttpPatchResponseDecodeError       = "failed while decoding http patch response: %w"
	HttpDeleteResponseDecodeError      = "failed while decoding http delete response: %w"
	HttpGetResponseDecodeError         = "failed while decoding http get response: %w"
)

// Encryption Error Messages
const (
	EncryptionError                    = "encryption failed "
	EncryptionProcessError             = "error during encryption: %w"
	PreAuthEncryptKeysStorageMainError = "failed to store preAuth encryption keys"
	PreAuthEncryptKeysStorageError     = "failed to store preAuth encryption keys: %w"
	PreAuthEncryptKeysLoadingError     = "failed to load server public key from file: %w"
	RequestPayloadEncryptionError      = "failed to encrypt request payload: %w"
	ClientPublicKeyEncryptionError     = "failed to encrypt client public key: %w"
)

// Decryption Error Messages
const (
	DecryptionError                              = "decryption failed "
	DecryptPreAuthenticationServerPublicKeyError = "failed to decrypt pre-authentication server public key: %w"
	DecryptTokenDataError                        = "failed to decrypt token data: %w"
	DecryptEncryptedBlockError                   = "failed to decrypt encrypted block: %w"
)

// JWT Middleware Error Messages
const (
	JWTTokenMissingError        = "no token was provided in the authorization header"
	JWTTokenBearerMissingError  = "no bearer token was provided in the authorization header"
	JWTTokenExpiredError        = "the access token has expired. please generate a new token from the refresh token"
	JWTRefreshTokenExpiredError = "the refresh token has expired. please login again"
	UnmarshalTokenDataError     = "failed to unmarshal token data: %w"
	TokenDataEncryptionError    = "failed to encrypt token data: %w"
	JWTRefreshTokenInvalidError = "invalid refresh token"
	JWTAccessTokenInvalidError  = "invalid access token"
	JWTInvalidTokenError        = "invalid token: %w"
	ExtractClaimsError          = "unable to extract claims"
	ExtractPayloadError         = "unable to extract token payload"
	JWTTokenDataEncryptionError = "unable to encrypt token data	: %w"
	JWTTokenCompressionError    = "unable to compress token data: %w"
	JWTTokenDecompressionError  = "unable to decompress token data: %w"
)

// CryptoHandshake Middleware Error Messages
const (
	DeviceRetrieveError       = "failed to retrieve the device with the given device id: %w"
	DeviceNotFoundError       = "device not found with given device id"
	RequestBodyReadError      = "failed to read request body: %w"
	ParseBFFPrivateKeyError   = "failed to parse bff private key: %w"
	ParseDevicePublicKeyError = "failed to parse device public key: %w"
	DecryptRequestBodyError   = "failed to decrypt request body: %w"
	EncryptResponseBodyError  = "failed to encrypt response body: %w"
	MarshalResponseError      = "failed to marshal response: %w"
)

// HeaderCheck Middleware Error Messages
const (
	XRequestIdKeyMissingError   = "xRequestId key is missing in header"
	DeviceIdKeyMissingError     = "deviceId key is missing in header"
	AppVersionKeyMissingError   = "app version key is missing in header"
	SourceKeyMissingError       = "source key is missing in header"
	AppInstallIdKeyMissingError = "app install id key is missing in header"
	UserAgentKeyMissingError    = "user agent id key is missing in header"
	TimeStampKeyMissingError    = "timestamp key is missing in header"
	TimeStampInvalidFormatError = "timestamp must be in epoch milliseconds format"
	TimeStampDateNotMatchError  = "the device time has not been set to update automatically, please check the settings"
)

// Unmarshal and Marshal Error Messages
const (
	MarshalStringError                 = "failed to marshal string: %w"
	AccountDetailsResponseDecodeError  = "failed to decode account details response: %w"
	AccountDetailsWrapperResponseError = "failed to get response from account details wrapper: %w"
)

// CryptoRSA Error Messages
const (
	UnexpectedPEMBlockTypeError = "unexpected PEM block type"
	CaptureResponseWriterError  = "capture response writer is not set in the context"
)

// SSL Environment Error Messages
const (
	SSLEnvironmentVariablesMissingError = "SSL certificate or key is missing from environment variables"
)

// Date Error Messages
const (
	DateParseError = "failed to parse date using any of the provided layouts"
)

// Authentication Error Messages
const (
	GreaterThanError = "gt"
)

// Test Case Error Messages
const (
	UnexpectedResponseError = "Body content did not match with %v"
)

// Generic Error Messages
const (
	GenericJSONErrorMessage = "error"
	GenericJSONMessage      = "message"
	ValidationErrors        = "validation errors"
)

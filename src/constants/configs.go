package constants

// Configuration Constants
const (
	RouterV1Config                    = "v1"
	RouterV2Config                    = "v2"
	ApplicationJSONTypeConfig         = "application/json"
	JsonConfig                        = "json"
	LoggerConfig                      = "logger"
	ApplicationConfig                 = "application"
	RedisConfig                       = "redis"
	ApiConfig                         = "apis"
	ModelConfig                       = "model"
	PostgresConfig                    = "postgres"
	SwaggerInfoHttpSchemeConfig       = "http"
	SwaggerInfoHttpsSchemeConfig      = "https"
	NestAPIMockConfig                 = "nestAPIResponse"
	NestAPICallConfig                 = "NestAPICall"
	EndPointConfig                    = "endPoint"
	ServerPortConfig                  = "serverPort.serverPort"
	ResponseConfig                    = "response"
	RawResponseConfig                 = "rawResponse"
	RequestPayloadConfig              = "requestPayload"
	RequestURLConfig                  = "requestURL"
	AllowHeaderOriginConfig           = "Origin"
	ExposeHeaderContentLengthConfig   = "Content-Length"
	SSLCertificateCRTConfig           = "SSL_CERTIFICATE_CRT"
	SSLCertificateKeyConfig           = "SSL_CERTIFICATE_KEY"
	UseMocksConfig                    = "appConfig.UseMocks"
	UseDBMocksConfig                  = "appConfig.UseDBMocks"
	EnableUIBFFEncDecConfig           = "appConfig.EnableUIBFFEncDec"
	UseFrontendErrorFormatConfig      = "appConfig.UseFrontendErrorFormat"
	SwaggerHostKey                    = "swagger.swaggerHost"
	RedisURLKey                       = "hostUrl"
	FilePathKey                       = "filePath.encryptionKeysPath"
	AccessTokenSecretKey              = "token.accessSecretKey"
	RefreshTokenSecretKey             = "token.refreshSecretKey"
	AccessTokenExpiryInDaysKey        = "token.accessTokenExpiryInDays"
	RefreshTokenExpiryInDaysKey       = "token.refreshTokenExpiryInDays"
	EnableTokenCompressionKey         = "token.enableTokenCompression"
	EnableRateLimitConfig             = "appConfig.EnableRateLimit"
	EnableMRVDataConfig               = "appConfig.EnableMRVData"
	RateLimitIntervalInSecondConfig   = "appConfig.RateLimitIntervalInSecond"
	RateLimitRequestPerIntervalConfig = "appConfig.RateLimitRequestPerInterval"
	BannerURLPrefixConfig             = "appConfig.BannerURLPrefix"
	EnableOpenTelemetryConfig         = "appConfig.EnableOpenTelemetry"
	RateLimitIntervalInSecondEnvKey   = "RATE_LIMIT_INTERVAL_IN_SECOND"
	RateLimitRequestPerIntervalEnvKey = "RATE_LIMIT_REQUEST_PER_INTERVAL"
	UIConfig                          = "UIConfig"
	OLTPEndPointEnvConfig             = "OLTP_HTTP_ENDPOINT_URL"
	APPEnvironment                    = "APP_ENVIRONMENT"
	UseMocksEnvKey                    = "USE_MOCKS"
	UseDBMocksEnvKey                  = "USE_DB_MOCKS"
	EnableUIBFFEncDecEnvKey           = "ENABLE_UIBFF_ENCRYPT_DECRYPT"
	EnableOpenTelemetryEnvKey         = "ENABLE_OPEN_TELEMETRY"
	EnableRateLimitEnvKey             = "ENABLE_RATE_LIMIT"
)

// Custom validaton constants
const (
	ValidateEnumConfig     = "ValidateEnum"
	RetentionDateConfig    = "RetentionDateValidation"
	ScannerTypeValueConfig = "ScannerTypeValueValidation"
	DataOfBirthConfig      = "DateOfBirthValidaton"
	PANConfig              = "PANValidation"
	BidLengthValidation    = "BidLengthValidation"
)

// Configuration Keys
const (
	GetInitialKey           = "omnenest.GetInitialKey"
	GetPreAuthenticationKey = "omnenest.GetPreAuthenticationKey"
)

// NEST Configuration Constants
const (
	EnableNestEncryptionConfig = "appConfig.EnableNestEncryption"
	KeySize                    = 512
	JKey                       = "jKey"
	JData                      = "jData"
)

const (
	NestAPIRestBaseURL         = "restbaseurl"
	NestAPIScannerBaseURL      = "scannerbaseurl"
	NestAPIGlobalSearchBaseURL = "globalsearchbaseurl"
	NestAPIIPOBaseURL          = "ipobaseurl"
	NestAPITypeToURLMapping    = "nestapitypetourlmapping"
)

// Postgres Configuration Constants
const (
	PostgresHostKey             = "host"
	PostgresPortKey             = "port"
	PostgresUserKey             = "user"
	PostgresPasswordKey         = "password"
	PostgresDBNameKey           = "dbName"
	PostgresSSLModeKey          = "sslMode"
	PostgresTimeZoneKey         = "TimeZone"
	PostgresIsMockConnectionKey = "isMockConnection"
	PostgresDriverName          = "postgres"
	PostgresHostEnv             = "POSTGRES_HOST"
	PostgresPortEnv             = "POSTGRES_PORT"
	PostgresUserEnv             = "POSTGRES_USER"
	PostgresPasswordEnv         = "POSTGRES_PASSWORD"
	PostgresDBNameEnv           = "POSTGRES_DB_NAME"
)

// Redis Configuration Constants
const (
	RedisConfigUsername = "default"
	RedisConfigPassword = "hb9hlBCZ3EUykxX1VGfoP6yq2Fj9SKtL"
)

// Crypto Configuration Constants
const (
	KeyTypePublic          = "PUBLIC KEY"
	KeyTypePrivate         = "PRIVATE KEY"
	ObtainedHashedPassword = "Received hashed password"
)

// Mock Configuration Constants
const (
	MockResponseKey         = "mockresponse"
	MockEncryptionNeededKey = "encryptionneeded"
)

// REST Methods Constants
const (
	PostMethod    = "POST"
	GetMethod     = "GET"
	PatchMethod   = "PATCH"
	DeleteMethod  = "DELETE"
	PutMethod     = "PUT"
	OptionsMethod = "OPTIONS"
)

// Encryption Key File Names
const (
	PreAuthServerKey     = "preAuthServerPublicKey.pem"
	ClientPublicKey      = "clientPublicKey.pem"
	ClientPrivateKey     = "clientPrivateKey.pem"
	PreAuthServerKeyHash = "preAuthServerKeyHash"
)

// Redis Configuration Configuration
const (
	PoolSize     = "redis.PoolSize"
	MinIdleConns = "redis.MinIdleConns"
	MaxConnAge   = "redis.MaxConnAge"
	PoolTimeout  = "redis.PoolTimeout"
	ReadTimeout  = "redis.ReadTimeout"
)

// ScyllaDB Constants
const (
	ScyllaDBYamlFile        = "resources/scylladb.yml"
	ScyllaDBSelectionPolicy = "us-east-1"
	ScyllaDBKeySpace        = "omnenest_scylladb_test"
)

// Custom Validations Constants
const (
	CustomValidatorTag                   = "CustomValidation"
	DateOfBirthFieldName                 = "DateOfBirth"
	PANFieldName                         = "PAN"
	DateOfBirthFormat                    = `^\d{4}-\d{2}-\d{2}$`
	PANFormat                            = `^[A-Z]{5}[0-9]{4}[A-Z]$`
	DateOfBirthFormatMatch               = "2006-01-02"
	RetentionDateFormatMatch             = "2/1/2006"
	RetentionDateFieldName               = "RetentionTimestamp"
	RetentionDateFormat                  = `^(\d{1,2})/(\d{1,2})/\d{4}$` //`^(0[1-9]|[12][0-9]|3[01])-(0[1-9]|1[0-2])-\d{4}$`
	DateFormatMatch                      = `(\d{2}/\d{2}/\d{4} \d{2}:\d{2}:\d{2})|(\d{2}/\d{2}/\d{4})`
	SpreadInstrumentTypeScripTokenFormat = `^\d+ \d+$`
	NonDigitSequence                     = `[^0-9]+`
	Digit                                = `\d+`
	WatchListLimitError                  = `^nest internal server error: number of scrips already is :  (\d+) the number of scrips chosen is :(\d+) the maximum scrips can be added is : (\d+) #end#$`
	WatchListMaxScripsMatch              = `maximum scrips can be added is : \d+`
	DecimalZeroOrComma                   = `\.0+$|,`
	BrokerRecommendationPattern          = `\s+([\d,]+(?:\.\d+)?)RS`
)

// Research Call Response Field Extraction Tags
const (
	InitialPriceTag = "INITIATION"
	StopLossTag     = "SL"
	TargetTag       = "TARGET"
	PriceAtCallTag  = "PRICE AT CALL"
)

// regex patterns key
const (
	NonDigitSequenceKey               = "nonDigitSequence"
	DigitKey                          = "digit"
	InstrumentTypeScripTokenFormatKey = "instrumentTypeScripTokenFormat"
	DecimalZeroOrCommaKey             = "decimalZeroOrComma"
	InitialPriceTagKey                = "initialPriceTag"
	StopLossTagKey                    = "stopLossTag"
	TargetTagKey                      = "targetTag"
	PriceAtCallTagKey                 = "priceAtCallTag"
	DateFormatMatchKey                = `dateFormatMatch`
	WatchListLimitErrorKey            = `watchListLimitError`
	WatchListMaxScripsMatchKey        = `watchListMaxScripsMatch`
)

// Header JWT Context Configuration
const (
	UserID               = "userId"
	PanNumber            = "panNumber"
	ServerKeyPair        = "serverKeyPair"
	UserSessionToken     = "userSessionToken"
	TokenPayloadClaims   = "tokenPayload"
	BFFCtxPublicKey      = "bffPublicKey"
	BFFCtxPrivateKey     = "bffPrivateKey"
	DeviceCtxPublicKey   = "devicePublicKey"
	TokenExpiration      = "expiration"
	AccessToken          = "accessToken"
	RefreshToken         = "refreshToken"
	TokenType            = "tokenType"
	TokenPayload         = "tokenPayload"
	AccountID            = "accountId"
	IssueID              = "issueId"
	SourceKey            = "source"
	CriteriaAttributeKey = "criteriaAttribute"
	ProductAlias         = "productAlias"
	ClearingOrg          = "clearingOrg"
	BranchName           = "branchName"
	EquitySIPMode        = "1"
	EnabledExchangesKey  = "enabledExchanges"
)

// DNS String Configuration
const (
	DNSString = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s"
)

// Path for test configs file
const (
	TestConfig             = "../../../../setupTest/testConfigs"
	MockResponseConfigPath = "../../../../utils/mockResources"
)

// Space Config
const (
	BlankSpace = " "
	EmptySpace = ""
)

// Banners Env variable keys
const (
	BannerURLPrefixEnvKey = "BANNER_URL_PREFIX"
)

// NestAPIUrl Env variable keys
const (
	RestBaseUrlKey         = "REST_BASE_URL"
	ScannerBaseUrlKey      = "SCANNER_BASE_URL"
	GlobalSearchBaseUrlKey = "GLOBAL_SEARCH_BASE_URL"
	IPOBaseUrlKey          = "IPO_BASE_URL"
)

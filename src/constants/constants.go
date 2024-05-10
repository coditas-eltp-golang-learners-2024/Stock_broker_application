package constants

const (
	ConnectionSuccessful             = "Connected successfully"
	RunningServerPort                = "Running Server on port : %v"
	PostgresConnectionSuccessful     = "Successfully connected to Postgres"
	PostgresMockConnectionSuccessful = "Successfully connected to Mock Postgres"
	RedisConnectionSuccessful        = "Successfully connected to Redis"
	NestAPICallSuccessful            = "NestAPICall successful"
	NestAPICallWithEncryptedData     = "NestAPICall request with encrypted data"
	NestInternalServerErrorStartTag  = "nest internal server error: " // This is the start tag for the nest internal server error message
	NestInternalServerErrorEndTag    = " #end#"                       // This is the end tag for the nest internal server error message
	GenericErrorKey                  = "generic"                      // This is the key for generic error to be passed as the field name in response
	BFFResponseSuccessMessage        = "success"                      // This is the success message for the BFF response
	ContentType                      = "Content-Type"
	ValidationErrorKey               = "validation"
	AvailableRateLimitKey            = "Available-Limit"
	Password                         = "password"
	OTP                              = "otp"
)

// Nest Request Constants
const (
	NestMRVTrueValue  = "Y"
	NestMRVFalseValue = "N"
)

// Validate Password Constants
const (
	ValidatePassword       = "validatePassword"
	ValidatePasswordFormat = "ValidatePasswordFormat"
)

// Password Regex: Currently not working so unused
const (
	PasswordRegex = `^[a-zA-Z0-9!@#$%^&*()_+=\-[\]{};:'",.<>/?]{8,}$`
)

// Response Message Constants
const (
	ResponseMessageKey    = "message"
	ApiOkStatusMessage    = "Ok"
	ApiNotOkStatusMessage = "Not_Ok"
	Status500             = "500"
	Authorization         = "Authorization"
	Bearer                = "Bearer "
)

// Request Message Constants
const (
	NestInputZeroValue   = "0"
	NestInputNAValue     = "NA"
	OrderTypeMarketValue = "MKT"
	OrderTypeSLM         = "SL-M"
	OrderTypeSPValue     = "SP"
	OrderTypeLimitValue  = "L"
	OrderTypeSLValue     = "SL"
)

// Validation Constants
const (
	GreaterThanValue = "gt"
	LessThanValue    = "lt"
)

// Column Name Constants
const (
	Id                           = "id"
	Username                     = "username"
	EmailId                      = "email"
	PhoneNumber                  = "phone_number"
	SequenceId                   = "sequence_id"
	IsAccountBlocked             = "is_account_blocked"
	UnblockRequested             = "unblock_requested"
	UpdatedAt                    = "updated_at"
	UserAccessToken              = "user_access_token"
	UserRefreshToken             = "user_refresh_token"
	FirstName                    = "first_name"
	LastName                     = "last_name"
	DeviceId                     = "device_id"
	BFFPublicKey                 = "bff_public_key"
	BFFPrivateKey                = "bff_private_key"
	DevicePublicKey              = "device_public_key"
	AccountId                    = "account_id"
	BrokerId                     = "broker_id"
	BrokerName                   = "broker_name"
	UserType                     = "user_type"
	EnabledExchange              = "enabled_exchange"
	EnabledOrderType             = "enabled_order_type"
	EnabledProductType           = "enabled_product_type"
	ProductAliasKey              = "product_alias"
	TransactionFlag              = "transaction_flag"
	DefaultMarketWatchlistName   = "default_market_watchlist_name"
	BranchId                     = "branch_id"
	MarketWatchCount             = "market_watch_count"
	WebLink                      = "web_link"
	PasswordSpecialCharacterFlag = "password_special_character_flag"
	UserPrivileges               = "user_privileges"
	YSXExchangeFlag              = "ysx_exchange_flag"
	CriteriaAttribute            = "criteria_attribute"
	InterOpClearingOrg           = "clearing_org"
	AttributeOrderType           = "attribute_order_type"
	EquitySIPOperationMode       = "equity_sip_operation_mode"
	PriceTypeColumn              = "price_type"
	ExchangeColumn               = "exchange"
	RetentionTypeColumn          = "retention_type"
	TitleColumn                  = "title"
	ContentSnippetColumn         = "content_snippet"
	LinkColumn                   = "link"
	PubDateColumn                = "pub_date"
	SourceColumn                 = "source"
	IndicesOrderingColumn        = "indices_ordering"
	FAQColumn                    = "faq"
	Banners                      = "banner"
	Properties                   = "properties"
	SebiNumber                   = "sebi_number"
	CreatedAt                    = "created_at"
	EpochTimestamp               = "epochtimestamp"
	WatchlistName                = "watchlist_name"
	UserId                       = "user_id"
)

// SQL Query Entry Keys
const (
	DeviceIdSQLEntry = "device_id = ?"
	JDataSQLEntry    = "?jData="
	JKeySQLEntry     = "&jKey="
)

// NEST API URL Keys
const (
	OmnenestKey    = "omnenest"
	NestAPITypeKey = "nestapitypetourlmapping"
)

// Table Names Constants
const (
	UserTable                  = "users"
	WatchlistTable             = "watchlist"
	DeviceTable                = "devices"
	BrokerTable                = "brokers"
	StaticExchangeConfigTables = "static_exchange_config"
	NewsAndFeedTable           = "news_and_feeds"
)

// JSON Names
const (
	ExchangeKey        = "exchange"
	ScripTokenKey      = "scripToken"
	TransactionTypeKey = "transactionType"
	QuantityKey        = "quantity"
	TradingSymbolKey   = "tradingSymbol"
)

// NEST API Constants
const (
	VendorCode = ""
)

// NEST API short values
const (
	OrderTypeLimitShort  = "L"
	OrderTypeMarketShort = "MKT"
	ActionBuyShort       = "B"
	ActionSellShort      = "S"
)

// Field Names
const (
	Action                      = "Action"
	OrderType                   = "OrderType"
	OrderStatus                 = "OrderStatus"
	InputOrderStatus            = "InputOrderStatus"
	OrderDate                   = "OrderDate"
	ValidityDate                = "ValidityDate"
	TransactionType             = "TransactionType"
	Leg1TransactionType         = "Leg1TransactionType"
	Leg2TransactionType         = "Leg2TransactionType"
	Leg3TransactionType         = "Leg3TransactionType"
	Leg4TransactionType         = "Leg4TransactionType"
	PriceType                   = "PriceType"
	PositionType                = "PositionType"
	ProductCode                 = "ProductCode"
	RetentionType               = "RetentionType"
	ExchangeName                = "ExchangeName"
	OriginalPriceType           = "OriginalPriceType"
	SendAlertsOn                = "SendAlertsOn"
	BFFOrderStatus              = "BFFOrderStatus"
	InstrumentName              = "InstrumentName"
	Option                      = "Option"
	AlertType                   = "AlertType"
	SourceField                 = "Source"
	Segment                     = "Segment"
	ScannerType                 = "ScannerType"
	ScannerTypeValue            = "ScannerTypeValue"
	ConversionTypeField         = "ConversionType"
	CurrentProductCode          = "CurrentProductCode"
	TargetProductCode           = "TargetProductCode"
	AfterMarketOrderFlag        = "AfterMarketOrderFlag"
	IPOStatus                   = "IPOStatus"
	OFSStatus                   = "OFSStatus"
	ExchangeSegment             = "ExchangeSegment"
	ClientSubCategoryCode       = "ClientSubCategoryCode"
	BFFTriggerStatus            = "BFFTriggerStatus"
	PaymentMode                 = "PaymentMode"
	DematAccountNumberFieldName = "DematAccountNumber"
	BFFBidStatus                = "BFFBidStatus"
	BidHistory                  = "BidHistory"
)

// Date and Time Constant
const (
	DateOldLayout1                 = "02-Jan-2006 15:04:05"
	DateOldLayout2                 = "02-01-2006"
	DateOldLayout3                 = "02-Jan-2006"
	DateOldLayout4                 = "02 Jan, 2006"
	DateOldLayout5                 = "2006-01-02"
	DateOldLayout6                 = "02 Jan,2006"
	DateOldLayout7                 = "02/01/2006"
	DateNewLayout                  = "02/01/2006 15:04:05"
	TimeTrimChars                  = " 00:00:00"
	PrometheusDateTimeFormat       = "06010215"
	PrometheusDateTimeMinuteFormat = "0601021504"
)

// Bff Response Constant Fields
const (
	OrderTypeMarket        = "Market"
	OrderTypeLimit         = "Limit"
	ShortCapsBuy           = "B"
	ShortCapsSell          = "S"
	ShortCapsNil           = "SO"
	ActionBuy              = "Buy"
	ActionSell             = "Sell"
	OrderStatusExecuted    = "Executed"
	OrderStatusOpen        = "Open"
	OrderStatusRejected    = "Rejected"
	OrderStatusComplete    = "complete"
	OrderStatusOpenStr     = "open"
	OrderStatusRejectedStr = "rejected"
	SLBMExchange           = "SLBM"
	PositionTypeDay        = "DAY"
)

// Date Format Fields
const (
	DateInField = "date"
)

var BFFToNestRequestMapping = map[string]string{

	//Transaction Type Mapping
	"S":      "S",
	"SELL":   "S",
	"B":      "B",
	"BUY":    "B",
	"BORROW": "B",
	"LEND":   "L",
	"RP":     "RP",
	"REPAY":  "RP",
	"RC":     "RC",
	"RECALL": "RC",
	//Price Type Mapping
	"MARKET":    "MKT",
	"M":         "MKT",
	"MKT":       "MKT",
	"LIMIT":     "L",
	"L":         "L",
	"SL-L":      "SL-L",
	"SL-M":      "SL-M",
	"SP":        "SP",
	"SP-M":      "SP-M",
	"TWO LEG":   "2L",
	"2L":        "2L",
	"THREE LEG": "3L",
	"3L":        "3L",
	"FOUR LEG":  "4L",
	"4L":        "4L",
	//Exchange Name Mapping
	"NSE":     "NSE",
	"BSE":     "BSE",
	"NSE IPO": "NSE IPO",
	"BSE IPO": "BSE IPO",
	"NSE OFS": "NSE OFS",
	"BSE OFS": "BSE OFS",
	"NCDEX":   "NCDEX",
	"NFO":     "NFO",
	"MCX":     "MCX",
	"CDS":     "CDS",
	"ICEX":    "ICEX",
	"NMCE":    "NMCE",
	"DGCX":    "DGCX",
	"MCXSX":   "MCXSX",
	"BFO":     "BFO",
	"NSEL":    "NSEL",
	"MCXSXCM": "MCXSXCM",
	"MCXSXFO": "MCXSXFO",
	"NDM":     "NDM",
	"BCD":     "BCD",
	"BSEMF":   "BSEMF",
	"NCO":     "NCO",
	"BCO":     "BCO",
	"SLBM":    "SLBM",
	//Position Type Mapping
	"DAY": "DAY",
	"NET": "NET",
	//Retention Type Mapping
	"IOC":   "IOC",
	"FOK":   "FOK",
	"GTC":   "GTC",
	"GTD":   "GTD",
	"GTDys": "GTDys",
	"GTDYS": "GTDys",
	"GTT":   "GTT",
	"OPG":   "OPG",
	"EOS":   "EOS",
	"COL":   "COL",
	//Product Code Mapping
	"NRML": "NRML",
	"CNC":  "CNC",
	"MIS":  "MIS",
	"CO":   "CO",
	"BO":   "BO",
	//Order Status Mapping
	"ACCEPTED":                               "accepted",
	"REJECTED":                               "rejected",
	"OPEN":                                   "open",
	"COMPLETE":                               "complete",
	"PUT ORDER REQ RECEIVED":                 "put order req received",
	"VALIDATION PENDING":                     "validation pending",
	"OPEN PENDING":                           "open pending",
	"MODIFY VALIDATION PENDING":              "modify validation pending",
	"MODIFY PENDING":                         "modify pending",
	"MODIFIED":                               "modified",
	"NOT MODIFIED":                           "not modified",
	"CANCEL PENDING":                         "cancel pending",
	"CANCELLED":                              "cancelled",
	"NOT CANCELLED":                          "not cancelled",
	"FROZEN":                                 "frozen",
	"AFTER MARKET ORDER REQ RECEIVED":        "after market order req received",
	"MODIFY AFTER MARKET ORDER REQ RECEIVED": "modify after market order req received",
	"CANCELLED AFTER MARKET ORDER":           "cancelled after market order",
	"LAPSED":                                 "Lapsed",
	"TRIGGER PENDING":                        "trigger pending", //If nest is accepting input as "Trigger pending", it needs to be handled here//
	"MODIFY ORDER REQ RECEIVED":              "pending",         //need to check all nest order status for ipo modification
	//Source Mapping
	"INDXALRT": "indxalrt",
	"SLTP":     "SLTP",
	"TTVOLFD":  "TTVOLFD",
	"TTVALFD":  "TTVALFD",
	"SVWAP":    "SVWAP",
	//ScannerType Mapping
	"OPEN-LOW":       "open_low",
	"OPEN-HIGH":      "open_high",
	"PRICE-SHOCKER":  "ltp",
	"RISING-FALLING": "riseFall",
	//ScannerTypeValue Mapping
	"LOW":                  "low",
	"HIGH":                 "high",
	"52HIGH":               "52-high",
	"52LOW":                "52-low",
	"1%-UPPERCIRCUIT":      "1%uppercircuit",
	"1%-LOWERCIRCUIT":      "1%lowercircuit",
	"UPPERCIRCUIT":         "uppercircuit",
	"LOWERCIRCUIT":         "lowercircuit",
	"PRICEUP-VOLUMEUP":     "priceup_volumeup",
	"PRICEUP-VOLUMEDOWN":   "priceup_volumedown",
	"PRICEDOWN-VOLUMEUP":   "pricedown_volumeup",
	"PRICEDOWN-VOLUMEDOWN": "pricedown_volumedown",
	//AlertType
	"STOCK": "stock",
	"INDEX": "index",
	"ORDER": "order",
	"Y":     "YES",
	"YES":   "YES",
	"N":     "NO",
	"NO":    "NO",
	//Status
	"ONGOING":  "C",
	"CLOSED":   "P",
	"UPCOMING": "F",
	"ALL":      "A",
	//Payment Mode
	"INVESTMENT-ACCOUNT": "IA",
	"DIRECT-PAYMENT":     "DP",
	//ClientSubCategory Mapping
	"SHAREHOLDER":   "SHA",
	"EMPLOYEE":      "EMP",
	"POLICY HOLDER": "POL",
	"INDIVIDUAL":    "IND",
	//Order Type
	"RETAIL":     "Retail",
	"NON-RETAIL": "Non-institutional",
}

var NestToBFFResponseMapping = map[string]string{

	//Order Type Mapping
	"L":   "Limit",
	"MKT": "Market",
	//Transaction Type Mapping
	"B":  "Buy",
	"S":  "Sell",
	"SO": "Squared-off",
	//Order Status Mapping
	"complete":                               "Executed",
	"open":                                   "Open",
	"rejected":                               "Rejected",
	"reject":                                 "Rejected",
	"put order req received":                 "ReqRecd",
	"validation pending":                     "InValidation",
	"open pending":                           "OpnPend",
	"modify validation pending":              "ModValidPend",
	"modify pending":                         "ModPending",
	"modified":                               "Modified",
	"not modified":                           "NotModified",
	"cancel pending":                         "CanclPending",
	"cancelled":                              "Cancelled",
	"not cancelled":                          "Not Cancelled",
	"frozen":                                 "Frozen",
	"after market order req received":        "AmoReqRecd",
	"modify after market order req received": "ModAmoReqRecd",
	"modify order req received":              "ModOrdReqRecd",
	"cancelled after market order":           "CancelledAmo",
	"Lapsed":                                 "Lapsed",
	"Trigger pending":                        "TrigPending",
	"trigger pending":                        "TrigPending",
	//Exchange Name Mapping
	"NSE":     "NSE",
	"BSE":     "BSE",
	"NCDEX":   "NCDEX",
	"NFO":     "NFO",
	"MCX":     "MCX",
	"CDS":     "CDS",
	"ICEX":    "ICEX",
	"NMCE":    "NMCE",
	"DGCX":    "DGCX",
	"MCXSX":   "MCXSX",
	"BFO":     "BFO",
	"NSEL":    "NSEL",
	"MCXSXCM": "MCXSXCM",
	"MCXSXFO": "MCXSXFO",
	"NDM":     "NDM",
	"BCD":     "BCD",
	"BSEMF":   "BSEMF",
	"NCO":     "NCO",
	"BCO":     "BCO",
	"SLBM":    "SLBM",
	//Price Type Mapping
	"SL-L": "SL-L",
	"SL-M": "SL-M",
	"SP":   "SP",
	"SP-M": "SP-M",
	"2L":   "TWO LEG",
	"3L":   "THREE LEG",
	"4L":   "FOUR LEG",
	//Product Code Mapping
	"NRML": "NRML",
	"CNC":  "CNC",
	"MIS":  "MIS",
	"CO":   "CO",
	"BO":   "BO",
	//SendAlertsOn Mapping
	"1": "SENTON_EMAIL",
	"2": "SENTON_SMS",
	"3": "SENTON_EMAIL_SMS",
	//AlertType Mapping
	"SECURITY LAST TRADE PRICE":          "stock",
	"TOTAL TRADED VOLUME FOR THE DAY":    "stock",
	"TOTAL TRADED VALUE FOR THE DAY":     "stock",
	"SECURITY VOLUME WEIGHTED AVG PRICE": "stock",
	"INDEX ALERT":                        "index",
	"INDEX VALUE":                        "index",
	//IPO Status
	"UPCOMING": "Upcoming",
	"ONGOING":  "Ongoing",
	"CLOSED":   "Closed",
	//ClientSubCategory Mapping
	"SHA": "Shareholder",
	"EMP": "Employee",
	"POL": "Policy Holder",
	"IND": "Individual",
}

var ScannersTypeMap = map[string]string{
	"low":                  "price-shocker",
	"high":                 "price-shocker",
	"52high":               "price-shocker",
	"52low":                "price-shocker",
	"uppercircuit":         "price-shocker",
	"lowercircuit":         "price-shocker",
	"1%-uppercircuit":      "price-shocker",
	"1%-lowercircuit":      "price-shocker",
	"priceup-volumeup":     "rising-falling",
	"priceup-volumedown":   "rising-falling",
	"pricedown-volumeup":   "rising-falling",
	"pricedown-volumedown": "rising-falling",
}

var ExchangeToExchangeSegmentMapping = map[string]string{
	"NSE":     "nse_cm",
	"BSE":     "bse_cm",
	"NCDEX":   "ncx_fo",
	"NFO":     "nse_fo",
	"MCX":     "mcx_fo",
	"CDS":     "cde_fo",
	"ICEX":    "icx_fo",
	"NMCE":    "nmc_fo",
	"DGCX":    "dgx_fo",
	"MCXSX":   "mcx_sx",
	"BFO":     "bse_fo",
	"NSEL":    "nsel_sm",
	"MCXSXCM": "mcx_cm",
	"MCXSXFO": "mcx_cmfo",
	"NDM":     "nse_dm",
	"BCD":     "bcs_fo",
	"BSEMF":   "bse_mf",
	"NCO":     "nse_com",
	"BCO":     "bse_com",
	"ANY":     "any",
	"SLBM":    "nse_slb",
}

var ExchangeSegmentToExchangeMapping = map[string]string{
	"nse_cm":   "NSE",
	"bse_cm":   "BSE",
	"ncx_fo":   "NCDEX",
	"nse_fo":   "NFO",
	"mcx_fo":   "MCX",
	"cde_fo":   "CDS",
	"icx_fo":   "ICEX",
	"nmc_fo":   "NMCE",
	"dgx_fo":   "DGCX",
	"mcx_sx":   "MCXSX",
	"bse_fo":   "BFO",
	"nsel_sm":  "NSEL",
	"mcx_cm":   "MCXSXCM",
	"mcx_cmfo": "MCXSXFO",
	"nse_dm":   "NDM",
	"bcs_fo":   "BCD",
	"bse_mf":   "BSEMF",
	"nse_com":  "NCO",
	"bse_com":  "BCO",
	"any":      "ANY",
	"nse_slb":  "SLBM",
	"bse_ipo":  "BSE",
	"nse_ipo":  "NSE",
	"bse_ofs":  "BSE",
	"nse_ofs":  "NSE",
}

var ExchangeToInterOpSegementMapping = map[string]string{
	"NSE":     "CASH",
	"BSE":     "CASH",
	"NCDEX":   "COM",
	"NFO":     "FO",
	"MCX":     "COM",
	"CDS":     "CUR",
	"BFO":     "FO",
	"MCXSXCM": "CASH",
	"BCD":     "CUR",
	"NCO":     "COM",
	"BCO":     "COM",
	"CASH":    "CASH",
	"FO":      "FO",
	"CUR":     "CUR",
	"SLBM":    "SLB",
}

var ExchangeToSegmentIndicatorMapping = map[string]string{
	"NSE":     "EQUITY",
	"BSE":     "EQUITY",
	"NCDEX":   "COMMODITY",
	"NFO":     "FNO",
	"MCX":     "COMMODITY",
	"CDS":     "CURRENCY",
	"BFO":     "FNO",
	"MCXSXCM": "EQUITY",
	"BCD":     "CURRENCY",
	"NCO":     "COMMODITY",
	"BCO":     "COMMODITY",
	"CASH":    "EQUITY",
	"FO":      "FNO",
	"CUR":     "CURRENCY",
	"SLBM":    "SLB",
}

// Nest to BFF Transaction Type Mapping for SLBM
var SLBMTransactionTypeMapping = map[string]string{
	"B":  "Borrow",
	"L":  "Lend",
	"RP": "Repay",
	"RC": "Recall",
}

// Conversion Type mapping
var ConversionTypeMapping = map[string]string{
	"DAY": "D",
	"CF":  "C",
}

// Common Constants
const (
	ExchangeBSE                  = "BSE"
	ExchangeNSE                  = "NSE"
	GroupRS                      = "RS"
	OFSRetailType                = "Retail"
	OFSNonRetailType             = "Non-Retail"
	ExchangeCashInterOp          = "CASH"
	ExchangeFnoInterOp           = "FNO"
	ExchangeCurrencyInterOp      = "CUR"
	SpreadContractIdentifier     = "SP-"
	DefaultDenominatorFloatValue = 1.0
)

var InstrumentFilterOut = map[string]bool{
	"CUR":    true,
	"COMDTY": true,
	"UNDCUR": true,
	"UNDIRD": true,
	"UNDIRC": true,
	"UNDIRT": true,
	"UL":     true,
	"AUCSO":  true,
	"INDEX":  true,
	"UNDCOM": true,
}

var PreferableExchangeToExchangeSegmentMapping = map[string]string{
	"CASH": "nse_cm",
	"FNO":  "nse_fo",
	"CUR":  "cde_fo",
}

const (
	SegmentIndicatorForSpreadInstrumentType = "SPREAD"
)

// Exchange Names
const (
	NFOExchange = "NFO"
	BFOExchange = "BFO"
	CDSExchange = "CDS"
)

// Alerts Field Type mapping
var AlertFieldTypeMapping = map[string]string{
	"SECURITY LAST TRADE PRICE":          "Last Traded Price",
	"TOTAL TRADED VOLUME FOR THE DAY":    "Total Traded Volume",
	"TOTAL TRADED VALUE FOR THE DAY":     "Total Traded Value",
	"SECURITY VOLUME WEIGHTED AVG PRICE": "Volume Weighted Avg Price",
}

// Exchange Segment Indicator
const (
	EquityExchangeSegmentIndicator    = "EQUITY"
	CommodityExchangeSegmentIndicator = "COMMODITY"
	FNOExchangeSegmentIndicator       = "FNO"
	CurrencyExchangeSegmentIndicator  = "CURRENCY"
)

// Product Code
const (
	BOProductCode   = "BO"
	COProductCode   = "CO"
	NRMLProductCode = "NRML"
	CNCProductCode  = "CNC"
	MISProductCode  = "MIS"
)

// Prometheus Metrics
const (
	NestOverallLatency = "NESTOverallLatency"
	APIRequestTime     = "APIRequestTime"
	ServiceNameLabel   = "ServiceName"

	HttpRequestMetricLabel  = "HTTP_Requests"
	NestOverallLatencyLabel = "NEST-Overall"
	NestAPILatencyLabel     = "NEST"
	BFFLatencyLabel         = "BFF"
	APILatencyLabel         = "Overall"
)

// NEST API URL Keys
const (
	AccountDetailsKey = "accountdetails"
)

// OLTP HTTP API URL Keys
const (
	OLTPHttpEndpointUrl  = "localhost:4318"
	OLTPHttpTimeoutInSec = 10
)

// Application Environment
const (
	DevEnvironment  = "DEV"
	ProdEnvironment = "PROD"
	QAEnvironment   = "QA"
)

// BFF ENC DEC Middleware Skipper
const (
	StocksIntradayAggrDataEndpoint = "/v1/stocks/intraday-aggr-data"
	ChartSource                    = "CHART"
	AutomationFlag                 = "AUTOMATION"
)

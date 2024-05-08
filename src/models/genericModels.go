package models

import (
	"stock_broker_application/src/constants"

	"github.com/dgrijalva/jwt-go"
)

type ErrorMessage struct {
	Key          string `json:"key,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type ErrorAPIResponse struct {
	Message []ErrorMessage `json:"errors,omitempty"`
	Error   string         `json:"error,omitempty"`
}

type Request struct {
	Type     string      `json:"type,omitempty"`     // Sender
	Method   string      `json:"method,omitempty"`   // Recipient
	ClientId string      `json:"clientId,omitempty"` // Content
	UserId   string      `json:"userId,omitempty"`   // Content
	Request  interface{} `json:"request,omitempty"`  // Content
}

type EncryptResponse struct {
	EncResponse string `json:"encResponse"`
}

type EncryptRequest struct {
	EncRequest string `json:"encRequest" validate:"required"`
}

type EncryptedNestAPIResponse struct {
	EncryptedResponse string `json:"jEncResp,omitempty"`
}

// type TokenData struct {
// 	UserId            string   `json:"uid"`
// 	UserSessionId     string   `json:"userSessionId"`
// 	BFFPublicKey      string   `json:"bffPublicKey"`
// 	BFFPrivateKey     string   `json:"bffPrivateKey"`
// 	DevicePublicKey   string   `json:"devicePublicKey"`
// 	AccountId         string   `json:"accountId"`
// 	BrokerName        string   `json:"brokerName"`
// 	BranchName        string   `json:"branchName"`
// 	ProductAlias      string   `json:"productAlias"`
// 	CriteriaAttribute []string `json:"criteriaAttribute"`
// 	ClearingOrg       string   `json:"clearingOrg"`
// 	EnabledExchanges  []string `json:"enabledExchange"`
// }

type ChannelResponse struct {
	ApiEndpoint string
	Response    []byte
	Error       error
	Metadata    interface{}
}

type HttpGoRoutineRequest struct {
	ApiEndpoint string
	Request     interface{}
	Metadata    interface{}
}

// claim model for email
type TokenData struct {
	UserId string `gorm:"column:id" json:"id"`
	jwt.StandardClaims
}

func (TokenData) TableName() string {
	return constants.UserTable
}
// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/auth/change-password": {
            "patch": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Change a user's password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Change Password"
                ],
                "summary": "Change Password",
                "parameters": [
                    {
                        "description": "Change Password Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ChangePassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Password changed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/forgot-password": {
            "post": {
                "description": "Updates user credentials based on the provided request.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Forgot Password"
                ],
                "summary": "Update user credentials",
                "operationId": "update-credentials",
                "parameters": [
                    {
                        "description": "Forgot Password Request JSON",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ForgotPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Password updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/signin": {
            "post": {
                "description": "Handle sign-in request and authenticate the user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SignIN"
                ],
                "summary": "Handle sign-in request",
                "parameters": [
                    {
                        "description": "Sign-in request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User authenticated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/signup": {
            "post": {
                "description": "Create a new user with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SignUP"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserSignUp"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/validate-otp": {
            "post": {
                "description": "Validates the OTP for a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Validate OTP"
                ],
                "summary": "Validate OTP",
                "parameters": [
                    {
                        "description": "OTP Request",
                        "name": "otpRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ValidateOTPRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OTP validated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "OTP is expired or invalid",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ChangePassword": {
            "type": "object",
            "required": [
                "newPassword",
                "oldPassword"
            ],
            "properties": {
                "newPassword": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 8,
                    "example": "Coditas@18"
                },
                "oldPassword": {
                    "type": "string",
                    "example": "S@nket123"
                }
            }
        },
        "models.ForgotPasswordRequest": {
            "type": "object",
            "required": [
                "email",
                "panCardNumber"
            ],
            "properties": {
                "NewPassword": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 8,
                    "example": "sample11110"
                },
                "email": {
                    "type": "string",
                    "example": "testUser@gmail.com"
                },
                "panCardNumber": {
                    "type": "string",
                    "example": "abgjhi6789"
                }
            }
        },
        "models.SignInRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 8,
                    "example": "Goat@018"
                },
                "userName": {
                    "type": "string",
                    "example": "virat"
                }
            }
        },
        "models.UserSignUp": {
            "type": "object",
            "required": [
                "email",
                "name",
                "panCard",
                "password",
                "phoneNumber",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "panCard": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 8
                },
                "phoneNumber": {
                    "type": "integer",
                    "maximum": 9999999999,
                    "minimum": 1000000000
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.ValidateOTPRequest": {
            "type": "object",
            "required": [
                "id",
                "otp"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "otp": {
                    "type": "integer",
                    "maximum": 9999,
                    "minimum": 1000
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Stock Broker Application",
	Description:      "api for Stock Broker using gin and gorm",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

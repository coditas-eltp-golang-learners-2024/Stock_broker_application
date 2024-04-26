package handler

import (
	"authentication/business"
	"authentication/commons/constants"
	"authentication/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"stock_broker_application/src/utils/validations"
)

// SignInHandler handles the sign-in request
// @Summary Handle sign-in request
// @Description Handle sign-in request and authenticate the user
// @Accept json
// @Produce json
// @Param request body models.SignInRequest true "Sign-in request body"
// @Success 200 {object} string "User authenticated successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Router /v1/authentication/signin [post]
func SignInHandler(userService *business.SignInService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var signInRequest models.SignInRequest

		// Bind JSON request body to SignInRequest struct
		if err := c.ShouldBindJSON(&signInRequest); err != nil {
			c.JSON(http.StatusBadRequest, constants.ErrorBadRequest)
			return
		}
		customValidator := validations.NewCustomValidator()
		// Register custom validation functions
		customValidator.Validator.RegisterValidation("email", validations.ValidateEmail)
		customValidator.Validator.RegisterValidation("strong_password", validations.ValidateStrongPassword)

		// Validate the struct using the custom validator
		if err := customValidator.ValidateStruct(c.Request.Context(), signInRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call SignIn method to authenticate user
		if err := userService.SignIn(signInRequest); err != nil {
			c.JSON(http.StatusUnauthorized, constants.ErrorMessageAuthenticationFailed)
			return
		}

		// Authentication successful
		c.JSON(http.StatusOK, constants.SuccessMessageSignIn)

	}
}

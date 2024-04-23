package handlers

import (
	"authentication/business"
	"authentication/models"
	"authentication/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(userSignUpService business.SignUpService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := utils.SignUpValidation(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := userSignUpService.SignUp(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	}
}

package router

import (
	"authentication/business"
	"authentication/constants"
	"authentication/handlers"
	"authentication/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	userRepoSignUp := repositories.NewUserSignUpInstance(db)
	userSignUpService := business.NewSignUpService(userRepoSignUp)
	router.POST(constants.SignUp, handlers.SignUp(*userSignUpService))

	return router

}

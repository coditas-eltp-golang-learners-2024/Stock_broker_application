package utils

import (
	"net/http"

	genericConstants "stock_broker_application/src/constants"
	genericModels "stock_broker_application/src/models"

	"github.com/gin-gonic/gin"
)

func SendBadRequest(ctx *gin.Context, err []genericModels.ErrorMessage) {
	ctx.JSON(http.StatusBadRequest, genericModels.ErrorAPIResponse{
		Message: err,
		Error:   http.StatusText(http.StatusBadRequest),
	})
}

func SendInternalServerError(ctx *gin.Context, err string) {
	message := genericModels.ErrorMessage{
		Key:          genericConstants.GenericErrorKey,
		ErrorMessage: err,
	}
	ctx.JSON(http.StatusInternalServerError, genericModels.ErrorAPIResponse{
		Message: []genericModels.ErrorMessage{message},
		Error:   http.StatusText(http.StatusInternalServerError),
	})
}

func SendConflictError(ctx *gin.Context, err string) {
	message := genericModels.ErrorMessage{
		Key:          genericConstants.GenericErrorKey,
		ErrorMessage: err,
	}
	ctx.JSON(http.StatusConflict, genericModels.ErrorAPIResponse{
		Message: []genericModels.ErrorMessage{message},
		Error:   http.StatusText(http.StatusConflict),
	})
}

func SendNoContentError(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, genericModels.ErrorAPIResponse{
		Message: []genericModels.ErrorMessage{},
		Error:   http.StatusText(http.StatusNoContent),
	})
}

func SendStatusOk(ctx *gin.Context, data interface{}) {
	value, exist := ctx.Get(genericConstants.GenericTokenMessage)
	if exist {
		ctx.JSON(http.StatusOK, gin.H{
			genericConstants.GenericJSONMessage:  data,
			genericConstants.GenericTokenMessage: value,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			genericConstants.GenericJSONMessage: data,
		})
	}

}

func SendNewDataCreated(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, data)
}

func SendUnauthorizedError(ctx *gin.Context, err string) {
	message := genericModels.ErrorMessage{
		Key:          genericConstants.GenericErrorKey,
		ErrorMessage: err,
	}
	ctx.JSON(http.StatusUnauthorized, genericModels.ErrorAPIResponse{
		Message: []genericModels.ErrorMessage{message},
		Error:   http.StatusText(http.StatusUnauthorized),
	})
}

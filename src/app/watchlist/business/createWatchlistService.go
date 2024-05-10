package business

import (
	"errors"
	"github.com/gin-gonic/gin"
	genericConstants "stock_broker_application/src/constants"
	serviceConstants "watchlist/commons/constants"
	serviceModel "watchlist/models"
	"watchlist/repositories"
)

type CreateWatchlistService struct {
	CreateWatchlistRepository repositories.CreateWatchlistRepository
}

func NewCreateWatchlistService(createWatchlistInstance repositories.CreateWatchlistRepository) *CreateWatchlistService {
	return &CreateWatchlistService{
		CreateWatchlistRepository: createWatchlistInstance,
	}
}

func (service *CreateWatchlistService) CreateWatchlistService(serviceModel serviceModel.CreateWatchlist, ctx *gin.Context) error {
	id := ctx.Value(genericConstants.Id).(uint16)
	condition := map[string]interface{}{
		genericConstants.UserId:        id,
		genericConstants.WatchlistName: serviceModel.WatchlistName,
	}
	if service.CreateWatchlistRepository.IsWatchlistExists(condition) {
		return errors.New(serviceConstants.CreateWatchlistAlreadyExistsError)
	}
	if !service.CreateWatchlistRepository.CreateWatchlist(condition) {
		return errors.New(serviceConstants.CreateWatchlistIsNotAddedError)
	}
	if !service.CreateWatchlistRepository.IsScripsAdded(condition, serviceModel) {
		return errors.New(serviceConstants.CreateWatchlistWrongScripError)
	}

	return nil
}

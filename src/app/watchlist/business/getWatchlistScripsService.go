package business

import (
	"errors"
	"github.com/gin-gonic/gin"
	genericConstants "stock_broker_application/src/constants"
	serviceConstants "watchlist/commons/constants"
	"watchlist/models"
	"watchlist/repositories"
)

type WatchlistScripsService struct {
	watchlistRepository repositories.WatchlistScripsRepository
}

func NewWatchlistScripsService(watchlistRepository repositories.WatchlistScripsRepository) *WatchlistScripsService {
	return &WatchlistScripsService{
		watchlistRepository: watchlistRepository,
	}
}

func (service *WatchlistScripsService) GetScrips(context *gin.Context, watchlistName string) (*models.GetWatchlistScrips, error) {
	userID := context.Value(genericConstants.Id).(uint16)
	condition := map[string]interface{}{
		genericConstants.UserId:       userID,
		genericConstants.WatchlistName: watchlistName,
	}

	watchlistID, err := service.watchlistRepository.GetWatchlistsByUserID(condition)
	if err != nil {
		return nil, err
	}

	if watchlistID == 0 {
		return nil, errors.New(serviceConstants.WatchlistNotFoundError)
	}

	stockIDSlice, err := service.watchlistRepository.GetStockIDsByWatchlistID(watchlistID)
	if err != nil {
		return nil, err
	}

	if len(stockIDSlice) == 0 {
		return  nil, errors.New(serviceConstants.NoStocksInWatchlistError)
	}

	scrips, err := service.watchlistRepository.GetScripsByStockID(stockIDSlice)
	if err != nil {
		return nil, err
	}

	return &scrips, nil
}
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
func (service *WatchlistScripsService) GetScrips(context *gin.Context, watchlistName string) ([]models.Scrip, error) {
	userID := context.Value(genericConstants.Id).(uint16)
	condition := map[string]interface{}{
		genericConstants.UserId:        userID,
		genericConstants.WatchlistName: watchlistName,
	}

	watchlistExists, err := service.watchlistRepository.CheckWatchlistExists(condition)
	if err != nil {
		return nil, err
	}
	if !watchlistExists {
		return nil, errors.New(serviceConstants.WatchlistNotFoundError)
	}

	watchlistID, err := service.watchlistRepository.GetWatchlistsByUserID(condition)
	if err != nil {
		return nil, err
	}
	if watchlistID == 0 {
		return nil, errors.New(serviceConstants.WatchlistIDNotFoundError) 
	}

	stockIDSlice, err := service.watchlistRepository.GetStockIDsByWatchlistID(watchlistID)
	if err != nil {
		return nil, err
	}
	if len(stockIDSlice) == 0 {
		return nil, errors.New(serviceConstants.NoStocksInWatchlistError)
	}

	scrips, err := service.watchlistRepository.GetScripsByStockID(stockIDSlice)
	if err != nil {
		return nil, err
	}

	return scrips, nil
}
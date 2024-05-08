package business

import (
	"errors"
	"fmt"
	genericConstants "stock_broker_application/src/constants"
	"watchlist/models"
	"watchlist/repositories"

	"github.com/gin-gonic/gin"
)

type WatchlistService struct {
	watchlistRepository repositories.WatchlistRepository
}

func NewWatchlistService(watchlistRepository repositories.WatchlistRepository) *WatchlistService {
	return &WatchlistService{
		watchlistRepository: watchlistRepository,
	}
}

func (service *WatchlistService) GetScrips(ctx *gin.Context, request models.WatchlistWithScripsRequest) error {
	// Check if the watchlist name exists for the given user ID
	userID := ctx.Value(genericConstants.Id).(uint16)
	watchlistExists, err := service.watchlistRepository.CheckWatchlistExists(userID, request.WatchlistName)
	if err != nil {
		return err
	}
	if !watchlistExists {
		return errors.New("watchlist not found for the given user ID and name")
	}

	// If the watchlist exists, continue with your logic
	watchlistIDSlice, err := service.watchlistRepository.GetWatchlistsByUserID(userID, request.WatchlistName)
	if err != nil {
		return err
	}
	fmt.Println("watchlist found", watchlistIDSlice)

	stockIDSlice, err := service.watchlistRepository.GetStockIDsByWatchlistID(watchlistIDSlice)
	if err != nil {
		return err
	}
	fmt.Println("stock found", stockIDSlice)
	scripsMap,err:=service.watchlistRepository.GetScripsByStockID(stockIDSlice)
	if err != nil {
		return err
	}
	fmt.Println("scrips found", scripsMap)

	return nil
}

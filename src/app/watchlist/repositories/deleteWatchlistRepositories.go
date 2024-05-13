package repositories

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/models"
	"watchlist/commons/constants"
)

type DeleteWatchlistRepository interface {
	DeleteScrips(ctx *gin.Context, watchlistName string, scrips []int) error
}

type watchlistDBRepository struct {
	DB *gorm.DB
}

func NewDeleteWatchlistRepository(db *gorm.DB) DeleteWatchlistRepository {
	return &watchlistDBRepository{DB: db}
}

func (repository *watchlistDBRepository) DeleteScrips(ctx *gin.Context, watchlistName string, scrips []int) error {
	userID, exists := ctx.Get(genericConstants.Id)
	if !exists {
		return errors.New(constants.UserIDNotFoundError)
	}

	var watchListID uint
	if err := repository.DB.Model(&models.Watchlist{}).Where("user_id = ? AND watchlist_name = ?", userID, watchlistName).Select("id").First(&watchListID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(constants.WatchlistNotFoundError)
		}
		return errors.New(constants.FailedtoFindWatchlistError)
	}

	var result []uint
	for _, stocksID := range scrips {
		var stockID []uint
		if err := repository.DB.Model(&models.Stocks{}).
			Where("id = ?", stocksID).Pluck("id", &stockID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New(constants.StockwithSymbolNotFoundError)
			}
			return errors.New(constants.StockNotFoundError)
		}
		result = append(result, stockID...)
	}

	// Delete entries from watchlist_stocks table
	for _, stockID := range result {
		if err := repository.DB.Where("watchlist_id = ? AND stocks_id = ?", watchListID, stockID).Delete(&models.WatchlistStock{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New(constants.WatchlistNotFoundError)
			}
			return errors.New(constants.FailToDeleteScripsError)
		}
	}
	return nil
}

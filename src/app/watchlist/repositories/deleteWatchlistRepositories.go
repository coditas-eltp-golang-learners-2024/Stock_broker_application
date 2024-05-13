package repositories

import (
	"errors"
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		return errors.New(constants.ErrorUserIDNotFound)
	}

	// Find watchlist ID
	var watchListID uint
	if err := repository.DB.Model(&models.Watchlist{}).Where("user_id = ? AND watchlist_name = ?", userID, watchlistName).Select("id").First(&watchListID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(constants.ErrorWatchlistNotFound)
		}
		return err
	}

	var result uint

	for _, symbol := range scrips {

		if err := repository.DB.Model(&models.Stocks{}).
			Where("symbol = ?", symbol).Pluck("id", &result).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("stock not found")
			}

			return err
		}
	}

	// Delete entries from watchlist_stocks table
	if err := repository.DB.Where("watchlist_id = ? AND stocks_id = ?", watchListID, result).Delete(&models.WatchlistStock{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(constants.ErrorWatchlistNotFound)
		}
		return errors.New(constants.ErrorFailedToDeleteScrips)
	}
	return nil
}

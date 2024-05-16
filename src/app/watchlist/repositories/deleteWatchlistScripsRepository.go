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
	DeleteScrips(ctx *gin.Context, watchlistName string, scrips []int, watchlistCondition map[string]interface{}) error
}

type watchlistDbRepository struct {
	DB *gorm.DB
}

func NewDeleteWatchlistRepository(db *gorm.DB) DeleteWatchlistRepository {
	return &watchlistDbRepository{DB: db}
}

func (repository *watchlistDbRepository) DeleteScrips(ctx *gin.Context, watchlistName string, scrips []int, watchlistCondition map[string]interface{}) error {

	var watchListID uint
	if err := repository.DB.Model(&models.Watchlist{}).Where(watchlistCondition).Pluck(genericConstants.Id, &watchListID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(constants.WatchlistNotFoundError)
		}
		return errors.New(constants.FailedtoFindWatchlistError)
	}

	for _, stockID := range scrips {
		deleteScripsCondition := map[string]interface{}{
			genericConstants.WatchlistID: watchListID,
			genericConstants.StocksID:    stockID,
		}
		if err := repository.DB.Where(deleteScripsCondition).Delete(&models.WatchlistStock{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New(constants.WatchlistNotFoundError)
			}
			return errors.New(constants.FailToDeleteScripsError)
		}
	}
	return nil
}

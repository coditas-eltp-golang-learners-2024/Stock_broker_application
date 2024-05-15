package repositories

import (
	"errors"
	"stock_broker_application/src/models"
	"watchlist/commons/constants"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetWatchlistsRepository interface {
	GetWatchlists(ctx *gin.Context, condition map[string]interface{}) ([]string, error)
}

type watchlistDBRepository struct {
	DB *gorm.DB
}

func NewGetWatclistsRepository(db *gorm.DB) GetWatchlistsRepository {
	return &watchlistDBRepository{DB: db}
}

func (repository *watchlistDBRepository) GetWatchlists(ctx *gin.Context, condition map[string]interface{}) ([]string, error) {
	var watchlistSlice []models.Watchlist
	var watchlistsNames []string
	var err error
	if err = repository.DB.Model(&models.Watchlist{}).Where(condition).Find(&watchlistSlice).Error; err != nil {
		return watchlistsNames, errors.New(constants.WatchlistNotFoundError)
	}

	for _, watchlist := range watchlistSlice {
		watchlistsNames = append(watchlistsNames, watchlist.WatchlistName)
	}

	return watchlistsNames, nil
}

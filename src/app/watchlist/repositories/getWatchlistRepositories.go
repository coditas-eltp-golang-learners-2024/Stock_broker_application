package repositories

import (
	"errors"
	genericModels "stock_broker_application/src/models"
	"watchlist/commons/constants"
	"watchlist/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetWatchlistsRepository interface {
	GetWatchlists(ctx *gin.Context, condition map[string]interface{}) (models.GetWatchlists, error)
}

type watchlistDBRepository struct {
	DB *gorm.DB
}

func NewGetWatclistsRepository(db *gorm.DB) GetWatchlistsRepository {
	return &watchlistDBRepository{DB: db}
}

func (repository *watchlistDBRepository) GetWatchlists(ctx *gin.Context, condition map[string]interface{}) (models.GetWatchlists, error) {
	var watchlistSlice []genericModels.Watchlist
	var watchlistsNames models.GetWatchlists
	var err error
	if err = repository.DB.Model(&genericModels.Watchlist{}).Where(condition).Find(&watchlistSlice).Error; err != nil {
		return watchlistsNames, errors.New(constants.WatchlistNotFoundError)
	}

	for _, watchlist := range watchlistSlice {
		watchlistsNames.Watchlist = append(watchlistsNames.Watchlist, watchlist.WatchlistName)
	}

	return watchlistsNames, nil
}

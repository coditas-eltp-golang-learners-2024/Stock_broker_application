package repositories

import (
	"stock_broker_application/src/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetWatchlistsRepository interface {
	GetWatchlists(ctx *gin.Context, condition map[string]interface{}) ([]string, error)
}

type userDBRepository struct {
	DB *gorm.DB
}

func NewGetWatclistsRepository(db *gorm.DB) GetWatchlistsRepository {
	return &userDBRepository{DB: db}
}

func (repository *userDBRepository) GetWatchlists(ctx *gin.Context, condition map[string]interface{}) ([]string, error) {
	var watchlistSlice []models.Watchlist
	var watchlistNames []string
	var err error
	if err = repository.DB.Model(&models.Watchlist{}).Where(condition).Find(&watchlistSlice).Error; err != nil {
		return watchlistNames, err
	}

	for _, watchlist := range watchlistSlice {
		watchlistNames = append(watchlistNames, watchlist.WatchlistName)
	}

	return watchlistNames, nil
}

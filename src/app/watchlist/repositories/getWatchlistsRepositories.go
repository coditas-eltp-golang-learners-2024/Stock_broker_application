package repositories

import (
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/models"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetWatchlistsRepository interface {
	GetWatchlists(ctx *gin.Context) (string, error)
}

type userDBRepository struct {
	DB *gorm.DB
}

func NewGetWatclistsRepository(db *gorm.DB) GetWatchlistsRepository {
	return &userDBRepository{DB: db}
}

func (repository *userDBRepository) GetWatchlists(ctx *gin.Context) (string, error) {
	var watchlistSlice []models.Watchlist
	var watchlists string
	userID := ctx.Value(genericConstants.UserID).(string)
	//return repository.DB.Where("user_id = ? AND watchlist_name = ?", userId, watchlist).Find(&watchlist).Error
	if err := repository.DB.Model(&models.Watchlist{}).Where("user_id = ?", userID).Find(&watchlistSlice).Error; err != nil {
		return watchlists, err
	}
	var watchlistNames []string
	for _, watchlist := range watchlistSlice {
		watchlistNames = append(watchlistNames, watchlist.WatchlistName)
	}
	watchlistsData := strings.Join(watchlistNames, ", ")

	return watchlistsData, nil
}

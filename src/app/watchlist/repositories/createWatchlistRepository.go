package repositories

import (
	serviceConstants "stock_broker_application/src/constants"
	dbModel "stock_broker_application/src/models"
	"watchlist/models"

	"gorm.io/gorm"
)

type CreateWatchlistRepository interface {
	IsWatchlistExists(condition map[string]interface{}) bool
	CreateWatchlist(condition map[string]interface{}) bool
	IsScripsAdded(condition map[string]interface{}, createWatchlistRequest models.CreateWatchlist) bool
}

type UserDBRepository struct {
	db *gorm.DB
}

func NewUserDBRepository(dataBase *gorm.DB) *UserDBRepository {
	return &UserDBRepository{db: dataBase}
}

func (user *UserDBRepository) IsWatchlistExists(condition map[string]interface{}) bool {
	var count int64
	if err := user.db.Model(&dbModel.Watchlist{}).Where(condition).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (user *UserDBRepository) CreateWatchlist(condition map[string]interface{}) bool {
	var count int64
	if err := user.db.Model(&dbModel.Watchlist{}).Create(condition).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (user *UserDBRepository) IsScripsAdded(condition map[string]interface{}, createWatchlistRequest models.CreateWatchlist) bool {
	var count int64
	var watchlistId uint
	var stocksId uint

	if err := user.db.Model(&dbModel.Watchlist{}).Where(condition).Pluck(serviceConstants.Id, &watchlistId).Error; err != nil {
		return false
	}
	for _, scripValue := range createWatchlistRequest.Scrips {
		query := map[string]interface{}{
			serviceConstants.Token: scripValue,
		}
		if err := user.db.Model(&dbModel.Stocks{}).Where(query).Pluck(serviceConstants.Id, &stocksId).Count(&count).Error; err != nil {
			return false
		}
		if count > 0 {
			newRecord := map[string]interface{}{
				serviceConstants.WatchlistID: watchlistId,
				serviceConstants.StocksID:    stocksId,
			}
			if err := user.db.Model(&dbModel.WatchlistStock{}).Create(&newRecord).Error; err != nil {
				return false
			}
		} else {
			return count > 0
		}

	}
	return count > 0
}

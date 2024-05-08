package repositories

import (
	"gorm.io/gorm"
	"stock_broker_application/src/models"
)

type WatchlistRepository interface {
	CheckWatchlistExists(userID uint16, watchlistName string) (bool, error)
	GetWatchlistsByUserID(userID uint16 , watchlistName string) ([]uint16, error)
	GetStockIDsByWatchlistID(watchlistIDSlice []uint16) ([]uint16, error)
	GetScripsByStockID(stockIDSlice []uint16) ([]map[string]interface{}, error)
}
type watchlistDBRepository struct {
	db *gorm.DB
}

func NewWatchlistRepository(db *gorm.DB) *watchlistDBRepository {
	return &watchlistDBRepository{db: db}
}
func (repo *watchlistDBRepository) CheckWatchlistExists(userID uint16, watchlistName string) (bool, error) {
	var count int64
	if err := repo.db.Model(&models.Watchlist{}).
		Where("user_id = ? AND watchlist_name = ?", userID, watchlistName).
		Count(&count).Error; err != nil {
		return count>0, err
	}
	return true, nil
}
func (repo *watchlistDBRepository) GetWatchlistsByUserID(userID uint16 , watchlistName string) ([]uint16, error) {
	var watchlists []models.Watchlist
	var watchlistIDs []uint16
	if err := repo.db.Model(&models.Watchlist{}).
		Where("user_id = ? AND watchlist_name = ?", userID, watchlistName).
		Find(&watchlists).
		Error; err != nil {
		return nil, err
	}
	for _, elem := range watchlists {
		watchlistIDs = append(watchlistIDs, uint16(elem.ID))
	}
	
	return watchlistIDs, nil
}
func (repo *watchlistDBRepository) GetStockIDsByWatchlistID(watchlistIDSlice []uint16) ([]uint16, error) {
    var stockIDs []uint16
    for _, watchlistID := range watchlistIDSlice {
        var stocks []models.WatchlistStock
        if err := repo.db.Model(&models.WatchlistStock{}).
            Where("watchlist_id = ?", watchlistID).
            Find(&stocks).
            Error; err != nil {
            return nil, err
        }
        for _, elem := range stocks {
            stockIDs = append(stockIDs, uint16(elem.StockID))
        }
    }
    return stockIDs, nil
}
func (repo *watchlistDBRepository) GetScripsByStockID(stockIDSlice []uint16) ([]map[string]interface{}, error){
	var scrips []map[string]interface{}
	for _, stockID := range stockIDSlice {
		var stock models.Stocks
		if err := repo.db.Model(&models.Stocks{}).
			Where("token = ?", stockID).
			Find(&stock).
			Error; err != nil {
			return nil, err
		}
		scrip:=map[string]interface{}{
			"token":stock.Token,
			"symbol":stock.Symbol,
		}
		scrips=append(scrips, scrip)
	}
	return scrips, nil
}

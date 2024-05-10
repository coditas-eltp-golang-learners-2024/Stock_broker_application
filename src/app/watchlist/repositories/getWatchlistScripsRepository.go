package repositories

import (
	dbmodels "stock_broker_application/src/models"
	"watchlist/models"
	"gorm.io/gorm"
)

type WatchlistScripsRepository interface {
	CheckWatchlistExists(userID uint16, watchlistName string) (bool, error)
	GetWatchlistsByUserID(userID uint16, watchlistName string) (uint16, error)
	GetStockIDsByWatchlistID(watchlistID uint16) ([]uint16, error)
	GetScripsByStockID(stockIDSlice []uint16) ([]models.Scrip, error)
}
type watchlistDBRepository struct {
	db *gorm.DB
}

func NewWatchlistRepository(db *gorm.DB) *watchlistDBRepository {
	return &watchlistDBRepository{db: db}
}
func (repo *watchlistDBRepository) CheckWatchlistExists(userID uint16, watchlistName string) (bool, error) {
	var count int64
	if err := repo.db.Model(&dbmodels.Watchlist{}).
		Where("user_id = ? AND watchlist_name = ?", userID, watchlistName).
		Count(&count).Error; err != nil {
		return count > 0, err
	}
	return true, nil
}
func (repo *watchlistDBRepository) GetWatchlistsByUserID(userID uint16, watchlistName string) (uint16, error) {
	var watchlists dbmodels.Watchlist
	if err := repo.db.Model(&dbmodels.Watchlist{}).
		Where("user_id = ? AND watchlist_name = ?", userID, watchlistName).
		Find(&watchlists).
		Error; err != nil {
		return 0, err
	}
	return watchlists.ID, nil
}
func (repo *watchlistDBRepository) GetStockIDsByWatchlistID(watchlist uint16) ([]uint16, error) {
	var stockIDs []uint16
	var stocks []dbmodels.WatchlistStock
	if err := repo.db.Model(&dbmodels.WatchlistStock{}).
		Where("watchlist_id = ?", watchlist).
		Find(&stocks).
		Error; err != nil {
		return nil, err
	}
	for _, elem := range stocks {
		stockIDs = append(stockIDs, uint16(elem.StockID))
	}
	return stockIDs, nil
}
func (repo *watchlistDBRepository) GetScripsByStockID(stockIDSlice []uint16) ([]models.Scrip, error) {
	var scrips []models.Scrip
	for _, stockID := range stockIDSlice {
		var scrip models.Scrip
		if err := repo.db.Model(&dbmodels.Stocks{}).
			Where("id = ?", stockID).
			Find(&scrip).
			Error; err != nil {
			return nil, err
		}
		scrips = append(scrips, scrip)
	}
	return scrips, nil
}

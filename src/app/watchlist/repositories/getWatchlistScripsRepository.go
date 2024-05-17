package repositories

import (
	genericConstants "stock_broker_application/src/constants"
	genericModels "stock_broker_application/src/models"
	"watchlist/models"

	"gorm.io/gorm"
)

type WatchlistScripsRepository interface {
	CheckWatchlistExists(condition map[string]interface{}) (bool, error)
	GetWatchlistsByUserID(condition map[string]interface{}) (uint, error)
	GetStockIDsByWatchlistID(watchlistID uint) ([]uint, error)
	GetScripsByStockID(stockIDSlice []uint) ([]models.Scrip, error)
}
type watchlistDBScripsRepository struct {
	db *gorm.DB
}

func NewWatchlistRepository(db *gorm.DB) *watchlistDBScripsRepository {
	return &watchlistDBScripsRepository{db: db}
}
func (repo *watchlistDBScripsRepository) CheckWatchlistExists(condition map[string]interface{}) (bool, error) {
	var count int64
	if err := repo.db.Model(&genericModels.Watchlist{}).
		Where(condition).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (repo *watchlistDBScripsRepository) GetWatchlistsByUserID(condition map[string]interface{}) (uint, error) {
	var watchlists genericModels.Watchlist
	if err := repo.db.Model(&genericModels.Watchlist{}).
		Where(condition).
		Find(&watchlists).
		Error; err != nil {
		return 0, err
	}
	return watchlists.ID, nil
}

func (repo *watchlistDBScripsRepository) GetStockIDsByWatchlistID(watchlist uint) ([]uint, error) {
	var stockIDs []uint
	var stocks []genericModels.WatchlistStock
	if err := repo.db.Model(&genericModels.WatchlistStock{}).
		Where(genericConstants.WatchlistID, watchlist).
		Find(&stocks).
		Error; err != nil {
		return nil, err
	}
	for _, elem := range stocks {
		stockIDs = append(stockIDs, elem.StockID)
	}
	return stockIDs, nil
}

func (repo *watchlistDBScripsRepository) GetScripsByStockID(stockIDSlice []uint) ([]models.Scrip, error) {
	var GetWatchlistScrips models.GetWatchlistScrips
	for _, stockID := range stockIDSlice {
		var scrip models.Scrip
		if err := repo.db.Model(&genericModels.Stocks{}).
			Where(genericConstants.StocksID, stockID).
			Find(&scrip).
			Error; err != nil {
			return nil, err
		}
		GetWatchlistScrips.WatchlistScrip = append(GetWatchlistScrips.WatchlistScrip, scrip)
	}
	return GetWatchlistScrips.WatchlistScrip, nil
}

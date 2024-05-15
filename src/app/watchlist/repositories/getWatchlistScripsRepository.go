package repositories

import (
	"gorm.io/gorm"
	genericConstants "stock_broker_application/src/constants"
	dbmodels "stock_broker_application/src/models"
	"watchlist/models"
)

type WatchlistScripsRepository interface {
	CheckWatchlistExists(condition map[string]interface{}) (bool, error)
	GetWatchlistsByUserID(condition map[string]interface{}) (uint, error)
	GetStockIDsByWatchlistID(watchlistID uint) ([]uint, error)
	GetScripsByStockID(stockIDSlice []uint) ([]models.Scrip, error)
}
type watchlistDBRepository struct {
	db *gorm.DB
}

func NewWatchlistRepository(db *gorm.DB) *watchlistDBRepository {
	return &watchlistDBRepository{db: db}
}
func (repo *watchlistDBRepository) CheckWatchlistExists(condition map[string]interface{}) (bool, error) {
	var count int64
	if err := repo.db.Model(&dbmodels.Watchlist{}).
		Where(condition).
		Count(&count).Error; err != nil {
		return count > 0, err
	}
	return true, nil
}
func (repo *watchlistDBRepository) GetWatchlistsByUserID(condition map[string]interface{}) (uint, error) {
	var watchlists dbmodels.Watchlist
	if err := repo.db.Model(&dbmodels.Watchlist{}).
		Where(condition).
		Find(&watchlists).
		Error; err != nil {
		return 0, err
	}
	return watchlists.ID, nil
}

func (repo *watchlistDBRepository) GetStockIDsByWatchlistID(watchlist uint) ([]uint, error) {
	var stockIDs []uint
	var stocks []dbmodels.WatchlistStock
	if err := repo.db.Model(&dbmodels.WatchlistStock{}).
		Where(genericConstants.WatchlistID+" = ?", watchlist).
		Find(&stocks).
		Error; err != nil {
		return nil, err
	}
	for _, elem := range stocks {
		stockIDs = append(stockIDs, elem.StockID)
	}
	return stockIDs, nil
}

func (repo *watchlistDBRepository) GetScripsByStockID(stockIDSlice []uint) ([]models.Scrip, error) {
	var scrips []models.Scrip
	for _, stockID := range stockIDSlice {
		var scrip models.Scrip
		if err := repo.db.Model(&dbmodels.Stocks{}).
			Where(genericConstants.StocksID+" = ?", stockID).
			Find(&scrip).
			Error; err != nil {
			return nil, err
		}
		scrips = append(scrips, scrip)
	}
	return scrips, nil
}

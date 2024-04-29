package repositories

import (
	watchlistModel "stock_broker_application/src/models"
	"watchlist/models"

	"gorm.io/gorm"
)

type EditWatchListRepository interface {
	RenameWatchlist(client *gorm.DB, watchlist *models.WatchlistRenameModel) error
}

type editWatchListRepository struct{}

func NewEditWatchListRepository() *editWatchListRepository {
	return &editWatchListRepository{}
}

func (repo *editWatchListRepository) RenameWatchlist(db *gorm.DB, watchlist *models.WatchlistRenameModel) error {
	err := db.Model(&watchlistModel.Watchlist{}).Where("watchlist_name = ?", watchlist.WatchlistName).Update("watchlist_name", watchlist.NewWatchlistName).Error
	if err != nil {
		return err
	}
	return nil
}

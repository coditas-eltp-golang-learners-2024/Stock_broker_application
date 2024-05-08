package repositories

import (
	"fmt"
	watchlistModel "stock_broker_application/src/models"
	"watchlist/commons/constants"
	"watchlist/models"

	"gorm.io/gorm"
)

type EditWatchListRepository interface {
	RenameWatchlist(client *gorm.DB, watchlist *models.WatchlistRenameModel, userId string) error
}

type editWatchListRepository struct{}

func NewEditWatchListRepository() *editWatchListRepository {
	return &editWatchListRepository{}
}

func (repo *editWatchListRepository) RenameWatchlist(db *gorm.DB, watchlist *models.WatchlistRenameModel, userId string) error {
	err := db.Model(&watchlistModel.Watchlist{}).Where("watchlist_name = ? and user_id=?", watchlist.WatchlistName, userId).Update("watchlist_name", watchlist.NewWatchlistName)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return fmt.Errorf(constants.ErrNoWatchlist)
	}
	return nil
}

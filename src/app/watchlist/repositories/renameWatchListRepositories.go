package repositories

import (
	"errors"
	genericConstants "stock_broker_application/src/constants"
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
	err := db.Model(&watchlistModel.Watchlist{}).Where(genericConstants.WatchlistName+" = ? and "+genericConstants.UserId+"=?", watchlist.WatchlistName, userId).Update(genericConstants.WatchlistName, watchlist.NewWatchlistName)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New(constants.ErrNoWatchlist)
	}
	return nil
}

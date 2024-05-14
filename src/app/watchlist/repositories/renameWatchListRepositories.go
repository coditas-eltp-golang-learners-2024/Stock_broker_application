package repositories

import (
	"errors"
	genericConstants "stock_broker_application/src/constants"
	watchlistModel "stock_broker_application/src/models"
	"watchlist/commons/constants"
	"watchlist/models"

	"gorm.io/gorm"
)

type RenameWatchListRepository interface {
	RenameWatchlist(client *gorm.DB, watchlist *models.WatchlistRenameModel, userId uint16) error
}

type renameWatchListRepository struct{}

func NewRenameWatchListRepository() *renameWatchListRepository {
	return &renameWatchListRepository{}
}

func (repo *renameWatchListRepository) RenameWatchlist(db *gorm.DB, watchlist *models.WatchlistRenameModel, userId uint16) error {
	err := db.Model(&watchlistModel.Watchlist{}).Where(genericConstants.WatchlistName+" = ? and "+genericConstants.UserId+"=?", watchlist.WatchlistName, userId).Update(genericConstants.WatchlistName, watchlist.NewWatchlistName)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New(constants.NoWatchlistError)
	}
	return nil
}

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
	RenameWatchlist(client *gorm.DB, watchlist *models.RenameWatchlist, userId uint16) error
}

type renameWatchListRepository struct{}

func NewRenameWatchListRepository() *renameWatchListRepository {
	return &renameWatchListRepository{}
}

func (repository *renameWatchListRepository) RenameWatchlist(db *gorm.DB, watchlist *models.RenameWatchlist, userId uint16) error {
	var Count int64
	count := db.Model(&watchlistModel.Watchlist{}).Where(genericConstants.WatchlistName+" = ? and "+genericConstants.UserId+"=?", watchlist.NewWatchlistName, userId).Count(&Count)
	if count.Error != nil {
		return count.Error
	}
	if Count > 0 {
		return errors.New(constants.WatchlistAlreadyExistsError)
	}
	err := db.Model(&watchlistModel.Watchlist{}).Where(genericConstants.WatchlistName+" = ? and "+genericConstants.UserId+"=?", watchlist.WatchlistName, userId).Update(genericConstants.WatchlistName, watchlist.NewWatchlistName)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New(constants.NoWatchlistError)
	}
	return nil
}

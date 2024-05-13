package repositories

import (
	"errors"
	genericConstants "stock_broker_application/src/constants"
	watchlistModel "stock_broker_application/src/models"
	"watchlist/commons/constants"
	"watchlist/models"

	"gorm.io/gorm"
)

type DeleteWatchListRepository interface {
	DeleteWatchlist(client *gorm.DB, user *models.WatchlistDeleteModel, userId uint16) error
}

type deleteWatchListRepository struct{}

func NewDeleteWatchListRepository() *deleteWatchListRepository {
	return &deleteWatchListRepository{}
}

func (repo *deleteWatchListRepository) DeleteWatchlist(client *gorm.DB, watchlist *models.WatchlistDeleteModel, userId uint16) error {
	err := client.Model(&watchlistModel.Watchlist{}).Where(genericConstants.WatchlistName+"= ? and "+genericConstants.UserId+"=?", watchlist.WatchlistName, userId).Delete(&watchlistModel.Watchlist{})
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New(constants.ErrNoWatchlist)
	}
	return nil
}

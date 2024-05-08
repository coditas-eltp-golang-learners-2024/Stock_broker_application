package repositories

import (
	"fmt"
	"log"
	watchlistModel "stock_broker_application/src/models"

	"watchlist/commons/constants"
	"watchlist/models"

	"gorm.io/gorm"
)

type DeleteWatchListRepository interface {
	DeleteWatchlist(client *gorm.DB, user *models.WatchlistDeleteModel, userId string) error
}

type deleteWatchListRepository struct{}

func NewDeleteWatchListRepository() *deleteWatchListRepository {
	return &deleteWatchListRepository{}
}

func (repo *deleteWatchListRepository) DeleteWatchlist(client *gorm.DB, watchlist *models.WatchlistDeleteModel, userId string) error {
	log.Println("Before deleting")
	log.Println(userId)
	err := client.Model(&watchlistModel.Watchlist{}).Where("watchlist_name = ? and user_id=?", watchlist.WatchlistName, userId).Delete(&watchlistModel.Watchlist{})

	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return fmt.Errorf(constants.ErrNoWatchlist)
	}
	return nil
}

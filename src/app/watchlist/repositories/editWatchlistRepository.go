package repositories

import (
	"errors"
	genericConstants "stock_broker_application/src/constants"
	genericModels "stock_broker_application/src/models"
	"watchlist/commons/constants"
	"watchlist/models"

	"gorm.io/gorm"
)

type EditWatchlistRepository interface {
	InsertScripsToWatchlist(editWatchlist models.EditWatchlistRequest, userId uint16, client *gorm.DB) error
}

type editWatchlistRepository struct {
}

func NewEditWatchlistRepository() *editWatchlistRepository {
	return &editWatchlistRepository{}
}

func (repository *editWatchlistRepository) InsertScripsToWatchlist(editWatchlist models.EditWatchlistRequest, userId uint16, client *gorm.DB) error {
	var scripsIds []uint
	if err := client.Model(&genericModels.Stocks{}).
		Where(genericConstants.Token+" IN (?)", editWatchlist.Scrips).
		Pluck("id", &scripsIds).
		Error; err != nil {
		return err
	}

	var watchlistId uint
	if err := client.Model(&genericModels.Watchlist{}).
		Where(genericConstants.UserId+" = ? AND "+genericConstants.WatchlistName+" = ?", userId, editWatchlist.WatchlistName).
		Pluck("id", &watchlistId).
		Error; err != nil {
		return err
	}

	var count int64
	if err := client.Model(&genericModels.WatchlistStock{}).
		Where(genericConstants.WatchlistID+" = ? AND "+genericConstants.StocksID+" IN (?)", watchlistId, scripsIds).
		Count(&count).
		Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New(constants.ScripsAlreadyAddedError)
	}

	var watchlistStocks []genericModels.WatchlistStock
	for _, scripID := range scripsIds {
		watchlistStocks = append(watchlistStocks, genericModels.WatchlistStock{
			WatchlistID: watchlistId,
			StockID:     scripID,
		})
	}
	if err := client.Create(&watchlistStocks).Error; err != nil {
		return err
	}

	return nil
}

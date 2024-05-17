package repositories

import (
	dbModels "stock_broker_application/src/models"

	"gorm.io/gorm"
)

type EditWatchlistRepository interface {
	InsertScripsToWatchlist(client *gorm.DB, watchlist *dbModels.Watchlist) error
}

type editWatchlistRepository struct {
}

func NewEditWatchlistRepository() *editWatchlistRepository {
	return &editWatchlistRepository{}
}

func (repository *editWatchlistRepository) InsertScripsToWatchlist(client *gorm.DB) error {

	return nil
}

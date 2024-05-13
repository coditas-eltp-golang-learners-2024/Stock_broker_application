package business

import (
	"errors"
	"github.com/gin-gonic/gin"
	"watchlist/repositories"
)

type DeleteWatchlistService struct {
	deleteWatchlistRepository repositories.DeleteWatchlistRepository
}

func NewDeleteWatchlistService(deleteWatchlistRepository repositories.DeleteWatchlistRepository) *DeleteWatchlistService {
	return &DeleteWatchlistService{
		deleteWatchlistRepository: deleteWatchlistRepository,
	}
}

func (service *DeleteWatchlistService) DeleteScripsFromWatchlist(ctx *gin.Context, watchlistName string, scrips []int) error {
	if watchlistName == "" || len(scrips) == 0 {
		return errors.New("watchlistName and scrips are required")
	}

	err := service.deleteWatchlistRepository.DeleteScrips(ctx, watchlistName, scrips)
	if err != nil {
		return err
	}

	return nil
}

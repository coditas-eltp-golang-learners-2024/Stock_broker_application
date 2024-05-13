package business

import (
	"errors"
	"github.com/gin-gonic/gin"
	"watchlist/commons/constants"
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
	err := service.deleteWatchlistRepository.DeleteScrips(ctx, watchlistName, scrips)
	if err != nil {
		return errors.New(constants.FailedToDeleteScripsfromWatchlistError)
	}

	return nil
}

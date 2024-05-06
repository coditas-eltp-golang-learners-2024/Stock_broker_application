package business

import (
	"watchlist/repositories"

	"github.com/gin-gonic/gin"
)

type NewGetWatchlistService interface {
	NewGetWatchlistService(ctx *gin.Context) (string, error)
}

type getWatchListSercvice struct {
	getWatchlistInterface repositories.GetWatchlistsRepository
}

func NewUsersService(userData repositories.GetWatchlistsRepository) NewGetWatchlistService {
	return &getWatchListSercvice{
		getWatchlistInterface: userData,
	}
}

func (repository *getWatchListSercvice) NewGetWatchlistService(ctx *gin.Context) (string, error) {

	return repository.getWatchlistInterface.GetWatchlists(ctx)
}

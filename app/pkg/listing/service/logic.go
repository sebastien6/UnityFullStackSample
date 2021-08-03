package service

import (
	"context"
	"fmt"

	"github.com/sebastien6/UnityFullStackSample/app/pkg/logger"

	apiv1 "github.com/sebastien6/UnityFullStackSample/app/pkg/api/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

type listingService struct {
	serviceRepo ServiceResository
	apiv1.UnimplementedGamesServiceServer
}

func NewListingService(serviceRepo ServiceResository) apiv1.GamesServiceServer {
	return &listingService{
		serviceRepo,
		apiv1.UnimplementedGamesServiceServer{},
	}
}

// GetGames retreive the list of games form database and return the result as a list
func (s *listingService) GetGames(ctx context.Context, _ *emptypb.Empty) (*apiv1.Games, error) {
	return s.serviceRepo.FindAll(ctx)
}

// RegisterGame a game into the database and return the updated game with its associated ID
func (s *listingService) RegisterGame(ctx context.Context, game *apiv1.Game) (*apiv1.Game, error) {
	return s.serviceRepo.Store(ctx, game)
}

// RegisterGames register a list of game to the database and return the updated games with their associated ID
func (s *listingService) RegisterGames(ctx context.Context, games *apiv1.Games) (*apiv1.Games, error) {
	results := apiv1.Games{}
	for i, game := range games.Listings {
		game, err := s.serviceRepo.Store(ctx, game)
		if err != nil {
			err := fmt.Errorf("an error occured at index %d of listings, error: %v", i, err)
			logger.Log.Sugar().Error(err)
			return nil, err
		} else {
			logger.Log.Sugar().Infof("game at index %d was successfully registered", i)
			results.Listings = append(results.Listings, game)
		}
	}
	return games, nil
}

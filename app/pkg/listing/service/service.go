package service

import (
	"context"

	apiv1 "github.com/sebastien6/UnityFullStackSample/app/pkg/api/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ListingService interface {
	GetGames(ctx context.Context, _ *emptypb.Empty) (*apiv1.Games, error)
	RegisterGame(ctx context.Context, game *apiv1.Game) (*apiv1.Game, error)
	RegisterGames(ctx context.Context, games *apiv1.Games) (*apiv1.Games, error)
}

package service

import (
	"context"

	listingv1 "github.com/sebastien6/UnityFullStackSample/app/pkg/api/v1"
)

type ServiceResository interface {
	FindAll(ctx context.Context) (*listingv1.Games, error)
	Store(ctx context.Context, game *listingv1.Game) (*listingv1.Game, error)
}

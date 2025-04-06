package database

import (
	"context"
	"phrasmotica/flyer-api/models"
)

type IDatabase interface {
	GetFlyers(ctx context.Context) (bool, []models.Flyer)
	AddFlyer(ctx context.Context, newFlyer *models.Flyer) bool

	GetPlayers(ctx context.Context) (bool, []models.Player)
	AddPlayer(ctx context.Context, newPlayer *models.Player) bool
}

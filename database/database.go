package database

import (
	"context"
	"phrasmotica/flyer-api/models"
)

type IDatabase interface {
	GetFlyers(ctx context.Context) (bool, []models.Flyer)
	AddFlyer(ctx context.Context, newFlyer *models.Flyer) bool
}

package database

import (
	"context"
	"phrasmotica/flyer-api/models"
)

type InMemoryDatabase struct {
	Flyers []models.Flyer
}

// GetFlyers implements IDatabase
func (d *InMemoryDatabase) GetFlyers(ctx context.Context) (bool, []models.Flyer) {
	return true, d.Flyers
}

// AddFlyer implements IDatabase
func (d *InMemoryDatabase) AddFlyer(ctx context.Context, newFlyer *models.Flyer) bool {
	d.Flyers = append(d.Flyers, *newFlyer)
	return true
}

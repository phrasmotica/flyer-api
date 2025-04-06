package database

import (
	"context"
	"phrasmotica/flyer-api/models"
)

type InMemoryDatabase struct {
	Flyers  []models.Flyer
	Players []models.Player
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

// GetPlayers implements IDatabase
func (d *InMemoryDatabase) GetPlayers(ctx context.Context) (bool, []models.Player) {
	return true, d.Players
}

// AddPlayer implements IDatabase
func (d *InMemoryDatabase) AddPlayer(ctx context.Context, newPlayer *models.Player) bool {
	d.Players = append(d.Players, *newPlayer)
	return true
}

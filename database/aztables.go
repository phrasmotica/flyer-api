package database

import (
	"context"
	"encoding/json"
	"phrasmotica/flyer-api/logging"
	"phrasmotica/flyer-api/models"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

type TableStorageDatabase struct {
	Client *aztables.ServiceClient
	Logger logging.ILogger
}

// GetFlyers implements IDatabase
func (d *TableStorageDatabase) GetFlyers(ctx context.Context) (bool, []models.Flyer) {
	return true, make([]models.Flyer, 0)
}

// AddFlyer implements IDatabase
func (d *TableStorageDatabase) AddFlyer(ctx context.Context, newFlyer *models.Flyer) bool {
	entity := aztables.EDMEntity{
		Entity: aztables.Entity{
			PartitionKey: "Flyer",
			RowKey:       newFlyer.ID,
		},
		Properties: map[string]interface{}{},
	}

	marshalled, err := json.Marshal(entity)
	if err != nil {
		d.Logger.Error(err)
		return false
	}

	_, addErr := d.Client.NewClient("Flyers").AddEntity(ctx, marshalled, nil)
	if addErr != nil {
		d.Logger.Error(addErr)
		return false
	}

	return true
}

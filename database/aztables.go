package database

import (
	"context"
	"encoding/json"
	"phrasmotica/flyer-api/logging"
	"phrasmotica/flyer-api/models"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

type TableStorageDatabase struct {
	Client *aztables.ServiceClient
}

func CreateTableStorageClient(connStr string) *aztables.ServiceClient {
	client, err := aztables.NewServiceClientFromConnectionString(connStr, &aztables.ClientOptions{
		ClientOptions: policy.ClientOptions{},
	})

	if err != nil {
		logging.Error.Fatal(err)
		return nil
	}

	return client
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
		logging.Error.Println(err)
		return false
	}

	_, addErr := d.Client.NewClient("Flyers").AddEntity(ctx, marshalled, nil)
	if addErr != nil {
		logging.Error.Println(addErr)
		return false
	}

	return true
}

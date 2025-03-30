package database

import (
	"context"
	"os"
	"phrasmotica/flyer-api/logging"
	"phrasmotica/flyer-api/models"
)

type IDatabase interface {
	GetFlyers(ctx context.Context) (bool, []models.Flyer)
	AddFlyer(ctx context.Context, newFlyer *models.Flyer) bool
}

func CreateDb() IDatabase {
	azureTablesConnStr := os.Getenv("AZURE_TABLES_CONNECTION_STRING")
	if azureTablesConnStr != "" {
		logging.Info.Println("Using data backend: Azure Table Storage")

		return &TableStorageDatabase{
			Client: CreateTableStorageClient(azureTablesConnStr),
		}
	}

	logging.Info.Println("Using data backend: In-Memory")

	return &InMemoryDatabase{}
}

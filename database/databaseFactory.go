package database

import (
	"os"
	"phrasmotica/flyer-api/logging"
	"phrasmotica/flyer-api/models"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

type DatabaseFactory struct {
	Logger logging.ILogger
}

func (f *DatabaseFactory) CreateTableStorageClient(connStr string) *aztables.ServiceClient {
	client, err := aztables.NewServiceClientFromConnectionString(connStr, &aztables.ClientOptions{
		ClientOptions: policy.ClientOptions{},
	})

	if err != nil {
		f.Logger.Fatal(err.Error())
		return nil
	}

	return client
}

func (f *DatabaseFactory) CreateDb() IDatabase {
	azureTablesConnStr := os.Getenv("AZURE_TABLES_CONNECTION_STRING")
	if azureTablesConnStr != "" {
		f.Logger.Info("Using data backend: Azure Table Storage")

		return &TableStorageDatabase{
			Client: f.CreateTableStorageClient(azureTablesConnStr),
		}
	}

	f.Logger.Info("Using data backend: In-Memory")

	return &InMemoryDatabase{
		Flyers:  make([]models.Flyer, 0),
		Players: make([]models.Player, 0),
	}
}

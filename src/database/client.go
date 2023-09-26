package database

import (
	"context"

	"github.com/angrypufferfish/goodm/src/exceptions"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDatabase *GoodmDatabase

type GoodmClient struct {
	*mongo.Client
	current *GoodmDatabase
}

func NewGoodmClient(client *mongo.Client) *GoodmClient {
	var goodmClient = &GoodmClient{}

	goodmClient.Client = client

	return goodmClient
}

func GetGoodmDatabase() *GoodmDatabase {
	if mongoDatabase == nil {
		panic(exceptions.ErrDbNotInitialized)
	}
	return mongoDatabase
}

func (odmClient *GoodmClient) UseDatabase(databaseName string, ctx *context.Context, opts ...*options.DatabaseOptions) *GoodmDatabase {

	goodmDatabase := NewGoodmDatabase(odmClient, databaseName, ctx, opts...)

	odmClient.current = goodmDatabase

	mongoDatabase = odmClient.current

	return odmClient.current
}

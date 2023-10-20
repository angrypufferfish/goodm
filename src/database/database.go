package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GoodmDatabase struct {
	*mongo.Database
	current *mongo.Collection
	Context *context.Context
}

func NewGoodmDatabase(odmClient *GoodmClient, databaseName string, ctx *context.Context, opts ...*options.DatabaseOptions) *GoodmDatabase {
	var goodmDatabase = &GoodmDatabase{
		Context: ctx,
	}

	goodmDatabase.Database = odmClient.Database(databaseName, opts...)

	return goodmDatabase
}

func (odmDatabase *GoodmDatabase) UseCollection(collection string) *mongo.Collection {
	odmDatabase.current = odmDatabase.Collection(collection)

	return odmDatabase.current
}

func (odmDatabase *GoodmDatabase) GetCurrenCollection() *mongo.Collection {
	return odmDatabase.current
}

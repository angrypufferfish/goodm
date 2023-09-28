package repository

import (
	"github.com/angrypufferfish/goodm/src/database"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func findOneAndReplace[A any, S any](db *database.GoodmDatabase, filter any, update any, opts ...*options.FindOneAndReplaceOptions) (*S, error) {

	var document S
	collection, err := database.GetCollection[A](db)

	if err != nil {
		return nil, err
	}

	result := collection.FindOneAndReplace(*db.Context, filter, update, opts...)

	if err != nil {
		return nil, err
	}

	document, err = serializeSingleResult[S](result)
	if err != nil {
		return nil, err
	}
	return &document, nil
}

func replaceOne[A any](db *database.GoodmDatabase, filter any, update any, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {

	collection, err := database.GetCollection[A](db)

	if err != nil {
		return nil, err
	}

	result, err := collection.ReplaceOne(*db.Context, filter, update, opts...)

	if err != nil {
		return nil, err
	}

	return result, nil
}

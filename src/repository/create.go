package repository

import (
	"github.com/angrypufferfish/goodm/src/database"
	"go.mongodb.org/mongo-driver/mongo"
)

func Insert[A any](model any) (*mongo.InsertOneResult, error) {
	db := database.GetGoodmDatabase()
	return insertOne[A](db, model)
}

func InsertMany[A any](models []A) (*mongo.InsertManyResult, error) {

	var interfaceSlice = make([]any, len(models))

	for i := 0; i < len(interfaceSlice); i++ {
		interfaceSlice[i] = models[0]
	}

	db := database.GetGoodmDatabase()

	return insertMany[A](db, interfaceSlice)
}

func InsertWithDatabase[A any](db *database.GoodmDatabase, model any) (*mongo.InsertOneResult, error) {
	return insertOne[A](db, model)
}

func InsertManyWithDatabase[A any](db *database.GoodmDatabase, models []A) (*mongo.InsertManyResult, error) {
	var interfaceSlice = make([]any, len(models))

	for i := 0; i < len(interfaceSlice); i++ {
		interfaceSlice[i] = models[0]
	}
	return insertMany[A](db, interfaceSlice)
}

func insertOne[A any](db *database.GoodmDatabase, model any) (*mongo.InsertOneResult, error) {

	collection, err := database.GetCollection[A](db)

	if err != nil {
		return nil, err
	}

	res, err := collection.InsertOne(*db.Context, model)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func insertMany[A any](db *database.GoodmDatabase, models []any) (*mongo.InsertManyResult, error) {

	collection, err := database.GetCollection[A](db)

	if err != nil {
		return nil, err
	}

	res, err := collection.InsertMany(*db.Context, models)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func TestInsertOnePrivate[A any](db *database.GoodmDatabase, model any) (*mongo.InsertOneResult, error) {
	return insertOne[A](db, model)
}

func TestInsertManyPrivate[A any](db *database.GoodmDatabase, models []any) (*mongo.InsertManyResult, error) {
	return insertMany[A](db, models)
}

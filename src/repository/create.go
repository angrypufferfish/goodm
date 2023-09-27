package repository

import (
	"github.com/angrypufferfish/goodm/src/database"
	"go.mongodb.org/mongo-driver/mongo"
)

func Insert[A interface{}](model interface{}) (*mongo.InsertOneResult, error) {
	db := database.GetGoodmDatabase()
	return insertOne[A](db, model)
}

func InsertMany[A interface{}](models []A) (*mongo.InsertManyResult, error) {

	var interfaceSlice = make([]interface{}, len(models))

	for i := 0; i < len(interfaceSlice); i++ {
		interfaceSlice[i] = models[0]
	}

	db := database.GetGoodmDatabase()

	return insertMany[A](db, interfaceSlice)
}

func InsertWithDatabase[A interface{}](db *database.GoodmDatabase, model interface{}) (*mongo.InsertOneResult, error) {
	return insertOne[A](db, model)
}

func InsertManyWithDatabase[A interface{}](db *database.GoodmDatabase, models []A) (*mongo.InsertManyResult, error) {
	var interfaceSlice = make([]interface{}, len(models))

	for i := 0; i < len(interfaceSlice); i++ {
		interfaceSlice[i] = models[0]
	}
	return insertMany[A](db, interfaceSlice)
}

func insertOne[A interface{}](db *database.GoodmDatabase, model interface{}) (*mongo.InsertOneResult, error) {

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

func insertMany[A interface{}](db *database.GoodmDatabase, models []interface{}) (*mongo.InsertManyResult, error) {

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

func TestInsertOnePrivate[A interface{}](db *database.GoodmDatabase, model interface{}) (*mongo.InsertOneResult, error) {
	return insertOne[A](db, model)
}

func TestInsertManyPrivate[A interface{}](db *database.GoodmDatabase, models []interface{}) (*mongo.InsertManyResult, error) {
	return insertMany[A](db, models)
}

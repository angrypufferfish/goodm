package query

import (
	"github.com/angrypufferfish/goodm/src/database"
	"go.mongodb.org/mongo-driver/mongo"
)

func Insert[A interface{}](model interface{}) (*mongo.InsertOneResult, error) {
	db := database.GetGoodmDatabase()
	return insert[A](db, model)
}

func InsertWithDatabase[A interface{}](db *database.GoodmDatabase, model interface{}) (*mongo.InsertOneResult, error) {
	return insert[A](db, model)
}

func insert[A interface{}](db *database.GoodmDatabase, model interface{}) (*mongo.InsertOneResult, error) {

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

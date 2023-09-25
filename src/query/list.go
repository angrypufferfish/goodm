package query

import (
	"github.com/angrypufferfish/goodm/src/database"
	"github.com/angrypufferfish/goodm/src/serializer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func List[A interface{}](filter any, opts ...*options.FindOptions) ([]A, error) {
	db := database.GetGoodmDatabase()
	return list[A](db, filter, opts...)
}

func ListWithDatabase[A interface{}](db *database.GoodmDatabase, filter any, opts ...*options.FindOptions) ([]A, error) {

	return list[A](db, filter, opts...)
}

func list[A interface{}](db *database.GoodmDatabase, filter any, opts ...*options.FindOptions) ([]A, error) {

	collection, err := database.GetCollection[A](db)

	if err != nil {
		return nil, err
	}

	cursor, err := collection.Find(*db.Context, filter, opts...)

	var results []bson.M
	if err = cursor.All(*db.Context, &results); err != nil {
		return nil, err
	}
	return serializer.SerializeList[A](results), nil
}

package query

import (
	"github.com/angrypufferfish/goodm/src/database"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Delete[A interface{}](filter any, opts ...*options.DeleteOptions) (*int64, error) {
	db := database.GetGoodmDatabase()
	return delete[A](db, filter, opts...)
}

func DeleteWithDatabase[A interface{}](db *database.GoodmDatabase, filter any, opts ...*options.DeleteOptions) (*int64, error) {

	return delete[A](db, filter, opts...)
}

func delete[A interface{}](db *database.GoodmDatabase, filter any, opts ...*options.DeleteOptions) (*int64, error) {

	collection, err := database.GetCollection[A](db)

	if err != nil {
		return nil, err
	}

	count, err := collection.DeleteOne(*db.Context, filter, opts...)

	if err != nil {
		return nil, err
	}
	return &count.DeletedCount, nil
}

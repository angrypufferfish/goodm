package repository

import (
	"github.com/angrypufferfish/goodm/src/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Delete[A interface{}](id primitive.ObjectID) (*int64, error) {
	db := database.GetGoodmDatabase()
	objectID, err := objectIdConvert(id)

	if err != nil {
		return nil, err
	}
	return deleteOne[A](db, map[string]primitive.ObjectID{"_id": *objectID})
}

func DeleteOne[A interface{}](filter any, opts ...*options.DeleteOptions) (*int64, error) {
	db := database.GetGoodmDatabase()
	return deleteOne[A](db, filter, opts...)
}

func DeleteOneWithDatabase[A interface{}](db *database.GoodmDatabase, filter any, opts ...*options.DeleteOptions) (*int64, error) {

	return deleteOne[A](db, filter, opts...)
}

func DeleteMany[A interface{}](filter any, opts ...*options.DeleteOptions) (*int64, error) {
	db := database.GetGoodmDatabase()
	return deleteMany[A](db, filter, opts...)
}

func DeleteManyWithDatabase[A interface{}](db *database.GoodmDatabase, filter any, opts ...*options.DeleteOptions) (*int64, error) {
	return deleteMany[A](db, filter, opts...)
}

func deleteOne[A interface{}](db *database.GoodmDatabase, filter any, opts ...*options.DeleteOptions) (*int64, error) {

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

func deleteMany[A interface{}](db *database.GoodmDatabase, filter any, opts ...*options.DeleteOptions) (*int64, error) {

	collection, err := database.GetCollection[A](db)

	if err != nil {
		return nil, err
	}

	count, err := collection.DeleteMany(*db.Context, filter, opts...)

	if err != nil {
		return nil, err
	}
	return &count.DeletedCount, nil
}

func TestDeleteOnePrivate[A interface{}](db *database.GoodmDatabase, filter any, opts ...*options.DeleteOptions) (*int64, error) {
	return deleteOne[A](db, filter, opts...)
}

func TestDeleteManyPrivate[A interface{}](db *database.GoodmDatabase, filter any, opts ...*options.DeleteOptions) (*int64, error) {
	return deleteMany[A](db, filter, opts...)
}

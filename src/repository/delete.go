package repository

import (
	"github.com/angrypufferfish/goodm/src/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Delete[A any](id primitive.ObjectID) (*int64, error) {
	db := database.GetGoodmDatabase()
	objectID, err := ObjectIdConvert(id)

	if err != nil {
		return nil, err
	}
	return deleteOne[A](db, map[string]primitive.ObjectID{"_id": *objectID})
}

func DeleteOne[A any](filter any, opts ...*options.DeleteOptions) (*int64, error) {
	db := database.GetGoodmDatabase()
	return deleteOne[A](db, filter, opts...)
}

func DeleteOneWithDatabase[A any](db *database.GoodmDatabase, filter any, opts ...*options.DeleteOptions) (*int64, error) {

	return deleteOne[A](db, filter, opts...)
}

func DeleteMany[A any](filter any, opts ...*options.DeleteOptions) (*int64, error) {
	db := database.GetGoodmDatabase()
	return deleteMany[A](db, filter, opts...)
}

func DeleteManyWithDatabase[A any](db *database.GoodmDatabase, filter any, opts ...*options.DeleteOptions) (*int64, error) {
	return deleteMany[A](db, filter, opts...)
}

func findOneAndDelete[A any, S any](db *database.GoodmDatabase, filter any, opts ...*options.FindOneAndDeleteOptions) (*S, error) {

	var result S
	collection, err := database.GetCollection[A](db)

	if err != nil {
		return nil, err
	}

	singleResult := collection.FindOneAndDelete(*db.Context, filter, opts...)

	result, err = serializeSingleResult[S](singleResult)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func deleteOne[A any](db *database.GoodmDatabase, filter any, opts ...*options.DeleteOptions) (*int64, error) {

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

func deleteMany[A any](db *database.GoodmDatabase, filter any, opts ...*options.DeleteOptions) (*int64, error) {

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

func TestDeleteOnePrivate[A any](db *database.GoodmDatabase, filter any, opts ...*options.DeleteOptions) (*int64, error) {
	return deleteOne[A](db, filter, opts...)
}

func TestDeleteManyPrivate[A any](db *database.GoodmDatabase, filter any, opts ...*options.DeleteOptions) (*int64, error) {
	return deleteMany[A](db, filter, opts...)
}

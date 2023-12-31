package repository

import (
	"github.com/angrypufferfish/goodm/src/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateByID[A any](id primitive.ObjectID, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	db := database.GetGoodmDatabase()

	return updateByID[A, A](db, id, update, opts...)
}

func UpdateOne[A any](filter any, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	db := database.GetGoodmDatabase()

	return updateOne[A, A](db, filter, update, opts...)
}

func UpdateMany[A any](filter any, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	db := database.GetGoodmDatabase()

	return updateMany[A, A](db, filter, update, opts...)
}

func FindOneAndUpdate[A any](filter any, update any, opts ...*options.FindOneAndUpdateOptions) (*A, error) {
	db := database.GetGoodmDatabase()

	return findOneAndUpdate[A, A](db, filter, update, opts...)
}

func FindOneAndUpdateWithSerializer[A any, S any](filter any, update any, opts ...*options.FindOneAndUpdateOptions) (*S, error) {
	db := database.GetGoodmDatabase()

	return findOneAndUpdate[A, S](db, filter, update, opts...)
}

func findOneAndUpdate[A any, S any](db *database.GoodmDatabase, filter any, update any, opts ...*options.FindOneAndUpdateOptions) (*S, error) {

	var document *S
	collection, err := database.GetCollection[A](db)

	if err != nil {
		return nil, err
	}

	result := collection.FindOneAndUpdate(*db.Context, filter, update, opts...)

	if err != nil {
		return nil, err
	}

	document, err = serializeSingleResult[S](result)

	if err != nil {
		return nil, err
	}

	return document, nil
}

func updateByID[A any, S any](db *database.GoodmDatabase, id primitive.ObjectID, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {

	collection, err := database.GetCollection[A](db)

	if err != nil {
		return nil, err
	}

	objectID, err := ObjectIdConvert(id)
	if err != nil {
		return nil, err
	}

	result, err := collection.UpdateByID(*db.Context, objectID, update, opts...)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func updateOne[A any, S any](db *database.GoodmDatabase, filter any, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {

	collection, err := database.GetCollection[A](db)

	if err != nil {
		return nil, err
	}

	result, err := collection.UpdateOne(*db.Context, filter, update, opts...)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func updateMany[A any, S any](db *database.GoodmDatabase, filter any, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {

	collection, err := database.GetCollection[A](db)

	if err != nil {
		return nil, err
	}

	result, err := collection.UpdateMany(*db.Context, filter, update, opts...)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func TestUpdateManyPrivate[A any, S any](db *database.GoodmDatabase, filter any, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return updateMany[A, S](db, filter, update, opts...)
}

func TestUpdateOnePrivate[A any, S any](db *database.GoodmDatabase, filter any, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return updateOne[A, S](db, filter, update, opts...)
}

func TestUpdateByIDPrivate[A any, S any](db *database.GoodmDatabase, id primitive.ObjectID, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return updateByID[A, S](db, id, update, opts...)
}

package repository

import (
	"github.com/angrypufferfish/goodm/src/database"
	"github.com/angrypufferfish/goodm/src/serializer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Get[A interface{}](id primitive.ObjectID) (*A, error) {
	db := database.GetGoodmDatabase()
	objectID, err := objectIdConvert(id)
	if err != nil {
		return nil, err
	}
	return findOne[A, A](db, map[string]primitive.ObjectID{"_id": *objectID})
}

func FindOne[A interface{}](filter any, opts ...*options.FindOneOptions) (*A, error) {
	db := database.GetGoodmDatabase()
	return findOne[A, A](db, filter, opts...)
}

func Find[A interface{}](filter any, opts ...*options.FindOptions) ([]A, error) {
	db := database.GetGoodmDatabase()
	return find[A, A](db, filter, opts...)
}

func FindOneWithDatabase[A interface{}](db *database.GoodmDatabase, filter any, opts ...*options.FindOneOptions) (*A, error) {

	return findOne[A, A](db, filter, opts...)
}

func FindWithDatabase[A interface{}](db *database.GoodmDatabase, filter any, opts ...*options.FindOptions) ([]A, error) {

	return find[A, A](db, filter, opts...)
}

func FindWithSerializer[A interface{}, S interface{}](filter any, opts ...*options.FindOptions) ([]S, error) {
	db := database.GetGoodmDatabase()
	return find[A, S](db, filter, opts...)
}

func FindOneWithSerializer[A interface{}, S interface{}](filter any, opts ...*options.FindOneOptions) (*S, error) {
	db := database.GetGoodmDatabase()
	return findOne[A, S](db, filter, opts...)
}

func find[A interface{}, S interface{}](db *database.GoodmDatabase, filter any, opts ...*options.FindOptions) ([]S, error) {

	collection, err := database.GetCollection[A](db)

	if err != nil {
		return nil, err
	}

	cursor, err := collection.Find(*db.Context, filter, opts...)

	var results []bson.M
	if err = cursor.All(*db.Context, &results); err != nil {
		return nil, err
	}
	return serializer.SerializeList[S](results), nil
}

func findOne[A interface{}, S interface{}](db *database.GoodmDatabase, filter any, opts ...*options.FindOneOptions) (*S, error) {

	collection, err := database.GetCollection[A](db)

	if err != nil {
		return nil, err
	}

	singleResult := collection.FindOne(*db.Context, filter, opts...)

	var result bson.M
	err = singleResult.Decode(&result)

	if err != nil {
		return nil, err
	}

	var document S = serializer.Serialize[S](result)

	return &document, nil
}

package repository

import (
	"github.com/angrypufferfish/goodm/src/database"
	"github.com/angrypufferfish/goodm/src/serializer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Get[A any](id primitive.ObjectID) (*A, error) {
	db := database.GetGoodmDatabase()
	objectID, err := ObjectIdConvert(id)
	if err != nil {
		return nil, err
	}
	return findOne[A, A](db, map[string]primitive.ObjectID{"_id": *objectID})
}

func GetWithSerializer[A any, S any](id primitive.ObjectID) (*S, error) {
	db := database.GetGoodmDatabase()
	objectID, err := ObjectIdConvert(id)
	if err != nil {
		return nil, err
	}
	return findOne[A, S](db, map[string]primitive.ObjectID{"_id": *objectID})
}

func FindOne[A any](filter any, opts ...*options.FindOneOptions) (*A, error) {
	db := database.GetGoodmDatabase()
	return findOne[A, A](db, filter, opts...)
}

func Find[A any](filter any, opts ...*options.FindOptions) ([]A, error) {
	db := database.GetGoodmDatabase()
	return find[A, A](db, filter, opts...)
}

func FindOneWithDatabase[A any](db *database.GoodmDatabase, filter any, opts ...*options.FindOneOptions) (*A, error) {

	return findOne[A, A](db, filter, opts...)
}

func FindWithDatabase[A any](db *database.GoodmDatabase, filter any, opts ...*options.FindOptions) ([]A, error) {

	return find[A, A](db, filter, opts...)
}

func FindWithSerializer[A any, S any](filter any, opts ...*options.FindOptions) ([]S, error) {
	db := database.GetGoodmDatabase()
	return find[A, S](db, filter, opts...)
}

func FindOneWithSerializer[A any, S any](filter any, opts ...*options.FindOneOptions) (*S, error) {
	db := database.GetGoodmDatabase()
	return findOne[A, S](db, filter, opts...)
}

func find[A any, S any](db *database.GoodmDatabase, filter any, opts ...*options.FindOptions) ([]S, error) {

	collection, err := database.GetCollection[A](db)

	if err != nil {
		return nil, err
	}

	cursor, err := collection.Find(*db.Context, filter, opts...)

	var results []bson.M
	if err = cursor.All(*db.Context, &results); err != nil {
		return nil, err
	}

	seralizedList, err := serializer.SerializeList[S](results)

	if err != nil {
		return nil, err
	}
	return seralizedList, nil
}

func findOne[A any, S any](db *database.GoodmDatabase, filter any, opts ...*options.FindOneOptions) (*S, error) {

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

	seralizedDocument, err := serializer.Serialize[S](result)

	if err != nil {
		return nil, err
	}

	return &seralizedDocument, nil
}

func TestFindOnePrivate[A any, S any](db *database.GoodmDatabase, filter any, opts ...*options.FindOneOptions) (*S, error) {
	return findOne[A, S](db, filter, opts...)
}

func TestFindPrivate[A any, S any](db *database.GoodmDatabase, filter any, opts ...*options.FindOptions) ([]S, error) {
	return find[A, S](db, filter, opts...)
}

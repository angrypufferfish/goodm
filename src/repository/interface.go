package repository

import (
	"github.com/angrypufferfish/goodm/src/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository[A HasGoodmObjectID] struct{}

type RepositoryServices[A HasGoodmObjectID] interface {
	Get(id primitive.ObjectID) (*A, error)
	Find(filter any, opts ...*options.FindOptions) ([]A, error)
	FindOne(filter any, opts ...*options.FindOneOptions) (*A, error)
	FindOneWithDatabase(db *database.GoodmDatabase, filter any, opts ...*options.FindOneOptions) ([]A, error)
	FindWithDatabase(db *database.GoodmDatabase, filter any, opts ...*options.FindOptions) ([]A, error)

	InsertSelf(u HasGoodmObjectID) (*A, error)
	Insert(model interface{}) (*mongo.InsertOneResult, error)
	InsertMany(models []interface{}) (*mongo.InsertManyResult, error)
	InsertWithDatabase(db *database.GoodmDatabase, model interface{}) (*mongo.InsertOneResult, error)
	InsertManyWithDatabase(db *database.GoodmDatabase, models []interface{}) (*mongo.InsertManyResult, error)

	Delete(id string) (*int64, error)
	DeleteOne(filter any, opts ...*options.DeleteOptions) (*int64, error)
	DeleteOneWithDatabase(db *database.GoodmDatabase, filter any, opts ...*options.DeleteOptions) (*int64, error)
	DeleteMany(filter any, opts ...*options.DeleteOptions) (*int64, error)
	DeleteManyWithDatabase(db *database.GoodmDatabase, filter any, opts ...*options.DeleteOptions) (*int64, error)
}

func (repo *Repository[A]) Get(id primitive.ObjectID) (*A, error) {
	return Get[A](id)
}

func (repo *Repository[A]) Find(filter any, opts ...*options.FindOptions) ([]A, error) {
	return Find[A](filter, opts...)
}

func (repo *Repository[A]) FindOne(filter any, opts ...*options.FindOneOptions) (*A, error) {
	return FindOne[A](filter, opts...)
}

func (repo *Repository[A]) FindOneWithDatabase(db *database.GoodmDatabase, filter any, opts ...*options.FindOneOptions) (*A, error) {
	return FindOneWithDatabase[A](db, filter, opts...)
}

func (repo *Repository[A]) FindWithDatabase(db *database.GoodmDatabase, filter any, opts ...*options.FindOptions) ([]A, error) {
	return FindWithDatabase[A](db, filter, opts...)
}

func (repo *Repository[A]) InsertSelf(u HasGoodmObjectID) (A, error) {
	var doc, err = InsertSelf[A, A](u)
	return *doc, err
}

func (repo *Repository[A]) Insert(model interface{}) (*mongo.InsertOneResult, error) {
	return Insert[A](model)
}

func (repo *Repository[A]) InsertMany(models []interface{}) (*mongo.InsertManyResult, error) {
	return InsertMany[A](models)
}

func (repo *Repository[A]) InsertWithDatabase(db *database.GoodmDatabase, model interface{}) (*mongo.InsertOneResult, error) {
	return InsertWithDatabase[A](db, model)
}

func (repo *Repository[A]) InsertManyWithDatabase(db *database.GoodmDatabase, models []interface{}) (*mongo.InsertManyResult, error) {
	return InsertManyWithDatabase[A](db, models)
}

func (repo *Repository[A]) Delete(id primitive.ObjectID) (*int64, error) {
	return Delete[A](id)
}

func (repo *Repository[A]) DeleteOne(filter any, opts ...*options.DeleteOptions) (*int64, error) {
	return DeleteOne[A](filter, opts...)
}

func (repo *Repository[A]) DeleteOneWithDatabase(db *database.GoodmDatabase, filter any, opts ...*options.DeleteOptions) (*int64, error) {
	return DeleteOneWithDatabase[A](db, filter, opts...)
}

func (repo *Repository[A]) DeleteMany(filter any, opts ...*options.DeleteOptions) (*int64, error) {
	return DeleteMany[A](filter, opts...)
}

func (repo *Repository[A]) DeleteManyWithDatabase(db *database.GoodmDatabase, filter any, opts ...*options.DeleteOptions) (*int64, error) {
	return DeleteManyWithDatabase[A](db, filter, opts...)
}

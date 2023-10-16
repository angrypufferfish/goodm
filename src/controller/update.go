package controller

import (
	"github.com/angrypufferfish/goodm/src/base"
	"github.com/angrypufferfish/goodm/src/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Update[D any](document base.IBaseDocument, update any, opts ...*options.UpdateOptions) (*D, error) {

	document.OnUpdate()

	return updateById[D](document.GetID(), update, opts...)
}

func UpdateMany[D any](filter any, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {

	return repository.UpdateMany[D](filter, update, opts...)
}

func updateById[D any](id primitive.ObjectID, update any, opts ...*options.UpdateOptions) (*D, error) {

	_, err := repository.UpdateByID[D](id, update, opts...)
	if err != nil {
		return nil, err
	}

	return Get[D](id), nil
}

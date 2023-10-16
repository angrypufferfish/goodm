package controller

import (
	"fmt"
	"reflect"

	"github.com/angrypufferfish/goodm/src/base"
	"github.com/angrypufferfish/goodm/src/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Remove[D any](document base.IBaseDocument) *D {

	remove[D](document)

	expectedType := reflect.ValueOf(document).Elem().Interface()
	castedDocument := (expectedType).(D)

	return &castedDocument
}

func RemoveByID[D any](id primitive.ObjectID) bool {

	count, err := repository.Delete[D](id)

	if err != nil {
		panic(err)
	}

	return *count != 1
}

func DeleteMany[D any](filter any, opts ...*options.DeleteOptions) (*int64, error) {

	return deleteMany[D](filter, opts...)
}

func remove[D any](document base.IBaseDocument) {

	documentID := document.GetID()
	count, err := repository.Delete[D](documentID)

	if err != nil {
		panic(err)
	}

	if *count != 1 {
		textError := fmt.Sprintf("Document with %v not found", documentID)
		panic(textError)
	}
}

func deleteMany[D any](filter any, opts ...*options.DeleteOptions) (*int64, error) {

	count, err := repository.DeleteMany[D](filter, opts...)

	if err != nil {
		return nil, err
	}

	return count, nil
}

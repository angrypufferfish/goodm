package controller

import (
	"github.com/angrypufferfish/goodm/src/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Get[D any](id primitive.ObjectID) *D {

	document, err := repository.Get[D](id)

	if err != nil {
		panic(err)
	}

	return document
}

func FindOne[D any](filters any, opts ...*options.FindOneOptions) *D {

	document, err := repository.FindOne[D](filters, opts...)

	if err != nil {
		panic(err)
	}

	return document
}

func FindOneAndSerialize[D any, S any](filters any, opts ...*options.FindOneOptions) *S {

	document, err := repository.FindOneWithSerializer[D, S](filters, opts...)

	if err != nil {
		panic(err)
	}

	return document
}

func Find[D any](filters any, opts ...*options.FindOptions) []D {

	allDocuments, err := repository.Find[D](filters, opts...)

	if err != nil {
		panic(err)
	}

	return allDocuments
}

func FindAndSerialize[D any, S any](filters any, opts ...*options.FindOptions) []S {

	allDocuments, err := repository.FindWithSerializer[D, S](filters, opts...)

	if err != nil {
		panic(err)
	}

	return allDocuments
}

func FindAll[D any](opts ...*options.FindOptions) []D {

	allDocuments, err := repository.Find[D](map[string]string{}, opts...)

	if err != nil {
		panic(err)
	}

	return allDocuments
}

func FindAllAndSerialize[D any, S any](opts ...*options.FindOptions) []S {

	allDocuments, err := repository.FindWithSerializer[D, S](map[string]string{}, opts...)

	if err != nil {
		panic(err)
	}

	return allDocuments
}

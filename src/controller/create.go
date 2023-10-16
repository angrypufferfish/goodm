package controller

import (
	"github.com/angrypufferfish/goodm/src/base"
	"github.com/angrypufferfish/goodm/src/repository"
)

func Create[D any](document base.IBaseDocument) *D {

	///EVENT
	document.OnCreate()

	return create[D, D](document)
}

func CreateAndSerialize[D any, S any](document base.IBaseDocument) *S {

	///EVENT
	document.OnCreate()

	return create[D, S](document)
}

func create[D any, S any](document base.IBaseDocument) *S {

	res, err := repository.Insert[D](document)

	if err != nil {
		panic(err)
	}

	objectID, err := repository.ObjectIdConvert(res.InsertedID)

	if err != nil {
		panic(err)
	}

	savedDocument, err := repository.GetWithSerializer[D, S](*objectID)

	if err != nil {
		panic(err)
	}

	return savedDocument
}

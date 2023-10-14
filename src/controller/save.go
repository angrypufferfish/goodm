package controller

import (
	"github.com/angrypufferfish/goodm/src/database"
	"github.com/angrypufferfish/goodm/src/repository"
)

func Save[D any](document database.IBaseDocument) (*D, error) {

	///EVENT
	document.OnCreate()

	return save[D, D](document)
}

func SaveAndSerialize[D any, S any](document database.IBaseDocument) (*S, error) {

	///EVENT
	document.OnCreate()

	return save[D, S](document)
}

func save[D any, S any](document database.IBaseDocument) (*S, error) {

	res, err := repository.Insert[D](document)

	if err != nil {
		return nil, err
	}

	objectID, err := repository.ObjectIdConvert(res.InsertedID)

	if err != nil {
		return nil, err
	}

	savedDocument, err := repository.GetWithSerializer[D, S](*objectID)

	if err != nil {
		return nil, err
	}
	return savedDocument, nil
}

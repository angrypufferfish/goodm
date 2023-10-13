package controller

import (
	"github.com/angrypufferfish/goodm/src/database"
	"github.com/angrypufferfish/goodm/src/repository"
)

func Remove[D any](document database.IBaseDocument) (bool, error) {

	count, err := repository.Delete[D](*document.GetBaseDocumentID())

	if err != nil {
		return false, err
	}

	return *count == 1, nil
}

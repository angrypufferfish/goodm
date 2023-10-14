package controller

import (
	"github.com/angrypufferfish/goodm/src/repository"
)

func List[D any](filters any) ([]D, error) {

	allDocuments, err := repository.Find[D](filters)

	if err != nil {
		return nil, err
	}

	return allDocuments, nil
}

func ListAndSerialize[D any, S any](filters any) ([]S, error) {

	allDocuments, err := repository.FindWithSerializer[D, S](filters)

	if err != nil {
		return nil, err
	}

	return allDocuments, nil
}

func ListAll[D any]() ([]D, error) {

	allDocuments, err := repository.Find[D](map[string]string{})

	if err != nil {
		return nil, err
	}

	return allDocuments, nil
}

func ListAllAndSerialize[D any, S any]() ([]S, error) {

	allDocuments, err := repository.FindWithSerializer[D, S](map[string]string{})

	if err != nil {
		return nil, err
	}

	return allDocuments, nil
}

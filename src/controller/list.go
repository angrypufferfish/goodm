package controller

import (
	"github.com/angrypufferfish/goodm/src/repository"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func List[D any](filters any, opts ...*options.FindOptions) ([]D, error) {

	allDocuments, err := repository.Find[D](filters, opts...)

	if err != nil {
		return nil, err
	}

	return allDocuments, nil
}

func ListAndSerialize[D any, S any](filters any, opts ...*options.FindOptions) ([]S, error) {

	allDocuments, err := repository.FindWithSerializer[D, S](filters, opts...)

	if err != nil {
		return nil, err
	}

	return allDocuments, nil
}

func ListAll[D any](opts ...*options.FindOptions) ([]D, error) {

	allDocuments, err := repository.Find[D](map[string]string{}, opts...)

	if err != nil {
		return nil, err
	}

	return allDocuments, nil
}

func ListAllAndSerialize[D any, S any](opts ...*options.FindOptions) ([]S, error) {

	allDocuments, err := repository.FindWithSerializer[D, S](map[string]string{}, opts...)

	if err != nil {
		return nil, err
	}

	return allDocuments, nil
}

package database

import (
	"reflect"
	"strings"

	"github.com/angrypufferfish/goodm/src/exceptions"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCollectionName[A any]() (*string, error) {

	var collection A

	var field reflect.StructField
	var ok bool = false

	if reflect.ValueOf(collection).Kind() == reflect.Ptr {
		//TODO: RIMUOVERE STRINGA HARDOCODED
		field, ok = reflect.TypeOf(collection).Elem().FieldByName("BaseDocument")
	} else {
		var collectionPointer *A
		//TODO: RIMUOVERE STRINGA HARDOCODED
		field, ok = reflect.TypeOf(collectionPointer).Elem().FieldByName("BaseDocument")
	}

	if ok != true {
		panic(exceptions.ErrCollectionNameNotDefined)
	}
	//TODO RIMUOVERE STRINGA HARDOCODED
	collectionName := strings.Split(field.Tag.Get("goodm"), ",")[0]

	return &collectionName, nil
}

func GetCollection[A any](db *GoodmDatabase) (*mongo.Collection, error) {

	collectionName, err := GetCollectionName[A]()
	if err != nil {
		return nil, err
	}
	collection := db.UseCollection(*collectionName)

	return collection, nil
}

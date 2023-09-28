package database

import (
	"reflect"
	"strings"

	"github.com/angrypufferfish/goodm/src/exceptions"
	"go.mongodb.org/mongo-driver/mongo"
)

func getCollectionName[A any]() (*string, error) {

	var collection A

	var field reflect.StructField
	var ok bool = false

	if reflect.ValueOf(collection).Kind() == reflect.Ptr {
		field, ok = reflect.TypeOf(collection).Elem().FieldByName("GoodmCollection")
	} else {
		var collectionPointer *A
		field, ok = reflect.TypeOf(collectionPointer).Elem().FieldByName("GoodmCollection")
	}

	if ok != true {
		panic(exceptions.ErrCollectionNameNotDefined)
	}
	collectionName := strings.Split(field.Tag.Get("goodm"), ",")[0]

	return &collectionName, nil
}

func GetCollection[A any](db *GoodmDatabase) (*mongo.Collection, error) {

	collectionName, err := getCollectionName[A]()
	if err != nil {
		return nil, err
	}
	collection := db.UseCollection(*collectionName)

	return collection, nil
}

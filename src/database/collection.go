package database

import (
	"reflect"
	"strings"

	"github.com/angrypufferfish/goodm/src/exceptions"
	"go.mongodb.org/mongo-driver/mongo"
)

func getCollectionName[A interface{}]() (*string, error) {

	var collection A

	var field reflect.StructField
	var ok bool = false

	if reflect.ValueOf(collection).Kind() == reflect.Ptr {
		field, ok = reflect.TypeOf(collection).Elem().FieldByName("goodmCollection")
	} else {
		var collectionPointer *A
		field, ok = reflect.TypeOf(collectionPointer).Elem().FieldByName("goodmCollection")
	}

	if ok != true {
		panic(exceptions.ErrCollectionNameNotDefined)
	}
	collectionName := strings.Split(field.Tag.Get("goodm"), ",")[0]

	return &collectionName, nil
}

func GetCollection[A interface{}](db *GoodmDatabase) (*mongo.Collection, error) {

	collectionName, err := getCollectionName[A]()
	if err != nil {
		return nil, err
	}
	collection := db.UseCollection(*collectionName)

	return collection, nil
}

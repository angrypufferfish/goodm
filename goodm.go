package goodm

import (
	"github.com/angrypufferfish/goodm/src/base"
	"github.com/angrypufferfish/goodm/src/controller"
	"github.com/angrypufferfish/goodm/src/query"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/** CREATE FUNCTIONS **/

func Create[D any](document base.IBaseDocument) *D {
	return controller.Create[D](document)
}

func CreateAndSerialize[D any, S any](document base.IBaseDocument) *S {
	return controller.CreateAndSerialize[D, S](document)
}

/** UPDATE FUNCTIONS **/
func UpdateSelf[D any](document base.IBaseDocument, opts ...*options.UpdateOptions) (*D, error) {

	return controller.Update[D](document, query.Set(document), opts...)
}

func Update[D any](document base.IBaseDocument, update any, opts ...*options.UpdateOptions) (*D, error) {

	return controller.Update[D](document, update, opts...)
}

func UpdateMany[D any](filter any, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return controller.UpdateMany[D](filter, update, opts...)
}

/** FIND FUNCTIONS **/

func Get[D any](id primitive.ObjectID) *D {
	return controller.Get[D](id)
}

func FindOne[D any](filters any, opts ...*options.FindOneOptions) *D {
	return controller.FindOne[D](filters, opts...)
}

func Find[D any](filters any, opts ...*options.FindOptions) []D {
	return controller.Find[D](filters, opts...)
}

func FindAll[D any](opts ...*options.FindOptions) []D {
	return controller.FindAll[D](opts...)
}

func FindOneAndSerialize[D any, S any](filters any, opts ...*options.FindOneOptions) *S {
	return controller.FindOneAndSerialize[D, S](filters, opts...)
}

func FindAndSerialize[D any, S any](filters any, opts ...*options.FindOptions) []S {
	return controller.FindAndSerialize[D, S](filters, opts...)
}

func FindAllAndSerialize[D any, S any](opts ...*options.FindOptions) []S {
	return controller.FindAllAndSerialize[D, S](opts...)
}

/** REMOVE FUNCTIONS **/

func Remove[D any](document base.IBaseDocument) *D {
	return controller.Remove[D](document)
}

func RemoveByID[D any](id primitive.ObjectID) bool {
	return controller.RemoveByID[D](id)
}

func DeleteMany[D any](filter any, opts ...*options.DeleteOptions) (*int64, error) {
	return controller.DeleteMany[D](filter, opts...)
}

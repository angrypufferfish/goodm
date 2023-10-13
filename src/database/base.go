package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type IBaseDocument interface {
	GetBaseDocumentID() *primitive.ObjectID
}

type BaseDocument struct {
	Id primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
}

func (bd *BaseDocument) GetBaseDocumentID() *primitive.ObjectID {
	return &bd.Id
}

package database

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IBaseDocument interface {
	GetBaseDocumentID() *primitive.ObjectID
	OnCreate()
	OnUpdate()
}

type BaseDocument struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

func (bd *BaseDocument) GetBaseDocumentID() *primitive.ObjectID {
	return &bd.Id
}

func (bd *BaseDocument) OnCreate() {
	if bd.CreatedAt.IsZero() {
		bd.CreatedAt = time.Now().UTC()
	}
	bd.UpdatedAt = time.Now().UTC()
}

func (bd *BaseDocument) OnUpdate() {
	bd.UpdatedAt = time.Now().UTC()
}

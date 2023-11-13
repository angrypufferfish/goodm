package base

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseDocument struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

func (bd *BaseDocument) GetID() primitive.ObjectID {
	return bd.Id
}

func (bd *BaseDocument) GetCreatedAt() time.Time {
	return bd.CreatedAt
}
func (bd *BaseDocument) GetUpdatedAt() time.Time {
	return bd.UpdatedAt
}

func (bd *BaseDocument) OnCreate() {
	//Rimuovere evento
	// o gestire per multiple create
	bd.CreatedAt = time.Now().UTC()
	bd.UpdatedAt = time.Now().UTC()

}

func (bd *BaseDocument) OnUpdate() {
	//Rimuovere evento
	// o gestire per multipli update
	bd.UpdatedAt = time.Now().UTC()
}

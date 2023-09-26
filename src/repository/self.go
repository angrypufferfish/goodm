package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GoodmObjectID struct {
	Id primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
}

func (g *GoodmObjectID) GetID() primitive.ObjectID {
	return g.Id
}

func (g GoodmObjectID) SetID(hexID any) {
	objectID, err := objectIdConvert(hexID)
	if err != nil {
		panic(err)
	}
	g.Id = *objectID
}

type HasGoodmObjectID interface {
	GetID() primitive.ObjectID
	SetID(hexID any)
}

func InsertSelf[S any, A HasGoodmObjectID](u HasGoodmObjectID) (*A, error) {

	res, err := Insert[A](u)
	if err != nil {
		return nil, err
	}

	doc, err := Get[A](res.InsertedID.(primitive.ObjectID))

	if err != nil {
		return nil, err
	}

	objectId := (*doc).GetID()

	u.SetID(objectId)

	return doc, nil
}

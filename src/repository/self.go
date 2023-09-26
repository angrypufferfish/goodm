package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IDGetterSetter interface {
	SetID(id primitive.ObjectID)
	GetID() primitive.ObjectID
}

func NewGoodmDoc[T IDGetterSetter](document T) *Repo[T] {
	return &Repo[T]{
		Document: document,
	}
}

type GoodmCollection[A IDGetterSetter] struct {
	Id primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
}

func (u *GoodmCollection[A]) ToGoodmDoc(document A) *Repo[A] {
	return NewGoodmDoc[A](document)
}

func (u *GoodmCollection[A]) SetID(id primitive.ObjectID) {
	u.Id = id
}

func (u *GoodmCollection[A]) GetID() primitive.ObjectID {
	return u.Id
}

type Repo[A IDGetterSetter] struct {
	Repository[A] `json:"-" bson:"-"`
	Document      A `json:"-,inline" bson:"-,inline"`
}

func (r *Repo[A]) Save() (*A, error) {
	res, err := r.Insert(r.Document)
	if err != nil {
		return nil, err
	}

	r.Document.SetID(res.InsertedID.(primitive.ObjectID))

	document, err := r.Get(res.InsertedID.(primitive.ObjectID))
	if err != nil {
		return nil, err
	}
	return document, nil
}

func (r *Repo[A]) Remove() (bool, error) {
	_, err := r.Delete(r.Document.GetID())
	if err != nil {
		return false, err
	}
	return true, nil
}

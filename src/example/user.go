package example

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserData struct {
	Address string `json:"address" bson:"address"`
	Cap     string `json:"cap" bson:"cap"`
}

type User struct {
	goodmCollection string `json:"-" bson:"-" goodm:"users"`

	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	LastName  string             `json:"lastName" bson:"lastName"`
	FirstName string             `json:"firstName" bson:"firstName"`
	UserData  UserData           `json:"userData" bson:"userData"`
}

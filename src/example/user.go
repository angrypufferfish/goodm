package example

import (
	"github.com/angrypufferfish/goodm/src/database"
)

type UserData struct {
	Address string `json:"address" bson:"address"`
	Cap     string `json:"cap" bson:"cap"`
}

type User struct {
	database.BaseDocument `json:"inline" bson:"inline" goodm:"users"`

	LastName  string     `json:"lastName" bson:"lastName"`
	FirstName string     `json:"firstName" bson:"firstName"`
	Colors    []UserData `json:"colors" bson:"colors"`
	UserData  UserData   `json:"userData" bson:"userData"`
}

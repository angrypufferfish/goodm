package example

import (
	"github.com/angrypufferfish/goodm/src/repository"
)

type UserData struct {
	Address string `json:"address" bson:"address"`
	Cap     string `json:"cap" bson:"cap"`
}

type User struct {
	repository.Repository[*User] `json:"-" bson:"-"`
	goodmCollection              string `json:"-" bson:"-" goodm:"users"`

	repository.GoodmObjectID `json:"inline" bson:"inline"`
	LastName                 string   `json:"lastName" bson:"lastName"`
	FirstName                string   `json:"firstName" bson:"firstName"`
	UserData                 UserData `json:"userData" bson:"userData"`
}

func (u User) Save() *User {
	user, err := u.InsertSelf(&u)
	if err != nil {
		panic(err)
	}
	return user
}

func (u User) Remove() {
	u.Delete(u.Id)
}

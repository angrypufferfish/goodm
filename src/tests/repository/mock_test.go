package repository_test

import (
	"github.com/angrypufferfish/goodm/src/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserMockSerializer struct {
	Id   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"Name" bson:"Name"`
}

type UserMock struct {
	repository.GoodmCollection[*UserMock] `json:"inline" bson:"inline" goodm:"users"`

	Name      string  `json:"Name" bson:"Name"`
	Username  string  `json:"Username" bson:"Username"`
	MailCount int     `json:"MailCount,omitempty" bson:"MailCount,omitempty"`
	Latitude  float32 `json:"Latitude,omitempty" bson:"Latitude,omitempty"`
	Longitude float32 `json:"Longitude,omitempty" bson:"Longitude,omitempty"`
}

func NewUserMock() *UserMock {
	return &UserMock{
		Name:      "name",
		Username:  "username",
		MailCount: 1,
		Latitude:  12.33342,
		Longitude: 21.523423,
	}
}

var findOneMocks []*UserMock = []*UserMock{
	{
		Name:      "TEUF($/Q%HFJ?£Q=£/U=IQ)",
		Username:  "username",
		MailCount: 1,
		Latitude:  12.33342,
		Longitude: 21.523423,
	},
	{
		Name:      "name",
		Username:  "UF($/Q%HFJ?£Q=£/U=",
		MailCount: -23413,
		Latitude:  12.33342,
		Longitude: 21.523423,
	},
	{
		Name:      " NA     ME",
		Username:  " USERNAME ",
		Latitude:  12.33342,
		Longitude: 21.523423,
	},
	{
		Name:      "",
		Username:  "username",
		MailCount: 1,
		Latitude:  12.33342,
		Longitude: 21.523423,
	},
	{
		Username:  "username",
		MailCount: 1,
		Latitude:  12.33342,
		Longitude: 21.523423,
	},
	{
		Name:      "name",
		Username:  "username",
		MailCount: 100000000,
	},
}

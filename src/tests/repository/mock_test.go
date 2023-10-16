package repository_test

import (
	"github.com/angrypufferfish/goodm/src/base"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserMockSerializer struct {
	Id   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"Name" bson:"Name"`
}

type UserMock struct {
	base.BaseDocument `json:"inline" bson:"inline" goodm:"users"`

	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"Name" bson:"Name"`
	Username  string             `json:"Username" bson:"Username"`
	MailCount int                `json:"MailCount,omitempty" bson:"MailCount,omitempty"`
	Latitude  float32            `json:"Latitude,omitempty" bson:"Latitude,omitempty"`
	Longitude float32            `json:"Longitude,omitempty" bson:"Longitude,omitempty"`
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

var SelfSaveValues []UserMock = []UserMock{
	{
		Name:      "Sonja Campbell",
		Username:  "Pearl Hughes",
		Latitude:  30.371403,
		Longitude: -8.584232,
		MailCount: 92209,
	},
	{
		Name:      "Mason Marquez",
		Username:  "Heidi Herrera",
		Latitude:  -65.530518,
		Longitude: -112.318413,
		MailCount: 34914,
	},
	{
		Name:      "Rebecca Kirby",
		Username:  "Montgomery Ewing",
		Latitude:  78.015018,
		Longitude: -161.803062,
		MailCount: 64878,
	},
	{
		Name:      "Alejandra Bernard",
		Username:  "Herman Chapman",
		Latitude:  88.859209,
		Longitude: 3.663377,
		MailCount: 26219,
	},
	{
		Name:      "Simmons Moses",
		Username:  "Dunlap Lynn",
		Latitude:  78.762338,
		Longitude: 1.553304,
		MailCount: 4808,
	},
	{
		Name:      "Carrie Macias",
		Username:  "Ellis Bradford",
		Latitude:  74.353196,
		Longitude: -84.705472,
		MailCount: 64423,
	},
	{
		Name:      "Jamie Wiley",
		Username:  "Long Mccarty",
		Latitude:  -5.673011,
		Longitude: 92.366622,
		MailCount: 55380,
	},
	{
		Name:      "Price Downs",
		Username:  "Katie Good",
		Latitude:  -85.13695,
		Longitude: -103.904056,
		MailCount: 70636,
	},
	{
		Name:      "Coleman Reilly",
		Username:  "Marcella Morin",
		Latitude:  -9.235132,
		Longitude: -170.063634,
		MailCount: 99099,
	},
	{
		Name:      "Debbie Powers",
		Username:  "Lynn Bond",
		Latitude:  -52.024289,
		Longitude: -64.72688,
		MailCount: 77533,
	},
	{
		Name:      "Spears Roberson",
		Username:  "Vonda Solomon",
		Latitude:  73.574969,
		Longitude: 179.053217,
		MailCount: 59436,
	},
	{
		Name:      "Leonor Hobbs",
		Username:  "Burnett Cannon",
		Latitude:  32.264041,
		Longitude: 59.571375,
		MailCount: 96909,
	},
	{
		Name:      "Mcconnell Thornton",
		Username:  "Conner Carson",
		Latitude:  -43.694978,
		Longitude: -58.778376,
		MailCount: 24635,
	},
	{
		Name:      "Zimmerman Bonner",
		Username:  "Mai Pollard",
		Latitude:  56.962317,
		Longitude: 154.244795,
		MailCount: 68102,
	},
}

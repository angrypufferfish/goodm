package serializer_test

import (
	"testing"

	"github.com/angrypufferfish/goodm/src/serializer"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Name string `json:"name" bson:"name"`
	Age  int    `json:"age" bson:"age"`
}

func NewUser() *User {
	return &User{
		Name: "test",
		Age:  18,
	}
}

func TestSerialize(t *testing.T) {

	user, err := serializer.Serialize[User](primitive.D{
		{"Name", "test"},
		{"Age", 18},
	})

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.IsType(t, NewUser(), user)

}

func TestSerializeError(t *testing.T) {
	user, err := serializer.Serialize[User](primitive.D{
		{"Name", 23},
		{"Age", primitive.NewObjectID()},
	})

	assert.NotNil(t, err)
	assert.Nil(t, user)

	userNil, err := serializer.Serialize[User](nil)

	assert.NotNil(t, err)
	assert.Nil(t, userNil)
}

func TestSerializeList(t *testing.T) {
	users, err := serializer.SerializeList[User]([]primitive.D{
		{
			{"Name", "test"},
			{"Age", 18},
		},
		{
			{"Name", "test"},
			{"Age", 18},
		},
	})

	assert.Nil(t, err)
	assert.NotNil(t, users)
	assert.Len(t, users, 2)
	assert.IsType(t, []User{}, users)
}

func TestSerializeListError(t *testing.T) {
	users, err := serializer.SerializeList[User]([]primitive.D{
		{},
		{
			{"Name", 23},
			{"Age", 18},
		},
	})

	assert.NotNil(t, err)
	assert.Nil(t, users)

	usersNil, err := serializer.SerializeList[User]([]primitive.D{nil})

	assert.NotNil(t, err)
	assert.Nil(t, usersNil)
}

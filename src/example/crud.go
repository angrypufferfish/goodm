package example

import (
	"github.com/angrypufferfish/goodm/src/query"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserID struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	LastName string             `json:"lastName" bson:"lastName"`
}

func CreateUser() {
	user := &User{
		FirstName: "Dario",
		LastName:  "Treglia",
	}

	query.Insert[*User](user)
}

func ListAll() ([]User, error) {

	var filter = map[string]string{}

	r, err := query.Find[User](filter)

	if err != nil {
		return nil, err
	}
	return r, nil
}

func ListUserByFirstName(firstName string) ([]UserID, error) {

	var filter = map[string]string{"firstName": firstName}

	r, err := query.FindWithSerializer[User, UserID](filter)

	if err != nil {
		return nil, err
	}
	return r, nil
}

func DeleteUserByFirstName(firstName string) (*int64, error) {
	var filter = map[string]string{"firstName": firstName}

	deletedCount, err := query.DeleteOne[*User](filter)

	if err != nil {
		return nil, err
	}

	return deletedCount, nil
}

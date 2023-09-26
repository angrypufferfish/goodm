package example

import (
	"github.com/angrypufferfish/goodm/src/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserID struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	LastName string             `json:"lastName" bson:"lastName"`
}

func CreateUser() {

	user := repository.NewGoodmDoc[*User](
		&User{
			FirstName: "Mario",
			LastName:  "Rossi",
		},
	)
	user.Save()

	/** OR **/

	usr := &User{
		FirstName: "Mario",
		LastName:  "Rossi",
	}
	userInstance := usr.ToGoodmDoc(usr)
	userInstance.Save()

	user.Remove()
	userInstance.Remove()
}

func ListAll() ([]User, error) {

	var filter = map[string]string{}

	r, err := repository.Find[User](filter)

	if err != nil {
		return nil, err
	}
	return r, nil
}

func ListUserByFirstName(firstName string) ([]UserID, error) {

	var filter = map[string]string{"firstName": firstName}

	r, err := repository.FindWithSerializer[User, UserID](filter)

	if err != nil {
		return nil, err
	}
	return r, nil
}

func DeleteUserByFirstName(firstName string) (*int64, error) {
	var filter = map[string]string{"firstName": firstName}

	deletedCount, err := repository.DeleteOne[*User](filter)

	if err != nil {
		return nil, err
	}

	return deletedCount, nil
}

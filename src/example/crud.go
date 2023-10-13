package example

import (
	"fmt"

	"github.com/angrypufferfish/goodm/src/controller"
	"github.com/angrypufferfish/goodm/src/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserID struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	LastName string             `json:"lastName" bson:"lastName"`
}

func Fi(name bool) {

}

func CreateUser() {

	user := &User{
		FirstName: "Mario",
		LastName:  "Rossi",
		UserData: UserData{
			Cap:     "dwdwd",
			Address: "via ",
		},
	}
	//TODO: GET field bson tag name and value with reflect and implement search in List controller
	//example: ListAllBy(Fi(op.Or(op.Eq(user.FirstName == "name"),op.Neq(user.FirstName == "name"))))

	usr, _ := controller.Save[User](user)
	fmt.Printf("%v", usr)

	ok, _ := controller.Remove[User](usr)
	fmt.Printf("REMOVED %v", ok)

}

func ListAll() ([]User, error) {

	r, err := controller.ListAll[User]()

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

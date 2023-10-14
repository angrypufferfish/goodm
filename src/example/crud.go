package example

import (
	"fmt"

	"github.com/angrypufferfish/goodm/src/controller"
	op "github.com/angrypufferfish/goodm/src/operator"
	"github.com/angrypufferfish/goodm/src/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserID struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	LastName string             `json:"lastName" bson:"lastName"`
}

func CreateUser() {

	user := &User{
		FirstName: "Mario",
		LastName:  "Rossi",
		Colors: []UserData{
			{
				Cap:     "dwdwd",
				Address: "via ",
			},
		},
		UserData: UserData{
			Cap:     "dwdwd",
			Address: "via ",
		},
	}

	usr, err := controller.Save[User](user)

	if err != nil {
		panic(err)
	}

	founded, err := controller.List[User](op.ElemMatch("colors",
		op.And(
			op.Ne(
				"cap",
				"dwdw",
			),
			op.Eq(
				"cap",
				"dwdwd",
			),
		),
	))

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", founded)

	ok, _ := controller.Remove[User](usr)
	if ok {
		fmt.Print("Removed")
	}

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

package example

import (
	"fmt"

	"github.com/angrypufferfish/goodm"
	"github.com/angrypufferfish/goodm/src/query"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserID struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	LastName string             `json:"lastName" bson:"lastName"`
}

func Example() {
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

	goodm.DeleteMany[User](query.Exists("userData", true))

	user = goodm.Create[User](user)

	user.FirstName = "firstname"
	user.LastName = "lastname"

	goodm.UpdateSelf[User](user)

	users := goodm.FindOne[User](
		query.And(
			query.Ne("firstName", "Te"),
			query.Eq("lastName", "lastname"),
		),
	)

	fmt.Printf("USERS -> %v", users)
}

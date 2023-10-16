package example

import (
	"fmt"

	"github.com/angrypufferfish/goodm"
	"github.com/angrypufferfish/goodm/src/query"
	"go.mongodb.org/mongo-driver/bson"
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

	goodm.Update[User](
		user,
		query.Update(
			query.Set(bson.E{"firstName", "Tre"}, bson.E{"lastName", "Treglia"}),
			query.CurrentDate(bson.E{"updatedAt", true}),
		),
	)

	users := goodm.FindOne[User](
		query.Or(
			query.Eq("firstName", "Tre"),
			query.Eq("lastName", "d"),
		),
	)

	fmt.Printf("USERS -> %v", users)
}

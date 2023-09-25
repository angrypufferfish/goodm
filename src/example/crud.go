package example

import (
	"github.com/angrypufferfish/goodm/src/query"
)

func CreateUser() {
	user := &User{
		FirstName: "Dario",
		LastName:  "Treglia",
	}

	query.Insert[*User](user)
}

func ListAll() ([]User, error) {

	var filter = map[string]string{}

	r, err := query.List[User](filter)

	if err != nil {
		return nil, err
	}
	return r, nil
}

func ListUserByFirstName(firstName string) ([]User, error) {

	var filter = map[string]string{"firstName": firstName}

	r, err := query.List[User](filter)

	if err != nil {
		return nil, err
	}
	return r, nil
}

func DeleteUserByFirstName(firstName string) (*int64, error) {
	var filter = map[string]string{"firstName": firstName}

	deletedCount, err := query.Delete[*User](filter)

	if err != nil {
		return nil, err
	}

	return deletedCount, nil
}

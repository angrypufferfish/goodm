package example

import (
	"context"
	"fmt"

	"github.com/angrypufferfish/goodm/src/connection"
)

func Run() {

	var goodm connection.Goodm
	ctx := context.Background()

	client, err := goodm.Connect("mongodb://localhost:27017", 10000)
	defer goodm.Disconnect()

	if err != nil {
		panic(err)
	}
	client.UseDatabase("test", &ctx)

	CreateUser()

	users, _ := ListUserByFirstName("Dario")
	fmt.Printf("%v", users)

	//DeleteUserByFirstName("Dario")

	users, _ = ListUserByFirstName("Dario")
	fmt.Printf("%v", users)
}

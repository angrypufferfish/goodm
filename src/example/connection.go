package example

import (
	"context"

	"github.com/angrypufferfish/goodm"
)

func Run() {

	var goodm goodm.Goodm
	ctx := context.Background()

	client, err := goodm.Connect("mongodb://localhost:27017", 10000)
	defer goodm.Disconnect()

	if err != nil {
		panic(err)
	}

	client.UseDatabase("test", &ctx)

	Example()
}

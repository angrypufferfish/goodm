package connection

import (
	"context"
	"fmt"
	"time"

	"github.com/angrypufferfish/goodm/src/database"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Goodm struct {
	ctx    *context.Context
	client *database.GoodmClient
}

func (odm *Goodm) ping() (*database.GoodmClient, error) {
	fmt.Printf("Trying to Ping MongoDB\n")

	err := odm.client.Ping(*odm.ctx, readpref.Primary())
	return odm.client, err
}

func (odm *Goodm) connectWithTimeoutContext(timeout time.Duration) context.CancelFunc {
	connectionTimeout := timeout * time.Second
	tmpContext, cancel := context.WithTimeout(context.Background(), connectionTimeout)
	odm.ctx = &tmpContext

	return cancel
}

func (odm *Goodm) connectWithMongoUri(uri string) (*database.GoodmClient, error) {
	fmt.Printf("Trying to connect to MongoDB uri: %s\n", uri)
	client, err := mongo.Connect(*odm.ctx, options.Client().ApplyURI(uri))
	odm.client = database.NewGoodmClient(client)

	return odm.client, err
}

func (odm *Goodm) ConnectMock(mtestClient *mongo.Client) *database.GoodmClient {
	odm.client = database.NewGoodmClient(mtestClient)
	return odm.client
}

func (odm *Goodm) Connect(uri string, timeout time.Duration) (*database.GoodmClient, error) {

	cancel := odm.connectWithTimeoutContext(timeout)
	defer cancel()

	_, err := odm.connectWithMongoUri(uri)

	if err != nil {
		return nil, err
	}

	_, err = odm.ping()

	return odm.client, err
}

func (odm *Goodm) Disconnect() {

	if err := odm.client.Disconnect(*odm.ctx); err != nil {
		panic(err)
	}

}

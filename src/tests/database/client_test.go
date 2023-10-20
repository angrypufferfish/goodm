package database_test

import (
	"context"
	"testing"

	"github.com/angrypufferfish/goodm"
	"github.com/angrypufferfish/goodm/src/database"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestNewGoodmClient(t *testing.T) {

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	goodmClient := database.NewGoodmClient(mt.Client)

	assert.Equal(t, mt.Client, goodmClient.Client)
}

func TestGoodmClientUseDatabase(t *testing.T) {
	var iGoodm goodm.Goodm

	ctx := context.Background()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("UseDatabase", func(mt *mtest.T) {

		db := iGoodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		///CHECK GLOBAL DB IN USE
		dbInUse := database.GetGoodmDatabase()

		assert.Equal(t, dbInUse, db)
		assert.NotNil(t, dbInUse)
		assert.NotNil(t, db)
	})
}

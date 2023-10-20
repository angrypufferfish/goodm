package database_test

import (
	"context"
	"testing"

	"github.com/angrypufferfish/goodm"
	"github.com/angrypufferfish/goodm/src/database"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestNewGoodmDatabaseContext(t *testing.T) {
	var iGoodm goodm.Goodm
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	ctx := context.Background()
	defer mt.Close()

	mt.Run("UseDatabaseContext", func(mt *mtest.T) {

		goodmClient := iGoodm.ConnectMock(mt.Client)
		goodmDatabase := database.NewGoodmDatabase(goodmClient, "test", &ctx)

		assert.Equal(t, &ctx, goodmDatabase.Context)
	})
}

func TestUseCollection(t *testing.T) {
	var iGoodm goodm.Goodm
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	ctx := context.Background()
	defer mt.Close()

	mt.Run("UseDatabaseContext", func(mt *mtest.T) {

		goodmClient := iGoodm.ConnectMock(mt.Client)
		goodmDatabase := database.NewGoodmDatabase(goodmClient, "test", &ctx)

		current := goodmDatabase.UseCollection("test_collection")
		assert.NotNil(t, goodmDatabase.GetCurrenCollection())
		assert.Equal(t, current, goodmDatabase.GetCurrenCollection())

		goodmDatabase.UseCollection("test_collection_different")
		assert.NotEqual(t, current, goodmDatabase.GetCurrenCollection())

	})
}

package repository_test

import (
	"context"
	"testing"

	"github.com/angrypufferfish/goodm"
	"github.com/angrypufferfish/goodm/src/repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUpdate(t *testing.T) {

	var mock *UserMock = NewUserMock()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	var goodm goodm.Goodm
	ctx := context.Background()

	mt.Run("SUCCESS - UpdateOne", func(mt *mtest.T) {

		db := goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"nModified", 1},
		})

		res, err := repository.TestUpdateOnePrivate[UserMock, UserMock](db, map[string]string{}, bson.D{
			{"$set", mock},
		})

		assert.NotEmpty(t, (*res).ModifiedCount, "ModifiedCount not found")
		assert.Equal(t, (*res).ModifiedCount, int64(1))
		assert.Nil(t, err, "Error is NotNil")
	})

	mt.Run("Error - UpdateOne", func(mt *mtest.T) {

		db := goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		mt.AddMockResponses(bson.D{
			{"ok", 0},
			{"nModified", 1},
		})

		res, err := repository.TestUpdateOnePrivate[UserMock, UserMock](db, map[string]string{}, bson.D{
			{"$set", mock},
		})

		assert.Nil(t, res, "Res is not Nil")
		assert.NotNil(t, err, "Error is Nil")
	})

	mt.Run("SUCCESS - UpdateMany", func(mt *mtest.T) {

		db := goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"nModified", 5},
		})

		res, err := repository.TestUpdateManyPrivate[UserMock, UserMock](db, map[string]string{}, bson.D{
			{"$set", mock},
		})

		assert.NotEmpty(t, (*res).ModifiedCount, "ModifiedCount not found")
		assert.Equal(t, (*res).ModifiedCount, int64(5))
		assert.Nil(t, err, "Error is NotNil")
	})

	mt.Run("Error - UpdateMany", func(mt *mtest.T) {

		db := goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		mt.AddMockResponses(bson.D{
			{"ok", 0},
			{"nModified", 5},
		})

		res, err := repository.TestUpdateManyPrivate[UserMock, UserMock](db, map[string]string{}, bson.D{
			{"$set", mock},
		})

		assert.Nil(t, res, "Res is not Nil")
		assert.NotNil(t, err, "Error is Nil")
	})

	mt.Run("SUCCESS - UpdateByID", func(mt *mtest.T) {

		db := goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"nModified", 1},
		})

		res, err := repository.TestUpdateByIDPrivate[UserMock, UserMock](db, primitive.NewObjectID(), bson.D{
			{"$set", mock},
		})

		assert.NotEmpty(t, (*res).ModifiedCount, "ModifiedCount not found")
		assert.Equal(t, (*res).ModifiedCount, int64(1))
		assert.Nil(t, err, "Error is NotNil")
	})

	mt.Run("Error - UpdateByID", func(mt *mtest.T) {

		db := goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		mt.AddMockResponses(bson.D{
			{"ok", 0},
			{"nModified", 1},
		})

		res, err := repository.TestUpdateByIDPrivate[UserMock, UserMock](db, primitive.NewObjectID(), bson.D{
			{"$set", mock},
		})

		assert.Nil(t, res, "Res is not Nil")
		assert.NotNil(t, err, "Error is Nil")
	})

	mt.Run("SUCCESS - Public UpdateByID", func(mt *mtest.T) {

		goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"nModified", 1},
		})

		res, err := repository.UpdateByID[UserMock](primitive.NewObjectID(), bson.D{
			{"$set", mock},
		})

		assert.NotEmpty(t, (*res).ModifiedCount, "ModifiedCount not found")
		assert.Equal(t, (*res).ModifiedCount, int64(1))
		assert.Nil(t, err, "Error is NotNil")
	})

	mt.Run("Error - Public UpdateByID", func(mt *mtest.T) {

		goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		mt.AddMockResponses(bson.D{
			{"ok", 0},
			{"nModified", 1},
		})

		res, err := repository.UpdateByID[UserMock](primitive.NewObjectID(), bson.D{
			{"$set", mock},
		})

		assert.Nil(t, res, "Res is not Nil")
		assert.NotNil(t, err, "Error is Nil")
	})
}

package repository_test

import (
	"context"
	"testing"

	"github.com/angrypufferfish/goodm/src/connection"
	"github.com/angrypufferfish/goodm/src/repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestInsert(t *testing.T) {

	var mock *UserMock = NewUserMock()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	var goodm connection.Goodm
	ctx := context.Background()

	mt.Run("SUCCESS - Insert", func(mt *mtest.T) {

		goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		res, err := repository.Insert[UserMock](mock)

		assert.NotEmpty(t, res.InsertedID, "InsertedID not found")
		assert.Nil(t, err, "Error is NotNil")
	})

	mt.Run("EXPECTED ERROR - Insert", func(mt *mtest.T) {

		goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})

		res, err := repository.Insert[UserMock](&UserMock{})

		assert.Nil(t, res, "Mongo.Insert is NotNil")
		assert.NotNil(t, err, "Error is Nil")
		assert.NotPanics(t, func() { repository.Insert[UserMock](&UserMock{}) })
	})

	mt.Run("SUCCESS - InsertWithDatabase", func(mt *mtest.T) {

		db := goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		res, err := repository.InsertWithDatabase[UserMock](db, mock)

		assert.NotEmpty(t, res.InsertedID, "InsertedID not found")
		assert.Nil(t, err, "Error is NotNil")
	})

	mt.Run("EXPECTED ERROR - InsertWithDatabase", func(mt *mtest.T) {

		db := goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})

		res, err := repository.InsertWithDatabase[UserMock](db, &UserMock{})

		assert.Nil(t, res, "Mongo.InsertWithDatabase is NotNil")
		assert.NotNil(t, err, "Error is Nil")
		assert.NotPanics(t, func() { repository.InsertWithDatabase[UserMock](db, &UserMock{}) })
	})

	mt.Run("SUCCESS - TestInsertOnePrivate", func(mt *mtest.T) {

		db := goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		res, err := repository.TestInsertOnePrivate[UserMock](db, mock)

		assert.NotEmpty(t, res.InsertedID, "InsertedID not found")
		assert.Nil(t, err, "Error is NotNil")
	})

	mt.Run("EXPECTED ERROR - TestInsertOnePrivate", func(mt *mtest.T) {

		db := goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})

		res, err := repository.TestInsertOnePrivate[UserMock](db, &UserMock{})

		assert.Nil(t, res, "Mongo.TestInsertOnePrivate is NotNil")
		assert.NotNil(t, err, "Error is Nil")
		assert.NotPanics(t, func() { repository.TestInsertOnePrivate[UserMock](db, &UserMock{}) })
	})
}

func TestInsertMany(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	var goodm connection.Goodm
	ctx := context.Background()

	mt.Run("SUCCESS - InsertMany", func(mt *mtest.T) {

		goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		users := []UserMock{
			*NewUserMock(),
			*NewUserMock(),
		}
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		res, err := repository.InsertMany[UserMock](users)

		assert.NotNil(t, res, "Res is Nil")
		assert.NotNil(t, res.InsertedIDs, "InsertedIDs is Nil")
		assert.Equal(t, len(users), len(res.InsertedIDs), "InsertedIDs is not equal to expected value")
		assert.Nil(t, err)
	})

	mt.Run("EXPECTED ERROR - InsertMany", func(mt *mtest.T) {

		goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		users := []UserMock{
			*NewUserMock(),
			*NewUserMock(),
			*NewUserMock(),
			*NewUserMock(),
		}
		mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})

		res, err := repository.InsertMany[UserMock](users)

		assert.NotNil(t, err)
		assert.Nil(t, res, "Res is Nil")
		assert.NotPanics(t, func() { repository.InsertMany[UserMock](users) })
	})

	mt.Run("SUCCESS - InsertManyWithDatabase", func(mt *mtest.T) {

		db := goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		users := []UserMock{
			*NewUserMock(),
			*NewUserMock(),
		}
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		res, err := repository.InsertManyWithDatabase[UserMock](db, users)

		assert.NotNil(t, res, "Res is Nil")
		assert.NotNil(t, res.InsertedIDs, "InsertedIDs is Nil")
		assert.Equal(t, len(users), len(res.InsertedIDs), "InsertedIDs is not equal to expected value")
		assert.Nil(t, err)
	})

	mt.Run("EXPECTED ERROR - InsertManyWithDatabase", func(mt *mtest.T) {

		db := goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		users := []UserMock{
			*NewUserMock(),
			*NewUserMock(),
			*NewUserMock(),
			*NewUserMock(),
		}
		mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})

		res, err := repository.InsertManyWithDatabase[UserMock](db, users)

		assert.NotNil(t, err)
		assert.Nil(t, res, "Res is Nil")
		assert.NotPanics(t, func() { repository.InsertManyWithDatabase[UserMock](db, users) })
	})

	mt.Run("SUCCESS - TestInsertManyPrivate", func(mt *mtest.T) {

		db := goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		users := []any{
			*NewUserMock(),
			*NewUserMock(),
		}
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		res, err := repository.TestInsertManyPrivate[UserMock](db, users)

		assert.NotNil(t, res, "Res is Nil")
		assert.NotNil(t, res.InsertedIDs, "InsertedIDs is Nil")
		assert.Equal(t, len(users), len(res.InsertedIDs), "InsertedIDs is not equal to expected value")
		assert.Nil(t, err)
	})

	mt.Run("EXPECTED ERROR - TestInsertManyPrivate", func(mt *mtest.T) {

		db := goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		users := []any{
			*NewUserMock(),
			*NewUserMock(),
			*NewUserMock(),
			*NewUserMock(),
		}
		mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})

		res, err := repository.TestInsertManyPrivate[UserMock](db, users)

		assert.NotNil(t, err)
		assert.Nil(t, res, "Res is Nil")
		assert.NotPanics(t, func() { repository.TestInsertManyPrivate[UserMock](db, users) })
	})
}

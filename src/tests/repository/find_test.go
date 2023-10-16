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

func TestFindOne(t *testing.T) {

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	var goodm goodm.Goodm
	ctx := context.Background()

	mt.Run("SUCCESS - TestFindOnePrivate", func(mt *mtest.T) {

		db := goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		expectedUser := &UserMock{
			Id:       primitive.NewObjectID(),
			Name:     "test*_*name",
			Username: "UserNameTest",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(0, "mock_db.users", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: expectedUser.Id},
			{Key: "Name", Value: expectedUser.Name},
			{Key: "Username", Value: expectedUser.Username},
		}))

		res, err := repository.TestFindOnePrivate[UserMock, UserMock](db, expectedUser)

		assert.Nil(t, err, "Error is NotNil")
		assert.Equal(t, *expectedUser, *res)
		assert.NotPanics(t, func() { repository.TestFindOnePrivate[UserMock, UserMock](db, expectedUser) })
	})

	mt.Run("SUCCESS - TestFindOnePrivate with serializer", func(mt *mtest.T) {

		db := goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		expectedUser := &UserMock{
			Id:       primitive.NewObjectID(),
			Name:     "jo hn",
			Username: "username with space",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(0, "mock_db.users", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: expectedUser.Id},
			{Key: "Name", Value: expectedUser.Name},
			{Key: "Username", Value: expectedUser.Username},
		}))

		res, err := repository.TestFindOnePrivate[UserMock, UserMockSerializer](db, expectedUser)

		assert.Nil(t, err, "Error is NotNil")
		assert.NotEmpty(t, (*res).Id)
		assert.Equal(t, expectedUser.Id, (*res).Id)
		assert.Equal(t, expectedUser.Name, (*res).Name)

		assert.NotPanics(t, func() { repository.TestFindOnePrivate[UserMock, UserMock](db, expectedUser) })
	})

	mt.Run("EXPECTED ERROR - TestFindOnePrivate", func(mt *mtest.T) {

		db := goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		expectedUser := UserMock{
			Name:     "errorname",
			Username: "Errorusername",
		}

		mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})

		res, err := repository.TestFindOnePrivate[UserMock, UserMock](db, expectedUser)

		assert.NotNil(t, err, "Error is Nil")
		assert.Nil(t, res)
		assert.NotPanics(t, func() { repository.TestFindOnePrivate[UserMock, UserMock](db, expectedUser) })
	})

}

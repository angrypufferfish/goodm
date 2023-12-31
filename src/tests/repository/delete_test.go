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

func TestDeleteOne(t *testing.T) {

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	var goodm goodm.Goodm
	ctx := context.Background()

	mt.Run("SUCCESS - TestDeleteOnePrivate", func(mt *mtest.T) {

		db := goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		expectedUser := &UserMock{
			Id:       primitive.NewObjectID(),
			Name:     "test*_*name",
			Username: "UserNameTest",
		}

		dataTest := [](map[int]bson.D){
			{1: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 1}}},
			{0: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 2}}},
			{0: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 142}}},
			{2: bson.D{{Key: "ok", Value: 0}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 1}}},
			{1: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 1}}},
		}

		for _, bs := range dataTest {
			for expected, b := range bs {

				mt.AddMockResponses(b)

				res, err := repository.TestDeleteOnePrivate[UserMock](db, map[string]primitive.ObjectID{
					"_id": expectedUser.Id,
				})

				if expected == 1 {
					assert.Nil(t, err)
					assert.Equal(t, *res, int64(1))
				} else if expected == 2 {
					assert.NotNil(t, err)
				} else {
					assert.NotEqual(t, *res, int64(1))
				}
			}
		}

	})
}

func TestDeleteMany(t *testing.T) {

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	var iGoodm goodm.Goodm
	ctx := context.Background()

	mt.Run("SUCCESS - TestDeleteOnePrivate", func(mt *mtest.T) {

		db := iGoodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		expectedUser := &UserMock{
			Name:     "test*_*name",
			Username: "UserNameTest",
		}

		dataTest := [](map[int]bson.D){
			{2: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 2}}},
			{142: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 142}}},
			{0: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 0}}},
			{1: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 1}}},
		}

		for _, bs := range dataTest {
			for expected, b := range bs {

				mt.AddMockResponses(b)

				res, err := repository.TestDeleteOnePrivate[UserMock](db, map[string]string{
					"Name": expectedUser.Name,
				})

				assert.Nil(t, err)
				assert.Equal(t, *res, int64(expected))

			}
		}

	})

	mt.Run("SUCCESS - TestDeleteManyPrivate", func(mt *mtest.T) {

		db := iGoodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		expectedUser := &UserMock{
			Name:     "test*_*name",
			Username: "UserNameTest",
		}

		dataTest := [](map[int]bson.D){
			{2: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 2}}},
			{142: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 142}}},
			{0: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 0}}},
			{1: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 1}}},
		}

		for _, bs := range dataTest {
			for expected, b := range bs {

				mt.AddMockResponses(b)

				res, err := repository.TestDeleteManyPrivate[UserMock](db, map[string]string{
					"Name": expectedUser.Name,
				})

				assert.Nil(t, err)
				assert.Equal(t, *res, int64(expected))

			}
		}

	})

	mt.Run("SUCCESS - DeleteOne", func(mt *mtest.T) {

		iGoodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		expectedUser := &UserMock{
			Name:     "test*_*name",
			Username: "UserNameTest",
		}

		dataTest := [](map[int]bson.D){
			{2: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 2}}},
			{142: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 142}}},
			{0: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 0}}},
			{1: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 1}}},
		}

		for _, bs := range dataTest {
			for expected, b := range bs {

				mt.AddMockResponses(b)

				res, err := repository.DeleteOne[UserMock](map[string]string{
					"Name": expectedUser.Name,
				})

				assert.Nil(t, err)
				assert.Equal(t, *res, int64(expected))

			}
		}

	})

	mt.Run("SUCCESS - Delete", func(mt *mtest.T) {

		iGoodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		expectedUser := &UserMock{
			Name:     "test*_*name",
			Username: "UserNameTest",
		}

		dataTest := [](map[int]bson.D){
			{2: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 2}}},
			{142: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 142}}},
			{0: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 0}}},
			{1: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 1}}},
		}

		for _, bs := range dataTest {
			for expected, b := range bs {

				mt.AddMockResponses(b)

				res, err := repository.Delete[UserMock](expectedUser.Id)

				assert.Nil(t, err)
				assert.Equal(t, *res, int64(expected))

			}
		}

	})

	mt.Run("SUCCESS - DeleteMany", func(mt *mtest.T) {

		iGoodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		expectedUser := &UserMock{
			Name:     "test*_*name",
			Username: "UserNameTest",
		}

		dataTest := [](map[int]bson.D){
			{2: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 2}}},
			{142: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 142}}},
			{0: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 0}}},
			{1: bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 1}}},
		}

		for _, bs := range dataTest {
			for expected, b := range bs {

				mt.AddMockResponses(b)

				res, err := repository.DeleteMany[UserMock](map[string]string{
					"Name": expectedUser.Name,
				})

				assert.Nil(t, err)
				assert.Equal(t, *res, int64(expected))

			}
		}

	})

}

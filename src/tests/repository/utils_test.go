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

func TestObjectIdConvert(t *testing.T) {

	obj := primitive.NewObjectID()
	var ids []any = []any{
		primitive.NewObjectID(),
		primitive.NewObjectID().Hex(),
		&obj,
	}

	var idsError []any = []any{
		nil,
		"NOT_OBJECT_ID",
		0,
		1,
		true,
		false,
	}

	for i := 0; i < len(ids); i++ {
		objectId, err := repository.ObjectIdConvert(ids[i])

		assert.Nil(t, err)
		assert.NotNil(t, objectId)
		expectedType := primitive.NewObjectID()
		assert.IsType(t, &expectedType, objectId)
	}

	for i := 0; i < len(idsError); i++ {
		objectId, err := repository.ObjectIdConvert(idsError[i])

		assert.NotNil(t, err)
		assert.Nil(t, objectId)
	}
}

func TestSerializeSingleResult(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	var goodm goodm.Goodm
	ctx := context.Background()
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

		res := db.UseCollection("users").FindOne(ctx, map[string]string{})

		serializedUser, err := repository.TestSerializeSingleResultPrivate[UserMockSerializer](res)

		assert.Nil(t, err, "Error is NotNil")
		assert.NotEmpty(t, (*serializedUser).Id)
		assert.Equal(t, expectedUser.Id, (*serializedUser).Id)
		assert.Equal(t, expectedUser.Name, (*serializedUser).Name)
		assert.IsType(t, &UserMockSerializer{}, serializedUser)
	})
}

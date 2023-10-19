package repository_test

import (
	"testing"

	"github.com/angrypufferfish/goodm/src/repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

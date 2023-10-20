package base_test

import (
	"testing"
	"time"

	"github.com/angrypufferfish/goodm/src/base"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestBaseDocument(t *testing.T) {
	id := primitive.NewObjectID()
	currentDate := time.Now().UTC()

	base := &base.BaseDocument{
		Id:        id,
		CreatedAt: currentDate,
		UpdatedAt: currentDate,
	}

	assert.Equal(t, id, base.Id)
	assert.Equal(t, currentDate, base.CreatedAt)
	assert.Equal(t, currentDate, base.UpdatedAt)

	assert.Equal(t, base.GetID(), id)
	assert.Equal(t, base.GetCreatedAt(), currentDate)
	assert.Equal(t, base.GetUpdatedAt(), currentDate)
}

func TestBaseDocumentOnCreateEvent(t *testing.T) {

	base := &base.BaseDocument{}
	base.OnCreate()

	assert.NotEmpty(t, base.GetCreatedAt())
	assert.NotEmpty(t, base.GetUpdatedAt())
}

func TestBaseDocumentOnUpdateEvent(t *testing.T) {

	base := &base.BaseDocument{}
	base.OnUpdate()

	assert.Empty(t, base.GetCreatedAt())
	assert.NotEmpty(t, base.GetUpdatedAt())
}

func TestBaseDocumentOnUpdateEventWithCreatedAtNotNil(t *testing.T) {
	currentDate := time.Now().UTC()
	base := &base.BaseDocument{
		CreatedAt: currentDate,
		UpdatedAt: currentDate,
	}
	base.OnUpdate()

	assert.NotEqual(t, base.GetCreatedAt(), base.GetUpdatedAt())
}

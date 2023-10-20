package database_test

import (
	"testing"

	"github.com/angrypufferfish/goodm/src/base"
	"github.com/angrypufferfish/goodm/src/database"
	"github.com/stretchr/testify/assert"
)

type InfoCollection struct {
	base.BaseDocument `json:",inline" bson:",inline" goodm:"info"`
}

type InfoWrongCollection struct {
	Name string `json:",inline" bson:",inline" goodm:""`
}

type Info struct {
	Name string `json:",inline" bson:",inline"`
}

func TestGetCollectionName(t *testing.T) {

	expectedName := "info"

	collectionName, err := database.GetCollectionName[InfoCollection]()

	assert.Nil(t, err)
	assert.NotNil(t, collectionName)
	assert.Equal(t, expectedName, *collectionName)
}

func TestGetCollectionNameError(t *testing.T) {

	///TEST WITHOUT TAG
	assert.Panics(t, func() {
		database.GetCollectionName[Info]()
	})

	///TEST WITH EMPTY TAG
	assert.Panics(t, func() {
		database.GetCollectionName[InfoWrongCollection]()
	})

}

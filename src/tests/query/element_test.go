package query_test

import (
	"testing"

	op_cmd "github.com/angrypufferfish/goodm/src/operators"
	"github.com/angrypufferfish/goodm/src/query"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestExists(t *testing.T) {
	field := "fieldToCheck"
	value1 := false
	value2 := true
	value3 := true

	expression1 := query.Exists(field, value1)
	expression2 := query.Exists(field, value2)
	expression3 := query.Exists(field, value3)

	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.EXISTS, value1}}}}, expression1)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.EXISTS, value2}}}}, expression2)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.EXISTS, value3}}}}, expression3)
}

func TestType(t *testing.T) {
	field := "fieldToCheck"

	value1 := bson.TypeBoolean
	value2 := bson.TypeObjectID
	value3 := bson.TypeString

	expression1 := query.Type(field, value1)
	expression2 := query.Type(field, value2)
	expression3 := query.Type(field, value3)

	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.TYPE, value1}}}}, expression1)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.TYPE, value2}}}}, expression2)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.TYPE, value3}}}}, expression3)
}

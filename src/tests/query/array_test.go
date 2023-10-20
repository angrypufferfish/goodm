package query_test

import (
	"testing"

	op_cmd "github.com/angrypufferfish/goodm/src/operators"
	"github.com/angrypufferfish/goodm/src/query"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestAll(t *testing.T) {
	field := "fieldToCheck"
	value1 := false
	value2 := true

	expression1 := query.All(field, value1)
	expression2 := query.All(field, value2)

	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.ALL, value1}}}}, expression1)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.ALL, value2}}}}, expression2)
}

func TestElemMatch(t *testing.T) {
	field := "fieldToCheck"
	value1 := "matchTarget1"
	value2 := "matchTarget2"

	expression1 := query.ElemMatch(field, value1)
	expression2 := query.ElemMatch(field, value2)

	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.ELEM_MATCH, value1}}}}, expression1)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.ELEM_MATCH, value2}}}}, expression2)
}

func TestSize(t *testing.T) {
	field := "fieldToCheck"
	value1 := false
	value2 := true

	expression1 := query.Size(field, value1)
	expression2 := query.Size(field, value2)

	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.SIZE, value1}}}}, expression1)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.SIZE, value2}}}}, expression2)
}

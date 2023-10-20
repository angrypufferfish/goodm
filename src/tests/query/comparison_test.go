package query_test

import (
	"testing"

	op_cmd "github.com/angrypufferfish/goodm/src/operators"
	"github.com/angrypufferfish/goodm/src/query"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestEq(t *testing.T) {
	field := "fieldToCheck"
	value1 := "ValueToMatch"
	value2 := true
	value3 := 23

	expression1 := query.Eq(field, value1)
	expression2 := query.Eq(field, value2)
	expression3 := query.Eq(field, value3)

	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.EQ, value1}}}}, expression1)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.EQ, value2}}}}, expression2)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.EQ, value3}}}}, expression3)
}

func TestGt(t *testing.T) {
	field := "fieldToCheck"
	value1 := -20
	value2 := 46
	value3 := 23

	expression1 := query.Gt(field, value1)
	expression2 := query.Gt(field, value2)
	expression3 := query.Gt(field, value3)

	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.GT, value1}}}}, expression1)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.GT, value2}}}}, expression2)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.GT, value3}}}}, expression3)
}

func TestGte(t *testing.T) {
	field := "fieldToCheck"
	value1 := -20
	value2 := 46
	value3 := 23

	expression1 := query.Gte(field, value1)
	expression2 := query.Gte(field, value2)
	expression3 := query.Gte(field, value3)

	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.GTE, value1}}}}, expression1)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.GTE, value2}}}}, expression2)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.GTE, value3}}}}, expression3)
}

func TestIn(t *testing.T) {
	field := "fieldToCheck"
	value1 := bson.A{"2", "23", "12"}
	value2 := bson.A{1, 2, -3}
	value3 := bson.A{"abcd", "Json", "Bson"}

	expression1 := query.In(field, value1)
	expression2 := query.In(field, value2)
	expression3 := query.In(field, value3)

	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.IN, value1}}}}, expression1)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.IN, value2}}}}, expression2)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.IN, value3}}}}, expression3)
}

func TestLt(t *testing.T) {
	field := "fieldToCheck"
	value1 := -20
	value2 := 46
	value3 := 23

	expression1 := query.Lt(field, value1)
	expression2 := query.Lt(field, value2)
	expression3 := query.Lt(field, value3)

	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.LT, value1}}}}, expression1)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.LT, value2}}}}, expression2)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.LT, value3}}}}, expression3)
}

func TestLte(t *testing.T) {
	field := "fieldToCheck"
	value1 := -20
	value2 := 46
	value3 := 23

	expression1 := query.Lte(field, value1)
	expression2 := query.Lte(field, value2)
	expression3 := query.Lte(field, value3)

	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.LTE, value1}}}}, expression1)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.LTE, value2}}}}, expression2)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.LTE, value3}}}}, expression3)
}

func TestNe(t *testing.T) {
	field := "fieldToCheck"
	value1 := "StringValue"
	value2 := 4627425234232
	value3 := true

	expression1 := query.Ne(field, value1)
	expression2 := query.Ne(field, value2)
	expression3 := query.Ne(field, value3)

	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.NE, value1}}}}, expression1)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.NE, value2}}}}, expression2)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.NE, value3}}}}, expression3)
}

func TestNin(t *testing.T) {
	field := "fieldToCheck"
	value1 := bson.A{"2", "23", "12"}
	value2 := bson.A{1, 2, -3}
	value3 := bson.A{"abcd", "Json", "Bson"}

	expression1 := query.Nin(field, value1)
	expression2 := query.Nin(field, value2)
	expression3 := query.Nin(field, value3)

	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.NIN, value1}}}}, expression1)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.NIN, value2}}}}, expression2)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.NIN, value3}}}}, expression3)
}

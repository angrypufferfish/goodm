package query_test

import (
	"testing"

	op_cmd "github.com/angrypufferfish/goodm/src/operators"
	"github.com/angrypufferfish/goodm/src/query"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestAnd(t *testing.T) {

	value1 := bson.D{{"$set", bson.D{{"field1", "value1"}}}}
	value2 := bson.D{{"$set", bson.D{{"field21", "value21"}}}}
	value3 := bson.D{{"$set", bson.D{{"field22", "value22"}}}}

	expression1 := query.And(value1)
	expression2 := query.And(value2, value3)

	assert.Equal(t, bson.D{{op_cmd.AND, bson.A{value1}}}, expression1)
	assert.Equal(t, bson.D{{op_cmd.AND, bson.A{value2, value3}}}, expression2)
}

func TestNot(t *testing.T) {

	field := "fieldToCheck"
	value1 := bson.D{{"$eq", bson.D{{"field22", "value22"}}}}
	value2 := bson.D{{"$ne", bson.D{{"field22", "value22"}}}}
	value3 := bson.D{{"$eq", bson.D{{"field22", "value22"}}}}

	expression1 := query.Not(field, value1)
	expression2 := query.Not(field, value2)
	expression3 := query.Not(field, value3)

	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.NOT, value1}}}}, expression1)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.NOT, value2}}}}, expression2)
	assert.Equal(t, bson.D{{field, bson.D{{op_cmd.NOT, value3}}}}, expression3)
}

func TestNor(t *testing.T) {

	value1 := bson.D{{"$ne", bson.D{{"field1", "value1"}}}}
	value2 := bson.D{{"$ne", bson.D{{"field21", "value21"}}}}
	value3 := bson.D{{"$eq", bson.D{{"field22", "value22"}}}}

	expression1 := query.Nor(value1)
	expression2 := query.Nor(value2, value3)

	assert.Equal(t, bson.D{{op_cmd.NOR, bson.A{value1}}}, expression1)
	assert.Equal(t, bson.D{{op_cmd.NOR, bson.A{value2, value3}}}, expression2)
}

func TestOr(t *testing.T) {

	value1 := bson.D{{"$eq", bson.D{{"field1", "value1"}}}}
	value2 := bson.D{{"$ne", bson.D{{"field21", "value21"}}}}
	value3 := bson.D{{"$eq", bson.D{{"field22", "value22"}}}}

	expression1 := query.Or(value1)
	expression2 := query.Or(value2, value3)

	assert.Equal(t, bson.D{{op_cmd.OR, bson.A{value1}}}, expression1)
	assert.Equal(t, bson.D{{op_cmd.OR, bson.A{value2, value3}}}, expression2)
}

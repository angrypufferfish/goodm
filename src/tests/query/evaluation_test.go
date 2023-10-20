package query_test

import (
	"testing"

	op_cmd "github.com/angrypufferfish/goodm/src/operators"
	"github.com/angrypufferfish/goodm/src/query"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestWhere(t *testing.T) {
	value1 := false
	value2 := "valueCheck"
	value3 := 23

	expression1 := query.Where(value1)
	expression2 := query.Where(value2)
	expression3 := query.Where(value3)

	assert.Equal(t, bson.D{{op_cmd.WHERE, value1}}, expression1)
	assert.Equal(t, bson.D{{op_cmd.WHERE, value2}}, expression2)
	assert.Equal(t, bson.D{{op_cmd.WHERE, value3}}, expression3)
}

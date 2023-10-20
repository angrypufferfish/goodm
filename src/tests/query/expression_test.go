package query_test

import (
	"testing"

	"github.com/angrypufferfish/goodm/src/query"
	"github.com/stretchr/testify/assert"
)

func TestNewPrimaryFieldExpression(t *testing.T) {

	operator := "$operator"
	field := "fieldToSet"
	value := "valueToSet"

	res := query.NewPrimaryFieldExpression(operator, field, value)

	assert.Equal(t, operator, *res.GetOperator())
	assert.Equal(t, field, *res.GetField())
	assert.Equal(t, value, *res.GetValue())

}

func TestNewPrimaryOperatorExpression(t *testing.T) {

	operator := "$operator"
	value := "valueToSet"

	res := query.NewPrimaryOperatorExpression(operator, value)

	assert.Equal(t, operator, *res.GetOperator())
	assert.Equal(t, value, *res.GetValue())
	assert.Nil(t, res.GetField())

}

package query

import (
	op_cmd "github.com/angrypufferfish/goodm/src/operators"
	"go.mongodb.org/mongo-driver/bson"
)

func And(expressions ...bson.D) bson.D {

	operations := mergeArrayList(expressions...)

	logic := NewPrimaryOperatorExpression(op_cmd.AND, operations)
	return logic.generatePrimaryOperator()
}

func Not(field string, expression bson.D) bson.D {

	logic := NewPrimaryFieldExpression(op_cmd.NOT, field, expression)
	return logic.generatePrimaryField()
}

func Nor(expressions ...bson.D) bson.D {

	operations := mergeArrayList(expressions...)

	logic := NewPrimaryOperatorExpression(op_cmd.NOR, operations)
	return logic.generatePrimaryOperator()
}

func Or(expressions ...bson.D) bson.D {

	operations := mergeArrayList(expressions...)

	logic := NewPrimaryOperatorExpression(op_cmd.OR, operations)
	return logic.generatePrimaryOperator()
}

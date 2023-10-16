package query

import (
	op_cmd "github.com/angrypufferfish/goodm/src/operators"
	"go.mongodb.org/mongo-driver/bson"
)

func Where(value any) bson.D {
	operation := NewPrimaryOperatorExpression(op_cmd.WHERE, value)
	return operation.generatePrimaryOperator()
}

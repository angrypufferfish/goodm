package query

import (
	op_cmd "github.com/angrypufferfish/goodm/src/operators"
	"go.mongodb.org/mongo-driver/bson"
)

func All(field string, value bool) bson.D {
	operation := NewPrimaryFieldExpression(op_cmd.ALL, field, value)
	return operation.generatePrimaryField()
}

func ElemMatch(field string, value any) bson.D {
	operation := NewPrimaryFieldExpression(op_cmd.ELEM_MATCH, field, value)
	return operation.generatePrimaryField()
}

func Size(field string, value bool) bson.D {
	operation := NewPrimaryFieldExpression(op_cmd.SIZE, field, value)
	return operation.generatePrimaryField()
}

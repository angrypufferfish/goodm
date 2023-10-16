package query

import (
	op_cmd "github.com/angrypufferfish/goodm/src/operators"
	"go.mongodb.org/mongo-driver/bson"
)

func Exists(field string, value bool) bson.D {
	operation := NewPrimaryFieldExpression(op_cmd.EXISTS, field, value)
	return operation.generatePrimaryField()
}

func Type(field string, value any) bson.D {
	operation := NewPrimaryFieldExpression(op_cmd.TYPE, field, value)
	return operation.generatePrimaryField()
}

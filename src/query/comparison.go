package query

import (
	op_cmd "github.com/angrypufferfish/goodm/src/operators"
	"go.mongodb.org/mongo-driver/bson"
)

func Eq(field string, value any) bson.D {
	operation := NewPrimaryFieldExpression(op_cmd.EQ, field, value)
	return operation.generatePrimaryField()
}

func Gt(field string, value any) bson.D {
	operation := NewPrimaryFieldExpression(op_cmd.GT, field, value)
	return operation.generatePrimaryField()
}

func Gte(field string, value any) bson.D {
	operation := NewPrimaryFieldExpression(op_cmd.GTE, field, value)
	return operation.generatePrimaryField()
}

func In(field string, value bson.A) bson.D {
	operation := NewPrimaryFieldExpression(op_cmd.IN, field, value)
	return operation.generatePrimaryField()
}

func Lt(field string, value any) bson.D {
	operation := NewPrimaryFieldExpression(op_cmd.LT, field, value)
	return operation.generatePrimaryField()
}

func Lte(field string, value any) bson.D {
	operation := NewPrimaryFieldExpression(op_cmd.LTE, field, value)
	return operation.generatePrimaryField()
}

func Ne(field string, value any) bson.D {
	operation := NewPrimaryFieldExpression(op_cmd.NE, field, value)
	return operation.generatePrimaryField()
}

func Nin(field string, value bson.A) bson.D {
	operation := NewPrimaryFieldExpression(op_cmd.NIN, field, value)
	return operation.generatePrimaryField()
}

package query

import (
	op_cmd "github.com/angrypufferfish/goodm/src/operators"
	"go.mongodb.org/mongo-driver/bson"
)

func Inc(expressions ...bson.E) bson.D {
	operation := NewPrimaryOperatorExpression(op_cmd.INC, expressions)
	return operation.generatePrimaryOperator()
}

func Min(expressions ...bson.E) bson.D {
	operation := NewPrimaryOperatorExpression(op_cmd.MIN, expressions)
	return operation.generatePrimaryOperator()
}

func Max(expressions ...bson.E) bson.D {
	operation := NewPrimaryOperatorExpression(op_cmd.MAX, expressions)
	return operation.generatePrimaryOperator()
}

func Mul(expressions ...bson.E) bson.D {
	operation := NewPrimaryOperatorExpression(op_cmd.MUL, expressions)
	return operation.generatePrimaryOperator()
}

func Rename(expressions ...bson.E) bson.D {
	operation := NewPrimaryOperatorExpression(op_cmd.RENAME, expressions)
	return operation.generatePrimaryOperator()
}

func Set(expressions any) bson.D {
	operation := NewPrimaryOperatorExpression(op_cmd.SET, expressions)
	return operation.generatePrimaryOperator()
}

func SetOnInsert(expressions ...bson.E) bson.D {
	operation := NewPrimaryOperatorExpression(op_cmd.SET_ON_INSERT, expressions)
	return operation.generatePrimaryOperator()
}

func Unset(expressions ...bson.E) bson.D {
	operation := NewPrimaryOperatorExpression(op_cmd.UNSET, expressions)
	return operation.generatePrimaryOperator()
}

func CurrentDate(expressions ...bson.E) bson.D {
	operation := NewPrimaryOperatorExpression(op_cmd.CURRENT_DATE, expressions)
	return operation.generatePrimaryOperator()
}

func AddToSet(expressions ...bson.E) bson.D {
	operation := NewPrimaryOperatorExpression(op_cmd.ADD_TO_SET, expressions)
	return operation.generatePrimaryOperator()
}

func Pop(expressions ...bson.E) bson.D {
	operation := NewPrimaryOperatorExpression(op_cmd.POP, expressions)
	return operation.generatePrimaryOperator()
}

func Pull(expressions ...bson.E) bson.D {
	operation := NewPrimaryOperatorExpression(op_cmd.PULL, expressions)
	return operation.generatePrimaryOperator()
}

func Push(expressions ...bson.E) bson.D {
	operation := NewPrimaryOperatorExpression(op_cmd.PUSH, expressions)
	return operation.generatePrimaryOperator()
}

func PullAll(expressions ...bson.E) bson.D {
	operation := NewPrimaryOperatorExpression(op_cmd.PULL_ALL, expressions)
	return operation.generatePrimaryOperator()
}

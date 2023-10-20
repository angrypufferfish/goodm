package query

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Expression struct {
	operator *string
	field    *string
	value    *any
}

func (exp *Expression) GetOperator() *string {
	return exp.operator
}

func (exp *Expression) GetField() *string {
	return exp.field
}

func (exp *Expression) GetValue() *any {
	return exp.value
}

func (exp *Expression) generatePrimaryField() bson.D {

	return bson.D{{
		*exp.field,
		bson.D{{
			*exp.operator,
			*exp.value,
		}},
	}}
}

func (exp *Expression) generatePrimaryOperator() bson.D {

	return bson.D{{
		*exp.operator,
		*exp.value,
	}}
}

func NewPrimaryFieldExpression(operator string, field string, value any) *Expression {

	return &Expression{
		operator: &operator,
		field:    &field,
		value:    &value,
	}
}

func NewPrimaryOperatorExpression(operator string, value any) *Expression {

	return &Expression{
		operator: &operator,
		value:    &value,
	}
}

func mergeOrderedList(expressions ...bson.D) bson.D {
	var updateExpressions bson.D

	for i := 0; i < len(expressions); i++ {
		expression := expressions[i]
		for t := 0; t < len(expression); t++ {
			element := expression[t]
			updateExpressions = append(updateExpressions, element)
		}
	}

	return updateExpressions
}

func mergeArrayList(expressions ...bson.D) bson.A {
	var updateExpressions bson.A

	for i := 0; i < len(expressions); i++ {
		expression := expressions[i]
		for t := 0; t < len(expression); t++ {
			element := expression[t]
			updateExpressions = append(updateExpressions, bson.D{element})
		}
	}

	return updateExpressions
}

func Update(expressions ...bson.D) bson.D {
	return mergeOrderedList(expressions...)
}

func Compare(expressions ...bson.D) bson.D {
	return mergeOrderedList(expressions...)
}

package op

import "go.mongodb.org/mongo-driver/bson"

var query = Operator{}

type Operator struct{}

type Commands = bson.D
type Array = bson.A

func (q *Operator) generateBsonArray(values []any) Array {
	bsonArray := Array{}
	for _, val := range values {
		bsonArray = append(bsonArray, val)
	}
	return bsonArray
}

func (q *Operator) OperatorValue(operator string, value any) Commands {
	return Commands{{
		operator,
		value,
	}}
}

func (q *Operator) OperatorCommand(operator string, cmds Commands) Commands {

	return Commands{{
		operator,
		cmds,
	}}
}

func (q *Operator) OperatorArray(operator string, values []any) Commands {
	bsonArray := q.generateBsonArray(values)

	return Commands{{
		operator,
		bsonArray,
	}}
}

func (q *Operator) FieldWithOperator(operator string, fieldName string, value any) Commands {
	return Commands{{
		fieldName,
		q.OperatorValue(operator, value),
	}}
}

func (q *Operator) FieldWithOperatorArray(operator string, fieldName string, values []any) Commands {

	return Commands{{
		fieldName,
		q.OperatorArray(operator, values),
	}}
}

func (q *Operator) FieldWithOperatorCommand(operator string, cmds ...Commands) Commands {

	///CAST TO []any
	array := make([]any, len(cmds))
	for i, e := range cmds {
		array[i] = e
	}

	return q.OperatorArray(operator, array)
}

func (q *Operator) FieldWithOperatorValue(operator string, fieldName string, cmds Commands) Commands {

	return Commands{{
		fieldName,
		q.OperatorValue(operator, cmds),
	}}
}

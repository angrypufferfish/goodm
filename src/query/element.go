package query

import (
	op_cmd "github.com/angrypufferfish/goodm/src/operators"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

func Exists(field string, value bool) bson.D {
	operation := NewPrimaryFieldExpression(op_cmd.EXISTS, field, value)
	return operation.generatePrimaryField()
}

func Type(field string, value bsontype.Type) bson.D {
	operation := NewPrimaryFieldExpression(op_cmd.TYPE, field, value)
	return operation.generatePrimaryField()
}

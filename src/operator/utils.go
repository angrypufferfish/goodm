package op

import (
	"github.com/angrypufferfish/goodm/src/exceptions"
	"go.mongodb.org/mongo-driver/bson"
)

func checkCommandsLen[Cmd any](cmds ...Cmd) {
	if len(cmds) < 2 {
		panic(exceptions.ErrOperatorNotEnoughParams)
	}
}

func checkArrayValueLen[Value any](cmds ...Value) {
	if len(cmds) < 1 {
		panic(exceptions.ErrOperatorNotEnoughParams)
	}
}

func generateBsonArray[E any](values []E) bson.A {
	bsonArray := bson.A{}
	for _, val := range values {
		bsonArray = append(bsonArray, val)
	}
	return bsonArray
}

func generateBsonArrayFromSplit[E any](values ...E) bson.A {
	bsonArray := bson.A{}
	for _, val := range values {
		bsonArray = append(bsonArray, val)
	}
	return bsonArray
}

func operatorCommands(operator string, cmd Commands) Commands {

	return Commands{{
		operator,
		cmd,
	}}
}

func operatorValue(operator string, value any) Commands {
	return Commands{{
		operator,
		value,
	}}
}

func operatorFieldValue(operator string, fieldName string, value any) Commands {
	return Commands{{
		fieldName,
		Commands{{
			operator,
			value,
		}},
	}}
}

func updateOperatorFieldValues(operator string, cmds Commands) Commands {
	return Commands{{
		operator,
		cmds,
	}}
}

func operatorArrayValue(operator string, fieldName string, values []any) Commands {

	bsonArray := generateBsonArray[any](values)

	return Commands{{
		fieldName,
		Commands{{
			operator,
			bsonArray,
		}},
	}}
}

func operatorArrayCommands(operator string, cmds ...Commands) Commands {
	checkCommandsLen[Commands](cmds...)

	commandsArray := generateBsonArray[Commands](cmds)

	return Commands{{
		operator,
		commandsArray,
	}}
}

func operatorArrayCommandsValue(operator string, fieldName string, cmds Commands) Commands {

	return Commands{{
		fieldName,
		Commands{{
			operator,
			cmds,
		}},
	}}
}

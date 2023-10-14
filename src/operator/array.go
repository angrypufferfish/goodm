package op

import op_cmd "github.com/angrypufferfish/goodm/src/operator_cmd"

func All(field string, values []any) Commands {
	return operatorArrayValue(op_cmd.ALL, field, values)
}

func ElemMatch(field string, cmds Commands) Commands {
	return operatorArrayCommandsValue(op_cmd.ELEM_MATCH, field, cmds)
}

func Size(field string, value any) Commands {
	return operatorFieldValue(op_cmd.SIZE, field, value)
}

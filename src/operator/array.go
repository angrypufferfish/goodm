package op

import op_cmd "github.com/angrypufferfish/goodm/src/operator_cmd"

func All(field string, values []any) Commands {
	return query.FieldWithOperatorArray(op_cmd.ALL, field, values)
}

func ElemMatch(field string, cmds Commands) Commands {
	return query.FieldWithOperatorValue(op_cmd.ELEM_MATCH, field, cmds)
}

func Size(field string, value any) Commands {
	return query.FieldWithOperator(op_cmd.SIZE, field, value)
}

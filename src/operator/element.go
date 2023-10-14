package op

import op_cmd "github.com/angrypufferfish/goodm/src/operator_cmd"

func Exists(field string, value bool) Commands {
	return query.FieldWithOperator(op_cmd.EXISTS, field, value)
}

func Type(field string, value any) Commands {
	return query.FieldWithOperator(op_cmd.TYPE, field, value)
}

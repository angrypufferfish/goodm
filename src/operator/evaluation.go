package op

import op_cmd "github.com/angrypufferfish/goodm/src/operator_cmd"

func Expr(cmd Commands) Commands {
	return query.OperatorCommand(op_cmd.EXPR, cmd)
}

func JsonSchema(cmd Commands) Commands {
	return query.OperatorCommand(op_cmd.JSONSCHEMA, cmd)
}

func Mod(field string, values []any) Commands {
	return query.FieldWithOperatorArray(op_cmd.MOD, field, values)
}

//TODO: Regex function
//TODO: Text function

func Where(value any) Commands {
	return query.OperatorValue(op_cmd.WHERE, value)
}

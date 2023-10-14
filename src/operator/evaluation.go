package op

import op_cmd "github.com/angrypufferfish/goodm/src/operator_cmd"

func Expr(cmd Commands) Commands {
	return operatorCommands(op_cmd.EXPR, cmd)
}

func JsonSchema(cmd Commands) Commands {
	return operatorCommands(op_cmd.JSONSCHEMA, cmd)
}

func Mod(field string, values []any) Commands {
	return operatorArrayValue(op_cmd.MOD, field, values)
}

//TODO: Regex function
//TODO: Text function

func Where(value any) Commands {
	return operatorValue(op_cmd.WHERE, value)
}

package op

import op_cmd "github.com/angrypufferfish/goodm/src/operator_cmd"

func And(cmds ...Commands) Commands {
	return query.FieldWithOperatorCommand(op_cmd.AND, cmds...)
}

func Not(cmds ...Commands) Commands {
	return query.FieldWithOperatorCommand(op_cmd.NOT, cmds...)
}

func Nor(cmds ...Commands) Commands {
	return query.FieldWithOperatorCommand(op_cmd.NOR, cmds...)
}

func Or(cmds ...Commands) Commands {
	return query.FieldWithOperatorCommand(op_cmd.OR, cmds...)
}

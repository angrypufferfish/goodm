package op

import op_cmd "github.com/angrypufferfish/goodm/src/operator_cmd"

func Set(cmds Commands) Commands {
	return query.OperatorCommand(op_cmd.SET, cmds)
}

func CurrentDate(cmds Commands) Commands {
	return query.OperatorCommand(op_cmd.CURRENT_DATE, cmds)
}

func Inc(cmds Commands) Commands {
	return query.OperatorCommand(op_cmd.INC, cmds)
}

func Min(cmds Commands) Commands {
	return query.OperatorCommand(op_cmd.MIN, cmds)
}

func Max(cmds Commands) Commands {
	return query.OperatorCommand(op_cmd.MAX, cmds)
}

func Mul(cmds Commands) Commands {
	return query.OperatorCommand(op_cmd.MUL, cmds)
}

func Rename(cmds Commands) Commands {
	return query.OperatorCommand(op_cmd.RENAME, cmds)
}

func SetOnInsert(cmds Commands) Commands {
	return query.OperatorCommand(op_cmd.SET_ON_INSERT, cmds)
}

func Unset(cmds Commands) Commands {
	return query.OperatorCommand(op_cmd.UNSET, cmds)
}

/**ARRAY**/

func AddToSet(cmds Commands) Commands {
	return query.OperatorCommand(op_cmd.ADD_TO_SET, cmds)
}

func Pop(cmds Commands) Commands {
	return query.OperatorCommand(op_cmd.POP, cmds)
}

func Pull(cmds Commands) Commands {
	return query.OperatorCommand(op_cmd.PULL, cmds)
}

func Push(cmds Commands) Commands {
	return query.OperatorCommand(op_cmd.PUSH, cmds)
}

func PullAll(cmds Commands) Commands {
	return query.OperatorCommand(op_cmd.PULL_ALL, cmds)
}

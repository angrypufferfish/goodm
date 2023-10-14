package op

import (
	op_cmd "github.com/angrypufferfish/goodm/src/operator_cmd"
)

// Matches values that are equal to a specified value.
func Eq(field string, value any) Commands {
	return operatorFieldValue(op_cmd.EQ, field, value)
}

// Matches values that are greater than a specified value.
func Gt(field string, value any) Commands {
	return operatorFieldValue(op_cmd.GT, field, value)
}

// Matches values that are greater than or equal to a specified value.
func Gte(field string, value any) Commands {
	return operatorFieldValue(op_cmd.GTE, field, value)
}

// Matches values that are less than a specified value.
func Lt(field string, value any) Commands {
	return operatorFieldValue(op_cmd.LT, field, value)
}

// Matches values that are less than or equal to a specified value.
func Lte(field string, value any) Commands {
	return operatorFieldValue(op_cmd.LTE, field, value)
}

// Matches all values that are not equal to a specified value.
func Ne(field string, value any) Commands {
	return operatorFieldValue(op_cmd.NE, field, value)
}

// Matches any of the values specified in an array.
func In(field string, values []any) Commands {
	return operatorArrayValue(op_cmd.IN, field, values)
}

// Matches none of the values specified in an array.
func NotIn(field string, values []any) Commands {
	return operatorArrayValue(op_cmd.NIN, field, values)
}

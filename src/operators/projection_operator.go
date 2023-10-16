package op_cmd

// PROJECTION
const (
	/*
	* Projects the first element in an array that matches the query condition.
	 */
	FIRST_ELEM string = "$"
	META       string = "$meta"
	SLICE      string = "$slice"
)

// MISCELLANEOUS
const (
	COMMENT string = "$comment"
	RAND    string = "$rand"
	NATURAL string = "$natural"
)

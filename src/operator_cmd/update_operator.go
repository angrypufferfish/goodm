package op_cmd

// FIELD
const (
	CURRENT_DATE  string = "$currentDate"
	INC           string = "$inc"
	MIN           string = "$min"
	MAX           string = "$max"
	MUL           string = "$mul"
	RENAME        string = "$rename"
	SET           string = "$set"
	SET_ON_INSERT string = "$setOnInsert"
	UNSET         string = "$unset"
)

// ARRAY
const (
	ADD_TO_SET string = "$addToSet"
	POP        string = "$pop"
	PULL       string = "$pull"
	PUSH       string = "$push"
	PULL_ALL   string = "$pullAll"

	POSITION string = "$position"
	SORT     string = "$sort"
	EACH     string = "$each"
	//SLICE    string = "$slice"
)

// BITWISE
const (
	BIT string = "$bit"
)

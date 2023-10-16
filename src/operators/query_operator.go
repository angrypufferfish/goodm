package op_cmd

// COMPARISON
const (
	EQ  string = "$eq"
	GT  string = "$gt"
	GTE string = "$gte"
	IN  string = "$in"
	LT  string = "$lt"
	LTE string = "$lte"
	NE  string = "$ne"
	NIN string = "$nin"
)

// LOGICAL
const (
	AND string = "$and"
	NOT string = "$not"
	NOR string = "$nor"
	OR  string = "$or"
)

// ELEMENT
const (
	EXISTS string = "$exists"
	TYPE   string = "$type"
)

// EVALUATION
const (
	EXPR       string = "$expr"
	JSONSCHEMA string = "$jsonschema"
	MOD        string = "$mod"
	REGEX      string = "$regex"
	TEXT       string = "$text"
	WHERE      string = "$where"
)

// GEOSPATIAL
const (
	GEO_INTERSECTS string = "$geoIntersects"
	GEO_WITHIN     string = "$geoWithin"
	NEAR           string = "$near"
	NEAR_SPHERE    string = "$nearSphere"

	BOX           string = "$box"
	CENTER        string = "$center"
	CENTER_SPHERE string = "$centerSphere"
	GEOMETRY      string = "$geometry"
	MAX_DISTANCE  string = "$maxDistance"
	MIN_DISTANCE  string = "$minDistance"
	POLYGON       string = "$polygon"
)

// ARRAY
const (
	ALL        string = "$all"
	ELEM_MATCH string = "$elemMatch"
	SIZE       string = "$size"
)

// BITWISE
const (
	BITS_ALL_CLEAR string = "$bitsAllClear"
	BITS_ALL_SET   string = "$bitsAllSet"
	BITS_ANY_CLEAR string = "$bitsAnyClear"
	BITS_ANY_SET   string = "$bitsAnySet"
)

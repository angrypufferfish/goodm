package exceptions

import "errors"

var (
	ErrDbNotInitialized         = errors.New("cannot parse NaN as a *big.Int")
	ErrCollectionNameNotDefined = errors.New("goodmCollection field have to be defined on Document definition")
)

package exceptions

import "errors"

var (
	ErrDbNotInitialized         = errors.New("Database not initialized")
	ErrCollectionNameNotDefined = errors.New("goodm tag is mandatory on Document definition")
)

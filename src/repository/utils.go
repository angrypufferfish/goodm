package repository

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func objectIdConvert(id any) (*primitive.ObjectID, error) {
	switch id.(type) {
	case string:
		objectID, err := primitive.ObjectIDFromHex(id.(string))
		if err != nil {
			panic(err)
		}
		return &objectID, nil
	case primitive.ObjectID:
		objectID := id.(primitive.ObjectID)
		return &objectID, nil
	default:
		return nil, fmt.Errorf("Id must be a string | primitive.ObjectID")
	}
}

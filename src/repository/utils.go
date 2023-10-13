package repository

import (
	"fmt"

	"github.com/angrypufferfish/goodm/src/serializer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ObjectIdConvert(id any) (*primitive.ObjectID, error) {
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

func serializeSingleResult[S any](singleResult *mongo.SingleResult) (S, error) {
	var result bson.M
	var seralizedDocument S

	err := singleResult.Decode(&result)

	if err != nil {
		return seralizedDocument, err
	}

	seralizedDocument, err = serializer.Serialize[S](result)

	if err != nil {
		return seralizedDocument, err
	}

	return seralizedDocument, nil
}

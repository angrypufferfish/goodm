package serializer

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func Serialize[A any](doc any) (A, error) {

	var data A

	bsonBytes, err := bson.Marshal(doc)

	if err != nil {
		return data, fmt.Errorf("BSON Marshal error - %s", err)
	}

	err = bson.Unmarshal(bsonBytes, &data)
	if err != nil {
		return data, fmt.Errorf("BSON Unmarshal error - %s", err)
	}

	return data, nil
}

func SerializeList[A any](docs []bson.D) ([]A, error) {
	var data []A

	for _, doc := range docs {
		serializedDoc, err := Serialize[A](doc)

		if err != nil {
			return data, err
		}
		data = append(data, serializedDoc)
	}

	return data, nil
}

package serializer

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func Serialize[A interface{}](doc any) A {
	var data A
	bsonBytes, err := bson.Marshal(doc)
	if err != nil {
		fmt.Printf("BSON Marshal error - %s", err)
	}

	err = bson.Unmarshal(bsonBytes, &data)
	if err != nil {
		fmt.Printf("BSON Unmarshal error - %s", err)
	}

	return data

}

func SerializeList[A interface{}](docs []bson.M) []A {
	var data []A

	for _, doc := range docs {
		data = append(data, Serialize[A](doc))
	}

	return data
}

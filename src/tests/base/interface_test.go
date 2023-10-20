package base_test

import (
	"testing"
	"time"

	"github.com/angrypufferfish/goodm/src/base"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestBaseDocumentImplementsInterfaceIBaseDocument(t *testing.T) {
	id := primitive.NewObjectID()
	currentDate := time.Now().UTC()

	doc := &base.BaseDocument{
		Id:        id,
		CreatedAt: currentDate,
		UpdatedAt: currentDate,
	}

	assert.Implements(t, (*base.IBaseDocument)(nil), doc)
}

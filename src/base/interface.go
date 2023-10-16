package base

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IBaseDocument interface {
	GetID() primitive.ObjectID
	GetUpdatedAt() time.Time
	GetCreatedAt() time.Time
	OnCreate()
	OnUpdate()
}

package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Command struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	DocumentId primitive.ObjectID `bson:"documentId" json:"documentId"`
	ItemId     primitive.ObjectID `bson:"itemId" json:"itemId"`
	Quantity   int                `bson:"quantity" json:"quantity"`
	Discount   int                `bson:"discount" json:"discount"`
}

type CreateCommandRequest struct {
	DocumentId primitive.ObjectID `bson:"documentId" json:"documentId"`
	ItemId     primitive.ObjectID `bson:"itemId" json:"itemId"`
	Quantity   int                `bson:"quantity" json:"quantity"`
	Discount   int                `bson:"discount" json:"discount"`
}

package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Documents struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Date     primitive.DateTime `bson:"date" json:"date"`
	ClientId primitive.ObjectID `bson:"clientId" json:"clientId"`
	UserId   primitive.ObjectID `bson:"userId" json:"userId"`
	Type     string             `bson:"type" json:"type"`
}

type CreateDocumentsRequest struct {
	ClientId primitive.ObjectID `bson:"clientId" json:"clientId"`
	UserId   primitive.ObjectID `bson:"userId" json:"userId"`
	Type     string             `bson:"type" json:"type"`
}

package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Client struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Firstname string             `bson:"name" json:"name"`
	Lastname  string             `bson:"lastname" json:"lastname"`
	Email     string             `bson:"email" json:"email"`
}

type CreateClientRequest struct {
	Firstname string `bson:"name" json:"firstname"`
	Lastname  string `bson:"lastname" json:"lastname"`
	Email     string `bson:"email" json:"email"`
}

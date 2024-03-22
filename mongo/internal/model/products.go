package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Category string             `bson:"category" json:"category"`
	Price    float32            `bson:"price" json:"price"`
	Stock    int                `bson:"stock" json:"stock"`
}

type CreateProductRequest struct {
	Name     string  `bson:"name" json:"name"`
	Category string  `bson:"category" json:"category"`
	Price    float32 `bson:"price" json:"price"`
	Stock    int     `bson:"stock" json:"stock"`
}

package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client    *mongo.Client
	Products  *mongo.Collection
	Clients   *mongo.Collection
	Commands  *mongo.Collection
	Documents *mongo.Collection
)

func Init(uri string, database string) error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	localClient, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return err
	}
	client = localClient

	Products = client.Database(database).Collection("products")
	Clients = client.Database(database).Collection("clients")
	Commands = client.Database(database).Collection("commands")
	Documents = client.Database(database).Collection("documents")

	err = client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err()
	return err
}
func Close() error {
	return client.Disconnect(context.Background())
}

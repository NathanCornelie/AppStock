package main

import (
	"Mongo/internal"
	"Mongo/internal/database"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"time"
)

var MongoUri = os.Getenv("MONGO_URI")

func main() {
	if MongoUri == "" {
		fmt.Println("MONGO_URI env var is not set")
		return
	}
	fmt.Println(primitive.NewDateTimeFromTime(time.Now()))
	err := database.Init(MongoUri, "developement")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected to MongoDB")
	defer func() {
		err := database.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	r := internal.DefineRouter()

	err = r.Run(":8080")
	if err != nil {
		return
	}

}

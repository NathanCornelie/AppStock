package main

import (
	"Mongo/internal"
	"Mongo/internal/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

var MongoUri = os.Getenv("MONGO_URI")

func main() {
	if MongoUri == "" {
		fmt.Println("MONGO_URI env var is not set")
		return
	}
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

	r := gin.Default()
	internal.AddProductRoutes(r)
	internal.AddClientsRoutes(r)
	r.Run(":8080")

}

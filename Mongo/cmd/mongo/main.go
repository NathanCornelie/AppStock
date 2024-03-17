package main

import (
	"Mongo/internal/database"
	"fmt"
)

var MongoUri = "mongodb+srv://nathan:AiV!M8aTrRqUEggQD_Hp@appstockdb.bxp0vgt.mongodb.net/?retryWrites=true&w=majority&appName=AppStockDB"

func main() {
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
			return
		}
	}()
}

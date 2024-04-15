package main

import (
	"Mongo/internal"
	"Mongo/internal/database"
	"fmt"
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
		fmt.Println("Erreur d'initiaisation de la DB : \n  \t- Vérifier la connexion internet \n \t- Vérifier les creds\n \t- l'IP n'est ptre pas dans la whitelist")
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

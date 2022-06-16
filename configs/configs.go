package configs

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB()  {
	// create a client
	client,err := mongo.NewClient(options.Client())

	if err != nil {
		log.Fatal(err)
	}
	
}
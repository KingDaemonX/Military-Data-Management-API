package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() {
	// create a client
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoURI()))

	if err != nil {
		log.Fatal(err)
	}

	cbg,cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	// connect client

}

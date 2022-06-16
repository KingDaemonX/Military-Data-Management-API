package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	// create a client
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoURI()))

	if err != nil {
		log.Fatal(err)
	}

	cbg, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	// connect client
	err = client.Connect(cbg)
	if err != nil {
		log.Fatal(err)
	}

	// ping database
	err = client.Ping(cbg, nil)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Database is not connected")
	}

	fmt.Println("Database Connected :) ...")
	return client

}


var clientt *mongo.Client = ConnectDB()

func DbCollection() *mongo.Collection {
	collections := clientt.Database("").Collection("")
	return collections
}
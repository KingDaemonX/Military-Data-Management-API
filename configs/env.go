package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
}

func MongoURI() string {
	loadEnv()
	return os.Getenv("MONGOURI")
}

func SecretKey() string {
	loadEnv()
	return os.Getenv("SECRET_KEY")
}

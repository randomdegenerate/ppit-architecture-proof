package mongodb

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)


const db = "warehouse"
const collName = "items"

var mongoClient * mongo.Client

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	connectionURI := os.Getenv("API_URI")
	if connectionURI == "" {
		fmt.Print("uri path is incorrect")
	}

	clientOptions := options.Client().ApplyURI(connectionURI)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		panic(err)
	}

	mongoClient = client
}


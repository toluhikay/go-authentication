package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// create a fuction of data base instance that returns a mongo client
func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	MongoDB := os.Getenv("MONGODB_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connection to database successful")

	return client
}

var Client *mongo.Client = DBinstance()

// create a function to access a particular collection in the dtabase
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("").Collection(collectionName)
	return collection
}

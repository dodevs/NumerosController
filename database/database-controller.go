package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DBInstance func
func DBInstance() *mongo.Client {
	MongoDbHost := os.Getenv("MONGODB_HOST")
	MongoDbPort := os.Getenv(("MONGODB_PORT"))

	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", MongoDbHost, MongoDbPort)))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected to MongoDB!\n")

	return client
}

//Client Database Instance
var Client *mongo.Client = DBInstance()

//OpenCollection is a function makes a connection with a collection in the database
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("cluster0").Collection(collectionName)
}

package main

import (
	"context"
	"fmt"
	"log"
	"test-task-auth-service-api/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// set client options
	clientOptions := options.Client().ApplyURI(config.Envs.MongoUri)

	// connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// check connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}
	log.Println("Connected to MongoDB")

	// disconnect from MongoDB
	err = client.Disconnect(context.Background())
	if err != nil {
		log.Fatal("Error disconnecting from MongoDB:", err)
	}
	fmt.Println("Disconnected from MongoDB")
}

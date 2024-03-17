package mongo_db

import (
	"context"
	"log"
	"test-task-auth-service-api/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() *mongo.Database {
	clientOptions := options.Client().ApplyURI(config.Envs.MongoUri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}

	return client.Database(config.Envs.MongoDatabaseName)
}

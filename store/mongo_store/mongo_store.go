package mongo_store

import "go.mongodb.org/mongo-driver/mongo"

type MongoStore struct {
	db *mongo.Database
}

func New(db *mongo.Database) *MongoStore {
	return &MongoStore{db}
}

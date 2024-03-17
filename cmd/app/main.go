package main

import (
	"test-task-auth-service-api/internal/config"
	"test-task-auth-service-api/internal/server"
	mongoStore "test-task-auth-service-api/internal/store/mongo_store"
	mongoDb "test-task-auth-service-api/pkg/db/mongo_db"
)

func main() {
	db := mongoDb.Init(config.Envs.MongoUri, config.Envs.MongoDatabaseName)
	store := mongoStore.New(db)

	server := server.New(config.Envs.Address, store)
	server.Run()
}

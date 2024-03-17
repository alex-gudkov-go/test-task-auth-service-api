package main

import (
	"test-task-auth-service-api/config"
	mongoDb "test-task-auth-service-api/db/mongo_db"
	"test-task-auth-service-api/server"
	mongoStore "test-task-auth-service-api/store/mongo_store"
)

func main() {
	db := mongoDb.Init()
	store := mongoStore.New(db)
	server := server.New(config.Envs.Address, store)
	server.Start()
}

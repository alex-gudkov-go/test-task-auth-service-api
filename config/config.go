package config

import (
	"log"
	"os"

	dotEnv "github.com/joho/godotenv"
)

type Config struct {
	Address           string
	MongoUri          string
	MongoDatabaseName string
}

var Envs = initEnvs()

func initEnvs() *Config {
	if err := dotEnv.Load(".env"); err != nil {
		log.Fatalln("Error loading .ENV file")
	}

	config := &Config{
		Address:           getEnvString("ADDRESS"),
		MongoUri:          getEnvString("MONGO_URI"),
		MongoDatabaseName: getEnvString("MONGO_DATABASE_NAME"),
	}

	log.Printf("Config: %+v", config)

	return config
}

func getEnvString(key string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		log.Fatalf("Env \"%s\" is not set\n", key)
	}

	return value
}

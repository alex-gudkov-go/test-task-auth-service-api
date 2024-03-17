package config

import (
	"log"
	"os"

	dotEnv "github.com/joho/godotenv"
)

type Config struct {
	MongoUri string
}

var Envs = initEnvs()

func initEnvs() *Config {
	if err := dotEnv.Load(".env"); err != nil {
		log.Fatalln("Error loading .ENV file")
	}

	config := &Config{
		MongoUri: getEnvString("MONGO_URI"),
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

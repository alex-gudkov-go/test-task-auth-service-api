package config

import (
	"log"
	"os"
	"strconv"

	dotEnv "github.com/joho/godotenv"
)

type Config struct {
	Address                      string
	MongoUri                     string
	MongoDatabaseName            string
	AccessTokenLifetimeInMinutes int
	AccessTokenSecret            string
	RefreshTokenSecret           string
}

var Envs = initEnvs()

func initEnvs() *Config {
	if err := dotEnv.Load(".env"); err != nil {
		log.Fatalln("Error loading .ENV file")
	}

	config := &Config{
		Address:                      getEnvString("ADDRESS"),
		MongoUri:                     getEnvString("MONGO_URI"),
		MongoDatabaseName:            getEnvString("MONGO_DATABASE_NAME"),
		AccessTokenLifetimeInMinutes: getEnvInt("ACCESS_TOKEN_LIFETIME_IN_MINUTES"),
		AccessTokenSecret:            getEnvString("ACCESS_TOKEN_SECRET"),
		RefreshTokenSecret:           getEnvString("REFRESH_TOKEN_SECRET"),
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

func getEnvInt(key string) int {
	value, ok := os.LookupEnv(key)

	if !ok {
		log.Fatalf("Env \"%s\" is not set", key)
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Env \"%s\" is not integer", key)
	}

	return intValue
}

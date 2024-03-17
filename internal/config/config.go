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
}

var Envs = initEnvs()

const envFileName = ".env"

func initEnvs() *Config {
	if err := dotEnv.Load(envFileName); err != nil {
		log.Fatalln("error loading env file")
	}

	config := &Config{
		Address:                      getEnvString("ADDRESS"),
		MongoUri:                     getEnvString("MONGO_URI"),
		MongoDatabaseName:            getEnvString("MONGO_DATABASE_NAME"),
		AccessTokenLifetimeInMinutes: getEnvInt("ACCESS_TOKEN_LIFETIME_IN_MINUTES"),
		AccessTokenSecret:            getEnvString("ACCESS_TOKEN_SECRET"),
	}

	log.Printf("Config: %+v", config)

	return config
}

func getEnvString(key string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		log.Fatalf("env variable \"%s\" is not set\n", key)
	}

	return value
}

func getEnvInt(key string) int {
	value, ok := os.LookupEnv(key)

	if !ok {
		log.Fatalf("env variable \"%s\" is not set\n", key)
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("env variable \"%s\" is not integer\n", key)
	}

	return intValue
}

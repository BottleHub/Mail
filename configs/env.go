package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Helps retrieve the MongoDB URI fron the env file
func EnvMongoURI() string {
	env, err := os.LookupEnv("MONGOURI")
	if !err {
		err := godotenv.Load()
		if err != nil {
			err := godotenv.Load(".env.test")
			if err != nil {
				log.Fatal("Error loading .env file: ", err)
			}
		}
		return os.Getenv("MONGOURI")
	}
	return env
}

// Helps retrieve the API Key fron the env file
func EnvApiKey() string {
	env, err := os.LookupEnv("API_KEY")
	if !err {
		err := godotenv.Load()
		if err != nil {
			err := godotenv.Load(".env.test")
			if err != nil {
				log.Fatal("Error loading .env file: ", err)
			}
		}
		return os.Getenv("API_KEY")
	}
	return env
}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentVariables struct {
	DATABASE_SSL_MODE string
	DATABASE_PASSWORD string
	DATABASE_USER     string
	DATABASE_NAME     string
	DATABASE_HOST     string
	DATABASE_PORT     string
	DATABASE_DIALECT  string
}

func Env() *EnvironmentVariables {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	env := EnvironmentVariables{
		DATABASE_SSL_MODE: os.Getenv("DATABASE_SSL_MODE"),
		DATABASE_PASSWORD: os.Getenv("DATABASE_PASSWORD"),
		DATABASE_USER:     os.Getenv("DATABASE_USER"),
		DATABASE_NAME:     os.Getenv("DATABASE_NAME"),
		DATABASE_HOST:     os.Getenv("DATABASE_HOST"),
		DATABASE_PORT:     os.Getenv("DATABASE_PORT"),
		DATABASE_DIALECT:  os.Getenv("DATABASE_DIALECT"),
	}
	return &env
}

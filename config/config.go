package config

import (
	"os"
	"strconv"
)

type Config struct {
	PORT string

	// databse variables
	POSTGRES_HOST     string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_NAME     string
	POSTGRES_PORT     int
}

var Envs = initConfig()

func initConfig() Config {
	// load environment variables
	// err := godotenv.Load("./../.env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	return Config{
		PORT: getEnvVariable("PORT", ":8080"),

		POSTGRES_HOST:     getEnvVariable("POSTGRES_HOST", "db_host"),
		POSTGRES_USER:     getEnvVariable("POSTGRES_USER", "db_user"),
		POSTGRES_PASSWORD: getEnvVariable("POSTGRES_PASSWORD", "db_password"),
		POSTGRES_NAME:     getEnvVariable("POSTGRES_NAME", "db_name"),
		POSTGRES_PORT:     getEnvVariableAsInt("POSTGRES_PORT", 5432),
	}
}

func getEnvVariable(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return fallback
}

func getEnvVariableAsInt(key string, fallback int) int {
	val := getEnvVariable(key, strconv.Itoa(fallback))
	valInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return valInt
}

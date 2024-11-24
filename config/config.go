package config

import "github.com/kacpertarka/restaurant/utils"

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
		PORT: utils.GetEnvVariable("PORT", ":8080"),

		POSTGRES_HOST:     utils.GetEnvVariable("POSTGRES_HOST", "db_host"),
		POSTGRES_USER:     utils.GetEnvVariable("POSTGRES_USER", "db_user"),
		POSTGRES_PASSWORD: utils.GetEnvVariable("POSTGRES_PASSWORD", "db_password"),
		POSTGRES_NAME:     utils.GetEnvVariable("POSTGRES_NAME", "db_name"),
		POSTGRES_PORT:     utils.GetEnvVariableAsInt("POSTGRES_PORT", 5432),
	}
}

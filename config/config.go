package config

import "os"

type Config struct {
	PORT string
}

func InitConfig()  Config{
	return Config{
		PORT: getEnvVariable("PORT", ":8080"),
	}
}


func getEnvVariable(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return fallback
}
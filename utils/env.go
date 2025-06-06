package utils

import (
	"log"

	"github.com/joho/godotenv"
)

type Environment struct {
	AccessKey string
	SecretKey string
}

func SetEnvironment(AccessKey string, SecretKey string) {
	env_variables := map[string]string{
		"AWS_ACCESS_KEY_ID":     AccessKey,
		"AWS_SECRET_ACCESS_KEY": SecretKey,
	}
	err := godotenv.Write(env_variables, ".env")
	if err != nil {
		log.Fatal(err)
	}
}

func ClearEnvironment() {
	env_variables := map[string]string{
		"AWS_ACCESS_KEY_ID":     "",
		"AWS_SECRET_ACCESS_KEY": "",
	}
	err := godotenv.Write(env_variables, ".env")
	if err != nil {
		log.Fatal(err)
	}
}

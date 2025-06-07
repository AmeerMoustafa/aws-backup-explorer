package utils

import (
	"embed"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type EnvVariables struct {
	AccessKey string `mapstructure:"AWS_ACCESS_KEY_ID"`
	SecretKey string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
}

var EnvironmentVariables EnvVariables

func (EnvironmentVariables *EnvVariables) LoadEnvVariables(envFile embed.FS) {
	file, err := envFile.Open(".env")
	if err != nil {
		log.Fatalln("Error: Could not open environment file")
	}

	defer file.Close()

	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadConfig(file)
	if err != nil {
		log.Fatalln(err)

	}

	err = viper.Unmarshal(&EnvironmentVariables)
	if err != nil {
		log.Fatalln(err)

	}

}

func SetEnvironment(AccessKey string, SecretKey string, EnvironmentVariables *EnvVariables) {
	viper.Set("AWS_ACCESS_KEY_ID", AccessKey)
	viper.Set("AWS_SECRET_ACCESS_KEY", SecretKey)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println(err)
	}

}

func ClearEnvironment() {
	viper.Set("AWS_ACCESS_KEY_ID", "")
	viper.Set("AWS_SECRET_ACCESS_KEY", "")
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println(err)
	}
}

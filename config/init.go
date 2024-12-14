package config

import (
	"github.com/spf13/viper"
	"log"
)

type EnvType struct {
	AppPort         string `mapstructure:"APP_PORT"`
	MongoUrl        string `mapstructure:"MONGO_URL"`
	DbName          string `mapstructure:"MONGO_DB_NAME"`
	SecretKey       string `mapstructure:"SECRET_KEY"`
	SwaggerUser     string `mapstructure:"SWAGGER_USER"`
	SwaggerPassword string `mapstructure:"SWAGGER_PASSWORD"`
}

var Env *EnvType

func Init() {
	if Env != nil {
		return
	}

	env := EnvType{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	Env = &env
}

package config

import (
	"log"

	"github.com/spf13/viper"
)

type env struct {
	JWTSecret string `mapstructure:"JWT_SECRET"`
	Port      string `mapstructure:"PORT"`
}

func ConfigEnv() *env {
	env := env{}
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}

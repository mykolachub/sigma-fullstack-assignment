package config

import (
	"log"

	"github.com/spf13/viper"
)

type env struct {
	JWTSecret    string `mapstructure:"JWT_SECRET"`
	Port         string `mapstructure:"PORT"`
	DBUser       string `mapstructure:"DB_USER"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	DBName       string `mapstructure:"DB_NAME"`
	DBSSLMode    string `mapstructure:"DB_SSLMODE"`
	DBConnString string `mapstructure:"DB_CONN_STRING"`
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

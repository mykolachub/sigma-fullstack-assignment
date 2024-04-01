package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	JWTSecret  string
	Port       string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	Test       string
}

func ConfigEnv() *Env {
	var env Env
	err := envconfig.Process("SIGMA", &env)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &env
}

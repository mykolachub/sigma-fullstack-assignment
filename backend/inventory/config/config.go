package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	HttpPort                string `envconfig:"HTTP_PORT"`
	GrpcPort                string `envconfig:"GRPC_SERVER_PORT"`
	GrpcInventoryClientPort string `envconfig:"GRPC_INVENTORY_CLIENT_PORT"`
	PostgresDBUser          string `envconfig:"POSTGRES_DBUSER"`
	PostgresDBPassword      string `envconfig:"POSTGRES_DBPASSWORD"`
	PostgresDBName          string `envconfig:"POSTGRES_DBNAME"`
	PostgresDBSSLMode       string `envconfig:"POSTGRES_DBSSLMODE"`
	PostgresDBPort          string `envconfig:"POSTGRES_DBPORT"`
}

func ConfigEnv() *Env {
	var env Env
	err := envconfig.Process("SIGMA_INVENTORY", &env)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &env
}

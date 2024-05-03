package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	HttpPort           string `envconfig:"HTTP_PORT"`
	JWTSecret          string `envconfig:"JWT_SECRET"`
	PostgresDBUser     string `envconfig:"POSTGRES_DBUSER"`
	PostgresDBPassword string `envconfig:"POSTGRES_DBPASSWORD"`
	PostgresDBName     string `envconfig:"POSTGRES_DBNAME"`
	PostgresDBPort     string `envconfig:"POSTGRES_DBPORT"`
	PostgresDBHost     string `envconfig:"POSTGRES_DBHOST"`
	PostgresDBSSLMode  string `envconfig:"POSTGRES_DBSSLMODE"`
	AerospikeHostname  string `envconfig:"AEROSPIKE_HOSTNAME"`
	AerospikePort      int    `envconfig:"AEROSPIKE_PORT"`
}

func ConfigEnv() *Env {
	var env Env
	err := envconfig.Process("SIGMA_USER", &env)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &env
}

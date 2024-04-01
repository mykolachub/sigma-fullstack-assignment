package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	Port               string `envconfig:"SIGMA_APP_PORT"`
	JWTSecret          string `envconfig:"JWT_SECRET"`
	PostgresDBUser     string `envconfig:"POSTGRES_DBUSER"`
	PostgresDBPassword string `envconfig:"POSTGRES_DBPASSWORD"`
	PostgresDBName     string `envconfig:"POSTGRES_DBNAME"`
	PostgresDBSSLMode  string `envconfig:"POSTGRES_DBSSLMODE"`
	AerospikeHostname  string `envconfig:"AEROSPIKE_HOSTNAME"`
	AerospikePort      int    `envconfig:"AEROSPIKE_PORT"`
}

func ConfigEnv() *Env {
	var env Env
	err := envconfig.Process("SIGMA", &env)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &env
}

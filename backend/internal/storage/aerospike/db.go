package aerospike

import "github.com/aerospike/aerospike-client-go/v7"

type AerospikeConfig struct {
	Hostname string
	Port     int
}

func InitDBConnection(dbConfig AerospikeConfig) (*aerospike.Client, aerospike.Error) {
	return aerospike.NewClient(dbConfig.Hostname, dbConfig.Port)
}

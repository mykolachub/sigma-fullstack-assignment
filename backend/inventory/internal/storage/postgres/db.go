package postgres

import (
	"database/sql"
	"fmt"
)

type PostgresConfig struct {
	DBUser     string
	DBName     string
	DBPassword string
	DBSSLMode  string
	DBPort     string
}

func InitDBConnection(dbConfig PostgresConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s port=%s sslmode=%s", dbConfig.DBUser, dbConfig.DBName, dbConfig.DBPassword, dbConfig.DBPort, dbConfig.DBSSLMode)
	return sql.Open("postgres", connStr)
}

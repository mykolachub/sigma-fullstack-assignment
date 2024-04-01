package postgres

import (
	"database/sql"
	"fmt"
)

type DbConfig struct {
	DBUser     string
	DBName     string
	DBPassword string
	DBSSLMode  string
}

func InitDBConnection(dbConfig DbConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", dbConfig.DBUser, dbConfig.DBName, dbConfig.DBPassword, dbConfig.DBSSLMode)
	return sql.Open("postgres", connStr)
}

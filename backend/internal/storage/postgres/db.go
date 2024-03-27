package postgres

import (
	"database/sql"
	"fmt"
	"sigma-test/config"
)

func InitDBConnection() (*sql.DB, error) {
	env := config.ConfigEnv()
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", env.DBUser, env.DBName, env.DBPassword, env.DBSSLMode)
	return sql.Open("postgres", connStr)
}

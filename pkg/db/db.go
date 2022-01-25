package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func GetConnection(host string, port int, user, password, dbName string) error {
	if password == "" {
		password = `''`
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = database.Ping()
	if err != nil {
		return err
	}
	DB = database
	return nil
}

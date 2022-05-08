package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

var DB *sql.DB

func SetStats() {
	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(2)
	DB.SetConnMaxIdleTime(time.Minute * 3)
	DB.SetConnMaxLifetime(time.Minute * 3)
}

func GetConnection(host string, port int, user, password, dbName string) error {
	if DB.Ping() == nil {
		return nil
	}

	if password == "" {
		password = `''`
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
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
	SetStats()
	return nil
}

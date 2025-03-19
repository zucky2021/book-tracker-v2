package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	MAX_RETRIES int = 5
)

func SetupDatabase() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)

	var db *sql.DB
	var err error

	// リトライ処理
	for i := 0; i < MAX_RETRIES; i++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				return db, nil
			}
		}
		fmt.Printf("Database connection failed. Retrying in 5 seconds... (%d/5)\n", i+1)
		time.Sleep(5 * time.Second)
	}

	return nil, fmt.Errorf("could not connect to the database: %v", err)
}

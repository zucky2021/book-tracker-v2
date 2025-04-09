package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	maxRetries int = 5
)

// DB接続確立
func ConnectDB() *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	var db *sql.DB
	var err error
	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("MySQLへの接続失敗 (試行 %d/%d): %v", i+1, maxRetries, err)
		} else {
			pingErr := db.Ping()
			if pingErr != nil {
				log.Printf("MySQLへの接続失敗 (試行 %d/%d): %v", i+1, maxRetries, pingErr)
			} else {
				log.Println("MySQLに接続成功")
				return db
			}
		}
		time.Sleep(5 * time.Second)
	}

	log.Fatalf("MySQLへの接続に失敗しました: %v", err)
	return nil
}

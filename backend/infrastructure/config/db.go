package config

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	maxRetries int = 5
)

var (
	db   *gorm.DB
	once sync.Once
)

type DBConfig struct {
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbName     string
}

func GetDB() *gorm.DB {
	return connectDB(&DBConfig{
		dbHost:     os.Getenv("DB_HOST"),
		dbPort:     os.Getenv("DB_PORT"),
		dbUser:     os.Getenv("DB_USER"),
		dbPassword: os.Getenv("DB_PASSWORD"),
		dbName:     os.Getenv("DB_NAME"),
	})
}

func connectDB(config *DBConfig) *gorm.DB {
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.dbUser, config.dbPassword, config.dbHost, config.dbPort, config.dbName)

		var err error
		for i := range maxRetries {
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				log.Printf("MySQLへの接続失敗 (試行 %d/%d): %v", i+1, maxRetries, err)
			} else {
				log.Println("MySQLへの接続に成功")
				return
			}

			time.Sleep(5 * time.Second)
		}

		log.Fatalf("MySQLへの接続に失敗しました: %v", err)

	})

	return db
}

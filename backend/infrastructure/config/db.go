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
	maxRetries      int = 5
	retryInterval       = 5 * time.Second
	maxIdleConns        = 10
	maxOpenConns        = 10
	maxConnLifetime     = 10 * time.Minute
)

var (
	dbConnections *DBConnections
	once          sync.Once
)

type DBConfig struct {
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbName     string
}

type DBConnections struct {
	Writer *gorm.DB
	Reader *gorm.DB
}

func GetDBConnections() *DBConnections {
	once.Do(func() {
		writer, err := connectDB(&DBConfig{
			dbHost:     os.Getenv("DB_HOST"),
			dbPort:     os.Getenv("DB_PORT"),
			dbUser:     os.Getenv("DB_USER"),
			dbPassword: os.Getenv("DB_PASSWORD"),
			dbName:     os.Getenv("DB_NAME"),
		})
		if err != nil {
			log.Fatalf("Failed to connect to writer DB: %v", err)
		}

		reader, err := connectDB(&DBConfig{
			dbHost:     os.Getenv("DB_READER_HOST"),
			dbPort:     os.Getenv("DB_READER_PORT"),
			dbUser:     os.Getenv("DB_READER_USER"),
			dbPassword: os.Getenv("DB_READER_PASSWORD"),
			dbName:     os.Getenv("DB_NAME"),
		})
		if err != nil {
			log.Fatalf("Failed to connect to reader DB: %v", err)
		}

		dbConnections = &DBConnections{
			Writer: writer,
			Reader: reader,
		}
	})
	return dbConnections
}

func connectDB(config *DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.dbUser, config.dbPassword, config.dbHost, config.dbPort, config.dbName)

	var db *gorm.DB
	var err error

	for i := range maxRetries {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Printf("MySQLへの接続失敗 (試行 %d/%d): %v", i+1, maxRetries, err)
		} else {
			sqlDB, err := db.DB()
			if err != nil {
				return nil, fmt.Errorf("failed to get sql.DB from gorm.DB: %v", err)
			}
			sqlDB.SetMaxIdleConns(maxIdleConns)
			sqlDB.SetMaxOpenConns(maxOpenConns)
			sqlDB.SetConnMaxLifetime(maxConnLifetime)
			log.Println("MySQLへの接続に成功")

			return db, nil
		}

		time.Sleep(5 * time.Second)
	}

	return nil, err
}

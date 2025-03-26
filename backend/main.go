package main

import (
	"fmt"
	"log"

	"backend/config"
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.SetupDatabase()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	r := gin.Default()

	config.SetupCORS(r)

	// ルーティング設定
	r.GET("/health", handlers.HealthCheckHandler(db))
	r.GET("/api/books", handlers.GetBooksHandler)

	// サーバーを起動（ポート8080）
	fmt.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}

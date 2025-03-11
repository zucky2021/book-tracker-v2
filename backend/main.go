package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORS設定 (フロントエンドと通信できるようにする)
	r.Use(cors.Default())

	// 書籍検索APIエンドポイント
	r.GET("/api/books", func(c *gin.Context) {
		query := c.Query("q")
		if query == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
			return
		}

		apiKey := os.Getenv("GOOGLE_BOOKS_API_KEY")
		url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=%s&key=%s", query, apiKey)

		resp, err := http.Get(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from Google Books API"})
			return
		}
		defer resp.Body.Close()

		var result map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			log.Printf("Error decoding response from Google Books API: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response from Google Books API"})
			return
		}
		c.JSON(http.StatusOK, result)
	})

	log.Println("Started server on :8080");
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
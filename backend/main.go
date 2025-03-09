package main

import (
	"encoding/json"
	"fmt"
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
		json.NewDecoder(resp.Body).Decode(&result)
		c.JSON(http.StatusOK, result)
	})

	r.Run(":8080") // サーバーを8080ポートで起動
}

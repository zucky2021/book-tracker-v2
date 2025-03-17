package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FIXME: 別課題で実装

func GetBooksHandler(apiKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("Authorization")
		if authToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization token is required"})
			return
		}

		shelfID := "7" // My Library の「マイブックス」を指定（変更可能）
		url := fmt.Sprintf("https://www.googleapis.com/books/v1/mylibrary/bookshelves/%s/volumes?key=%s", shelfID, apiKey)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create request"})
			return
		}

		req.Header.Add("Authorization", authToken)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil || resp.StatusCode != http.StatusOK {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "API request failed"})
			return
		}

		defer resp.Body.Close()

		c.JSON(http.StatusOK, gin.H{"message": "Books fetched successfully"})
	}
}

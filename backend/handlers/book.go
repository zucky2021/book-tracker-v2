package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetBookshelvesHandler(c *gin.Context) {
	userId := c.Query("userId")
	shelfId := c.Query("shelfId")

	if userId == "" || shelfId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId and shelfId are required"})
		return
	}

	// Google Books APIのURLを構築
	apiKey := os.Getenv("GOOGLE_BOOKS_API_KEY") // 環境変数からAPIキーを取得
	if apiKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Google Books API key is not set"})
		return
	}

	url := fmt.Sprintf(
		"https://www.googleapis.com/books/v1/users/%s/bookshelves/%s?key=%s",
		userId, shelfId, apiKey,
	)

	// Google Books APIにリクエストを送信
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch data from Google Books API"})
	}
	defer resp.Body.Close()

	// Google Books APIのレスポンスをそのまま返す
	c.Status(resp.StatusCode)
	c.Header("Content-Type", resp.Header.Get("Content-Type"))
	_, err = io.Copy(c.Writer, resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from Google Books API"})
	}
}

func GetBooksHandler(c *gin.Context) {
	userId := c.Query("userId")
	shelfId := c.Query("shelfId")
	startIndex := c.DefaultQuery("startIndex", "0")
	maxResults := c.DefaultQuery("maxResults", "40")

	if userId == "" || shelfId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId and shelfId are required"})
		return
	}

	// Google Books APIのURLを構築
	apiKey := os.Getenv("GOOGLE_BOOKS_API_KEY") // 環境変数からAPIキーを取得
	if apiKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Google Books API key is not set"})
		return
	}

	url := fmt.Sprintf(
		"https://www.googleapis.com/books/v1/users/%s/bookshelves/%s/volumes?startIndex=%s&maxResults=%s&key=%s",
		userId, shelfId, startIndex, maxResults, apiKey,
	)

	// Google Books APIにリクエストを送信
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch data from Google Books API"})
	}
	defer resp.Body.Close()

	// Google Books APIのレスポンスをそのまま返す
	c.Status(resp.StatusCode)
	c.Header("Content-Type", resp.Header.Get("Content-Type"))
	_, err = io.Copy(c.Writer, resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from Google Books API"})
	}
}

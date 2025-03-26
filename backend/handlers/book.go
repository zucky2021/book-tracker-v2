package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Google Books APIにリクエストを送信するヘルパー関数
func fetchFromGoogleBooksAPI(c *gin.Context, endpoint string) {
	apiKey := os.Getenv("GOOGLE_BOOKS_API_KEY") // 環境変数からAPIキーを取得
	if apiKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Google Books API key is not set"})
		return
	}

	// 完全なURLを構築
	url := fmt.Sprintf("%s&key=%s", endpoint, apiKey)

	// Google Books APIにリクエストを送信
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from Google Books API"})
		return
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

func GetBookshelvesHandler(c *gin.Context) {
	userId := c.Query("userId")
	shelfId := c.Query("shelfId")

	if userId == "" || shelfId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId and shelfId are required"})
		return
	}

	endpoint := fmt.Sprintf(
		"https://www.googleapis.com/books/v1/users/%s/bookshelves/%s",
		userId, shelfId,
	)

	fetchFromGoogleBooksAPI(c, endpoint)
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

	endpoint := fmt.Sprintf(
		"https://www.googleapis.com/books/v1/users/%s/bookshelves/%s/volumes?startIndex=%s&maxResults=%s",
		userId, shelfId, startIndex, maxResults,
	)

	fetchFromGoogleBooksAPI(c, endpoint)
}

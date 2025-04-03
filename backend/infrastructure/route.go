package infrastructure

import (
	"backend/controller"
	"backend/presenter"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, bookController *controller.BookController, bookPresenter *presenter.BookPresenter) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/api/books", func(c *gin.Context) {
		queryParams := map[string]string{
			"userId":     c.Query("userId"),
			"shelfId":    c.Query("shelfId"),
			"startIndex": c.DefaultQuery("startIndex", "0"),
			"maxResults": c.DefaultQuery("maxResults", "40"),
		}

		books, err := bookController.GetBooks(queryParams)
		if err != nil {

			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		bookPresenter.PresentBooks(c, books, err)
	})
}

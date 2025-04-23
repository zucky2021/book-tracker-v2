package infrastructure

import (
	"backend/controller"
	"backend/presenter"

	"github.com/gin-gonic/gin"
)

func InitRouter(
	r *gin.Engine,
	bookController *controller.BookController,
	bookPresenter *presenter.BookPresenter,
	bookshelfController *controller.BookshelfController,
	bookshelfPresenter *presenter.BookshelfPresenter,
	memoController *controller.MemoController,
) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":    "ok",
			"timestamp": c.Request.Header.Get("X-Request-Timestamp"),
		})
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

	r.GET("/api/bookshelf", func(c *gin.Context) {
		queryParams := map[string]string{
			"userId":  c.Query("userId"),
			"shelfId": c.Query("shelfId"),
		}

		bookshelf, err := bookshelfController.GetBookshelf(queryParams)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		bookshelfPresenter.PresentBookshelf(c, *bookshelf)
	})

	r.GET("/api/memo", memoController.GetMemo)
}

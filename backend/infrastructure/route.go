package infrastructure

import (
	"backend/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(
	r *gin.Engine,
	bookController *controller.BookController,
	bookshelfController *controller.BookshelfController,
	memoController *controller.MemoController,
) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":    "ok",
			"timestamp": c.Request.Header.Get("X-Request-Timestamp"),
		})
	})

	r.GET("/api/books", bookController.GetBooks)

	r.GET("/api/bookshelf", bookshelfController.GetBookshelf)

	r.POST("/api/memo", memoController.CreateMemo)
	r.GET("/api/memo/:memoId", memoController.GetMemo)
	r.PUT("/api/memo/:memoId", memoController.UpdateMemo)
	r.DELETE("api/memo/:memoId", memoController.DeleteMemo)
}

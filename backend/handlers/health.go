package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := db.Ping(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Database connection failed",
				"error":  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"db":     "Database connection successful",
		})
	}
}

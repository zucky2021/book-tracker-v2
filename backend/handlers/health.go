package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := db.Ping(); err != nil {
			log.Printf("Database health check failed: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"message": "Database connection failed",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"message": "Database connection successful",
		})
	}
}

package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

var PostgresDB *sql.DB

func PostgresCount(c *gin.Context) {
	var count int

	err := PostgresDB.QueryRow(
		"SELECT COUNT(*) FROM customers",
	).Scan(&count)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"database": "postgres",
		"count":    count,
	})
}

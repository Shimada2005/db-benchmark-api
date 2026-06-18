package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func BenchmarkCount(c *gin.Context) {

	var postgresCount int
	var mysqlCount int

	// PostgreSQL
	start := time.Now()

	err := PostgresDB.QueryRow(
		"SELECT COUNT(*) FROM customers",
	).Scan(&postgresCount)

	postgresTime := time.Since(start).Milliseconds()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// MySQL
	start = time.Now()

	err = MySQLDB.QueryRow(
		"SELECT COUNT(*) FROM customers",
	).Scan(&mysqlCount)

	mysqlTime := time.Since(start).Milliseconds()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"postgres": gin.H{
			"count":   postgresCount,
			"time_ms": postgresTime,
		},
		"mysql": gin.H{
			"count":   mysqlCount,
			"time_ms": mysqlTime,
		},
	})
}

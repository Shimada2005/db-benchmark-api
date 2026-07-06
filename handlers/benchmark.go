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

func BenchmarkSearch(c *gin.Context) {

	country := c.Query("country")

	// PostgreSQL
	start := time.Now()

	var postgresCount int

	err := PostgresDB.QueryRow(
		"SELECT COUNT(*) FROM customers WHERE country = $1",
		country,
	).Scan(&postgresCount)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	postgresTime := time.Since(start).Milliseconds()

	// MySQL
	start = time.Now()

	var mysqlCount int

	err = MySQLDB.QueryRow(
		"SELECT COUNT(*) FROM customers WHERE country = ?",
		country,
	).Scan(&mysqlCount)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	mysqlTime := time.Since(start).Milliseconds()

	c.JSON(http.StatusOK, gin.H{
		"country": country,
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
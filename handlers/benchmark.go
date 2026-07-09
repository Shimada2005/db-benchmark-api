package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const benchmarkLoop = 100

func BenchmarkCount(c *gin.Context) {
	var postgresCount int
	var mysqlCount int

	var postgresTotal int64
	var mysqlTotal int64

	// PostgreSQL
	for i := 0; i < benchmarkLoop; i++ {
		start := time.Now()

		err := PostgresDB.QueryRow(
			"SELECT COUNT(*) FROM customers",
		).Scan(&postgresCount)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		postgresTotal += time.Since(start).Microseconds()
	}

	// MySQL
	for i := 0; i < benchmarkLoop; i++ {
		start := time.Now()

		err := MySQLDB.QueryRow(
			"SELECT COUNT(*) FROM customers",
		).Scan(&mysqlCount)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		mysqlTotal += time.Since(start).Microseconds()
	}

	c.JSON(http.StatusOK, gin.H{
		"loops": benchmarkLoop,
		"postgres": gin.H{
			"count":      postgresCount,
			"average_us": postgresTotal / benchmarkLoop,
		},
		"mysql": gin.H{
			"count":      mysqlCount,
			"average_us": mysqlTotal / benchmarkLoop,
		},
	})
}

func BenchmarkSearch(c *gin.Context) {
	country := c.Query("country")

	var postgresCount int
	var mysqlCount int

	var postgresTotal int64
	var mysqlTotal int64

	// PostgreSQL
	for i := 0; i < benchmarkLoop; i++ {
		start := time.Now()

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

		postgresTotal += time.Since(start).Microseconds()
	}

	// MySQL
	for i := 0; i < benchmarkLoop; i++ {
		start := time.Now()

		err := MySQLDB.QueryRow(
			"SELECT COUNT(*) FROM customers WHERE country = ?",
			country,
		).Scan(&mysqlCount)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		mysqlTotal += time.Since(start).Microseconds()
	}

	c.JSON(http.StatusOK, gin.H{
		"country": country,
		"loops":   benchmarkLoop,
		"postgres": gin.H{
			"count":      postgresCount,
			"average_us": postgresTotal / benchmarkLoop,
		},
		"mysql": gin.H{
			"count":      mysqlCount,
			"average_us": mysqlTotal / benchmarkLoop,
		},
	})
}

func BenchmarkPrimarySearch(c *gin.Context) {
	id := c.Query("id")

	var postgresID string
	var mysqlID string

	var postgresTotal int64
	var mysqlTotal int64

	// PostgreSQL
	for i := 0; i < benchmarkLoop; i++ {
		start := time.Now()

		err := PostgresDB.QueryRow(
			"SELECT customer_id FROM customers WHERE customer_id = $1",
			id,
		).Scan(&postgresID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		postgresTotal += time.Since(start).Microseconds()
	}

	// MySQL
	for i := 0; i < benchmarkLoop; i++ {
		start := time.Now()

		err := MySQLDB.QueryRow(
			"SELECT customer_id FROM customers WHERE customer_id = ?",
			id,
		).Scan(&mysqlID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		mysqlTotal += time.Since(start).Microseconds()
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"loops": benchmarkLoop,
		"postgres": gin.H{
			"average_us": postgresTotal / benchmarkLoop,
		},
		"mysql": gin.H{
			"average_us": mysqlTotal / benchmarkLoop,
		},
	})
}

func BenchmarkCRUD(c *gin.Context) {
	const benchmarkLoop = 100

	var postgresInsertTotal, postgresUpdateTotal, postgresDeleteTotal int64
	var mysqlInsertTotal, mysqlUpdateTotal, mysqlDeleteTotal int64

	// -----------------
	// PostgreSQL
	// -----------------
	for i := 0; i < benchmarkLoop; i++ {

		id := fmt.Sprintf("benchmark_%d", time.Now().UnixNano())
		email := id + "@example.com"

		// INSERT
		start := time.Now()

		_, err := PostgresDB.Exec(`
			INSERT INTO customers
			(customer_id, first_name, last_name, company, city, country,
			 phone1, phone2, email, subscription_date, website)
			VALUES
			($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
		`,
			id,
			"Benchmark",
			"User",
			"TestCompany",
			"Osaka",
			"Japan",
			"0000000000",
			"0000000000",
			email,
			time.Now(),
			"https://example.com",
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "postgres insert: " + err.Error()})
			return
		}

		postgresInsertTotal += time.Since(start).Microseconds()

		// UPDATE
		start = time.Now()

		_, err = PostgresDB.Exec(
			`UPDATE customers SET company = $1 WHERE customer_id = $2`,
			"UpdatedCompany",
			id,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "postgres update: " + err.Error()})
			return
		}

		postgresUpdateTotal += time.Since(start).Microseconds()

		// DELETE
		start = time.Now()

		_, err = PostgresDB.Exec(
			`DELETE FROM customers WHERE customer_id = $1`,
			id,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "postgres delete: " + err.Error()})
			return
		}

		postgresDeleteTotal += time.Since(start).Microseconds()
	}

	// -----------------
	// MySQL
	// -----------------
	for i := 0; i < benchmarkLoop; i++ {

		id := fmt.Sprintf("benchmark_%d", time.Now().UnixNano())
		email := id + "@example.com"

		// INSERT
		start := time.Now()

		_, err := MySQLDB.Exec(`
			INSERT INTO customers
			(customer_id, first_name, last_name, company, city, country,
			 phone1, phone2, email, subscription_date, website)
			VALUES
			(?,?,?,?,?,?,?,?,?,?,?)
		`,
			id,
			"Benchmark",
			"User",
			"TestCompany",
			"Osaka",
			"Japan",
			"0000000000",
			"0000000000",
			email,
			time.Now(),
			"https://example.com",
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "mysql insert: " + err.Error()})
			return
		}

		mysqlInsertTotal += time.Since(start).Microseconds()

		// UPDATE
		start = time.Now()

		_, err = MySQLDB.Exec(
			`UPDATE customers SET company = ? WHERE customer_id = ?`,
			"UpdatedCompany",
			id,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "mysql update: " + err.Error()})
			return
		}

		mysqlUpdateTotal += time.Since(start).Microseconds()

		// DELETE
		start = time.Now()

		_, err = MySQLDB.Exec(
			`DELETE FROM customers WHERE customer_id = ?`,
			id,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "mysql delete: " + err.Error()})
			return
		}

		mysqlDeleteTotal += time.Since(start).Microseconds()
	}

	c.JSON(http.StatusOK, gin.H{
		"loops": benchmarkLoop,
		"postgres": gin.H{
			"insert_us": postgresInsertTotal / benchmarkLoop,
			"update_us": postgresUpdateTotal / benchmarkLoop,
			"delete_us": postgresDeleteTotal / benchmarkLoop,
		},
		"mysql": gin.H{
			"insert_us": mysqlInsertTotal / benchmarkLoop,
			"update_us": mysqlUpdateTotal / benchmarkLoop,
			"delete_us": mysqlDeleteTotal / benchmarkLoop,
		},
	})
}
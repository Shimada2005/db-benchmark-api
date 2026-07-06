package handlers

import (
	"database/sql"
	"net/http"

	"db-benchmark-api/models"

	"github.com/gin-gonic/gin"
)

var MySQLDB *sql.DB

func MySQLCount(c *gin.Context) {
	var count int

	err := MySQLDB.QueryRow(
		"SELECT COUNT(*) FROM customers",
	).Scan(&count)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"database": "mysql",
		"count":    count,
	})
}

func MySQLSearch(c *gin.Context) {

	country := c.Query("country")

	rows, err := MySQLDB.Query(`
		SELECT
			customer_id,
			first_name,
			last_name,
			company,
			city,
			country,
			phone1,
			phone2,
			email,
			subscription_date,
			website
		FROM customers
		WHERE country = ?
		LIMIT 100
	`, country)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()

	var customers []models.Customer

	for rows.Next() {

		var customer models.Customer

		err := rows.Scan(
			&customer.CustomerID,
			&customer.FirstName,
			&customer.LastName,
			&customer.Company,
			&customer.City,
			&customer.Country,
			&customer.Phone1,
			&customer.Phone2,
			&customer.Email,
			&customer.SubscriptionDate,
			&customer.Website,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		customers = append(customers, customer)
	}

	c.JSON(http.StatusOK, gin.H{
		"database": "mysql",
		"count":    len(customers),
		"data":     customers,
	})
}
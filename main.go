package main

import (
	"db-benchmark-api/config"
	"db-benchmark-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	postgresDB, err := config.ConnectPostgres()
	if err != nil {
		panic(err)
	}

	handlers.PostgresDB = postgresDB

	mysqlDB, err := config.ConnectMySQL()
	if err != nil {
		panic(err)
	}

	handlers.MySQLDB = mysqlDB

	r := gin.Default()

	// health
	r.GET("/health", handlers.Health)

	// postgres
	r.GET("/postgres/count", handlers.PostgresCount)

	// mysql
	r.GET("/mysql/count", handlers.MySQLCount)

	// benchmark
	r.GET("/benchmark/count", handlers.BenchmarkCount)

	//search
	r.GET("/postgres/search", handlers.PostgresSearch)
	r.GET("/mysql/search", handlers.MySQLSearch)
	r.GET("/benchmark/search", handlers.BenchmarkSearch)

	r.Run(":8080")
}

package main

import (
	"db-benchmark-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/health", handlers.Health)
	r.Run(":8080")
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/health", healthcheck)

	router.Run("localhost:8080")
}

// healthcheck responds that the server is running
func healthcheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Album Server is running!")
}

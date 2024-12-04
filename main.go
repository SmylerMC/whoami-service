package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func routeGetIp(c *gin.Context) {
	ip := c.ClientIP()
	c.Data(http.StatusOK, "text", []byte(ip))
}

func main() {
	fmt.Println("Starting whoami service...")
	router := gin.Default()
	router.GET("/ip", routeGetIp)
	router.Run("0.0.0.0:8080")
}

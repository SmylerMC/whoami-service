package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func routeGetIp(c *gin.Context) {
	ip := c.ClientIP()
	c.Data(http.StatusOK, "text", []byte(ip))
}

func main() {
	router := gin.Default()
	router.GET("/ip", routeGetIp)
	router.Run("localhost:8080")
}

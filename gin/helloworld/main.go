package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/user/:name", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello %s", c.Param("name"))
	})

	r.Run(":6666")
}

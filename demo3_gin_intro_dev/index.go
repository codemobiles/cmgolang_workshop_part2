package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleBookRequest(c *gin.Context) {
	from, to := c.Param("from"), c.Param("to")
	vichicle := c.Param("vichicle")
	c.JSON(http.StatusOK, gin.H{"result": "ok", "from": from, "to": to, "vichicle":vichicle})
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("1234"))
	})

	r.GET("/profile", func(c *gin.Context) {
		c.Data(200, "text/html; charset=utf-8", []byte("profile"))
	})

	r.GET("/login", func(c *gin.Context) {
		username, password := c.Query("username"), c.Query("password")
		c.JSON(http.StatusOK, gin.H{"result": "ok", "username": username, "password": password})
	})

	r.GET("/book/:from/:to/:vichicle", handleBookRequest)

	r.Run(":85")
}

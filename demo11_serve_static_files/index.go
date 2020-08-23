package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/static/file", func(c *gin.Context) {
		c.File("static/file.go")
	})

	var fs http.FileSystem = // ...
	router.GET("/fs/file", func(c *gin.Context) {
		c.FileFromFS("fs/file.go", fs)
	})
}

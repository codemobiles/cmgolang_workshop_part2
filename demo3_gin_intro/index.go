package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Data(200, "text/html; charset=utf-8", []byte("Root"))
	})

	r.GET("/profile", func(c *gin.Context) {
		c.Data(200, "text/html; charset=utf-8", []byte("profile"))
	})

	r.GET("/login", func(c *gin.Context) {
		username, password := c.Query("username"), c.Query("password")
		c.JSON(http.StatusOK, gin.H{"result": "ok", "username": username, "password": password})
	})

	r.GET("/book/:from/:to", func(c *gin.Context) {
		from, to := c.Param("from"), c.Param("to")
		c.JSON(http.StatusOK, gin.H{"result": "ok", "from": from, "to": to})
	})

	r.POST("/login", func(c *gin.Context) {
		var form LoginForm
		if c.ShouldBind(&form) == nil {
			if form.Username == "admin" && form.Password == "1234" {
				msg := fmt.Sprintf("you are logged with: %s,%s", form.Username, form.Password)
				c.JSON(200, gin.H{"status": msg})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		} else {
			c.JSON(401, gin.H{"status": "unable to bind data"})
		}
	})

	r.Run(":85")
}

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.String(http.StatusOK, "Login")
}

func Register(c *gin.Context) {
	c.String(http.StatusOK, "Register")
}

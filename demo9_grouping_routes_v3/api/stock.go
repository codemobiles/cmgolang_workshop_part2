package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListProduct(c *gin.Context) {
	c.String(http.StatusOK, "List Product")
}

func CreateProduct(c *gin.Context) {
	c.String(http.StatusOK, "Create Product")
}

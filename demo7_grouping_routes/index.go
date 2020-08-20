package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	c.String(http.StatusOK, "Login")
}

func register(c *gin.Context) {
	c.String(http.StatusOK, "Register")
}

//----------- Stock
func listProduct(c *gin.Context) {
	c.String(http.StatusOK, "List Product")
}

func createProduct(c *gin.Context) {
	c.String(http.StatusOK, "Create Product")
}

func main() {
	router := gin.Default()

	// Authen group
	authenAPI := router.Group("/authen")
	{
		authenAPI.GET("/login", login)
		authenAPI.GET("/register", register)
		authenAPI.GET("/profile", func(c *gin.Context) {
			c.String(http.StatusOK, "Profile")
		})
	}

	stockAPI := router.Group("/stock")
	{
		stockAPI.GET("/list", listProduct)
		stockAPI.GET("/create", createProduct)
	}

	router.Run(":85")

}

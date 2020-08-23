package api

import (
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	// Authen group
	authenAPI := router.Group("/authen")
	{
		authenAPI.GET("/login", Login)
		authenAPI.GET("/register", Register)
	}

	stockAPI := router.Group("/stock")
	{
		stockAPI.GET("/list", ListProduct)
		stockAPI.GET("/create", CreateProduct)
	}
}

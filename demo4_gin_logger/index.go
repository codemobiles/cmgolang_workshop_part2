package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	runningDir, _ := os.Getwd()

	errlogfile, _ := os.OpenFile(fmt.Sprintf("%s/gin_error.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	accesslogfile, _ := os.OpenFile(fmt.Sprintf("%s/gin_access.log", runningDir), os.O_CREATE|os.O_WRONLY, 0600)

	gin.DefaultErrorWriter = errlogfile
	gin.DefaultWriter = accesslogfile
	// gin.DisableConsoleColor()

	// r.Use(gin.Logger())

	// custom format logger
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	r.GET("/", func(c *gin.Context) {
		c.Data(200, "text/html; charset=utf-8", []byte("Root"))
	})

	r.GET("/profile", func(c *gin.Context) {
		c.Data(200, "text/html; charset=utf-8", []byte("profile"))
	})

	r.GET("/error", func(c *gin.Context) {
		errlogfile.WriteString("error\n")
		c.Data(200, "text/html; charset=utf-8", []byte("error"))
	})

	r.Run(":85")
}

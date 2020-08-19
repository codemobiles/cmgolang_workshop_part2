package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	runningDir, _ := os.Getwd()
	count := 0

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// single file
		count = count + 1
		file, _ := c.FormFile("file")
		extension := filepath.Ext(file.Filename)

		c.SaveUploadedFile(file, fmt.Sprintf("%s/uploaded/%d%s", runningDir, count, extension))
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	router.Run(":85")
}

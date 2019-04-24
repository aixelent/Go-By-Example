package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.WriteString("index router")
}

func main() {
	gonic := gin.Default()
	gonic.GET("/", Index)
	gonic.Run(":8080")
}

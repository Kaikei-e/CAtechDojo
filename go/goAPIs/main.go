package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"test": "hello world",
		})
	})

	router.Run(":8085")
}

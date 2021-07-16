package main

import (
	manageuser "app/manageUser"
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

	router.GET("/user/create", manageuser.CreateUser)

	router.Run(":8085")
}

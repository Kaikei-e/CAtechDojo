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

	router.GET("/user/get", manageuser.GetUserToken)
	router.POST("/user/create", manageuser.CreateUser)
	router.PUT("/user/update", manageuser.UpdateUserData)


	router.Run(":8085")
}

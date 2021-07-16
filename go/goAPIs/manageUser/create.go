package manageuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"test": "hello world",
	})
}

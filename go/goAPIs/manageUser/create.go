package manageuser

import (
	jwtauth "app/JWTauth"
	userstruct "app/userStruct"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context){
	theToken := jwtauth.JWTTokenMaker(ctx)

	var theUser userstruct.Users

	if err := ctx.ShouldBindJSON(&theUser); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}


	ctx.JSON(http.StatusOK, gin.H{
		"token": theToken,
	})
}

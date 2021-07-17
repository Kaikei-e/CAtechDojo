package manageuser

import (
	jwtauth "app/JWTauth"
	userstruct "app/userStruct"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserToken(ctx *gin.Context){
	theToken := ctx.Request.Header.Get("x-token")
	var user userstruct.Users

	if userParsed, err :=  jwtauth.ParseJWT(theToken); err != nil{
		log.Fatalln(err)
	}else{
		user = *userParsed
	}

	ctx.JSON(http.StatusOK, gin.H{
		"name": user.Name,
	})

}

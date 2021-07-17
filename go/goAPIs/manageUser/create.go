package manageuser

import (
	dbmanage "app/DBmanage"
	jwtauth "app/JWTauth"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context){

	theUser := jwtauth.JWTTokenMaker(ctx)


	db := dbmanage.DBconnection
	defer db().Close()

	log.Println(theUser)
	db().LogMode(true)

	db().Create(&theUser)


	ctx.JSON(http.StatusOK, gin.H{
		"token": theUser.Token,
	})
}





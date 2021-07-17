package manageuser

import (
	dbmanage "app/DBmanage"
	jwtauth "app/JWTauth"
	userstruct "app/userStruct"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUserData(ctx *gin.Context){
	theToken := ctx.Request.Header.Get("x-token")
	var user userstruct.Users
	var userUpdated userstruct.Users



	if err := ctx.ShouldBind(&userUpdated); err != nil{
		ctx.String(http.StatusBadRequest, "Bad request")
		log.Fatalln(err)

	}

	if userParsed, err :=  jwtauth.ParseJWT(theToken); err != nil{
		log.Fatalln(err)
	}else{
		user = *userParsed
	}

	user.Name = userUpdated.Name
	log.Println(user.Name)
	user.Token = theToken

	db := dbmanage.DBconnection
	defer db().Close()

	log.Println(theToken)
	db().LogMode(true)


	db().Table("users").Model(&user).Where("token=?", user.Token).Update("name", user.Name)

	ctx.Writer.WriteHeader(http.StatusOK)
}



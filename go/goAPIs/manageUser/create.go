package manageuser

import (
	jwtauth "app/JWTauth"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context){
	theToken := jwtauth.JWTTokenMaker(ctx)

	jwtauth.ParseJWT(theToken)

}

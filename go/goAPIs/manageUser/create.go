package manageuser

import (
	jwtauth "app/JWTauth"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context){
	jwtauth.JWTTokenMaker(ctx)

}

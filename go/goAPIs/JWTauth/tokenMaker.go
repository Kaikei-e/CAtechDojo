package jwtauth

import (
	userstruct "app/userStruct"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func JWTTokenMaker(ctx *gin.Context) string {
	user := userstruct.User{}
	if err := ctx.Bind(&user); err != nil{
		ctx.String(http.StatusBadRequest, "Bad request")
		log.Fatalln(err)

	}

	jToken := jwt.New(jwt.SigningMethodHS256)

	claims := jToken.Claims.(jwt.MapClaims)
	claims["name"] = user.Name
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 720).Unix()

	if err := godotenv.Load(); err != nil{
		log.Fatalln("Error loading .env file" +  err.Error())

	}

	secret := []byte(os.Getenv("SECRETKEY"))

	var tokenSignature string

	if tokenSig,  err := jToken.SignedString(secret); err != nil{
		log.Fatalln(err)
		}else{
		tokenSignature = tokenSig

	}

	return tokenSignature

}

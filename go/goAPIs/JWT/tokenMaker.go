package jwt

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func JWTTokenMaker(ctx *gin.Context, name string) error{
	jToken := jwt.New(jwt.SigningMethodHS256)

	claims := jToken.Claims.(jwt.MapClaims)
	claims["name"] = name
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 720).Unix()

	if err := godotenv.Load(); err != nil{
		log.Fatalln("Error loading .env file" +  err.Error())

	}

	secret := []byte(os.Getenv("SECRETKEY"))

	var tokenSignature string

	if tokenSig,  err := jToken.SignedString(secret); err != nil{
		log.Fatalln(err)
		return err
		}else{
		tokenSignature = tokenSig

	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"token": tokenSignature,
	})


}

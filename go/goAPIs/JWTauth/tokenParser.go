package jwtauth

import (
	userstruct "app/userStruct"
	"fmt"
	"log"
	"os"

	"github.com/form3tech-oss/jwt-go"
	"github.com/joho/godotenv"
)

func ParseJWT(jwtToken string) (*userstruct.User, error){
	if err := godotenv.Load(); err != nil{
		log.Fatalln("Error loading .env file" +  err.Error())

	}

	secret := []byte(os.Getenv("SECRETKEY"))

	parsedToken, err := jwt.Parse(jwtToken, func (theToken *jwt.Token) (interface{}, error) {
		if _, isOK := theToken.Method.(*jwt.SigningMethodHMAC) ; !isOK{
			return "", fmt.Errorf("Sigining method is wrong : %v", theToken.Header["alg"])

		}
		return []byte(secret), nil
	})

	if err != nil {
		log.Fatalln(err)
	}

	if parsedToken == nil {
		return nil, fmt.Errorf("The token is null %s: ", jwtToken)
	}

	claims, isOK := parsedToken.Claims.(jwt.MapClaims)
	if !isOK{
		return nil, fmt.Errorf("Claims are empty in %s", jwtToken)
	}

	userName, isOK := claims["name"].(string)
	if !isOK {
		return nil, fmt.Errorf("Can't find  id in %s", jwtToken)
	}

	return &userstruct.User{
		Name: userName,
	}, nil
}

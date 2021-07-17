package dbmanage

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func DBconnection() *gorm.DB{

	if err := godotenv.Load(); err != nil{
		log.Fatalln("Error loading .env file" +  err.Error())

	}

	pass := os.Getenv("MYSQL_ROOT_PASSWORD")

	DBMS := "mysql"
	user := "root"
	proto := "tcp(127.0.0.1:3306)"
	DBname := "CAtech"

	configure := user + ":" + pass + "@" + proto + "/" + DBname
	db, err := gorm.Open(DBMS, configure)
	if  err != nil {
		log.Fatalln(err)
	}

	return db

}

FROM golang:1.16.6

RUN mkdir /app
WORKDIR /app

COPY  ./goAPIs/ /app/

#RUN go mod init app

RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/auth0/go-jwt-middleware
RUN go get -u github.com/form3tech-oss/jwt-go
RUN go get -u github.com/joho/godotenv


EXPOSE 8085

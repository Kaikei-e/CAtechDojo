FROM golang:1.16.6

RUN mkdir /app
WORKDIR /app

COPY  ./goAPIs/ /app/

#RUN go mod init app

RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/jinzhu/gorm


EXPOSE 8085

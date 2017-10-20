FROM golang:onbuild

RUN ["go", "get", "-u", "github.com/jinzhu/gorm"]
RUN ["go", "get", "-u", "github.com/jinzhu/gorm/dialects/postgres"]
RUN ["go", "get", "-u", "github.com/labstack/echo"]

RUN ["apt-get", "update"]

EXPOSE 3000

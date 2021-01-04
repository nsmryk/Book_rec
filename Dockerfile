# Dockerfile
FROM golang:latest

WORKDIR /gin-book

# Install dependency module
RUN go get github.com/go-sql-driver/mysql \
    && go get -u github.com/gin-gonic/gin \
    && go get github.com/gorilla/mux \
    && go get -u github.com/jinzhu/gorm \
    && go get github.com/gin-contrib/cors \
    && go get gopkg.in/ini.v1 \
    && go get github.com/go-xorm/cmd/xorm

ENV PATH $PATH:/gin-book

#CMD ["go run main.go"]

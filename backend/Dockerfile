FROM golang:latest

ADD . /go/src/github.com/soeyusuke/rapsodier/api

RUN go get github.com/gin-gonic/gin
RUN go get github.com/gin-contrib/cors
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/gorm

ENV MYSQL_USER todos_user
ENV MYSQL_PASS todos_pass
ENV MYSQL_DATABASE todos_db

RUN go install github.com/soeyusuke/rapsodier/api

ENTRYPOINT /go/bin/api

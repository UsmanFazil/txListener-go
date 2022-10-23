# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./

RUN apk add build-base

# RUN go mod download 

RUN go get github.com/ethereum/go-ethereum
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/gorm
RUN go get github.com/pkg/errors

COPY . ./

RUN go build -o /godocker

CMD [ "/godocker" ]
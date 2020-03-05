FROM golang:1.14-alpine
LABEL maintainer="teruya ono"
WORKDIR /go/src/app
RUN apk add \
    alpine-sdk \
    git \
    && go get github.com/pilu/fresh \
    && go get github.com/go-sql-driver/mysql
CMD [ "fresh" ]

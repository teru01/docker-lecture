FROM golang:1.14-alpine
LABEL maintainer="teruya ono"
WORKDIR /go/src/app
RUN apk add --no-cache \
    alpine-sdk \
    git \
    && go get github.com/pilu/fresh
CMD [ "fresh" ]

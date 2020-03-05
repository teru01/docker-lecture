FROM golang:1.14-alpine as builder
WORKDIR /go/src/app
COPY ./main.go .
RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/app

FROM alpine:3.11
COPY --from=builder /go/bin/app /go/bin/app
RUN addgroup --system myapp &&\
    adduser --no-create-home --disabled-password --system --ingroup myapp myapp
USER myapp
ENTRYPOINT [ "/go/bin/app" ]



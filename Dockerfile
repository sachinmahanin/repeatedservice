FROM golang:alpine as builder
RUN apk update 
RUN apk add curl


ARG BUILD_VERSION
ARG COMMIT_SHA
ARG BUILD_SHORT_VERSION

ENV GOARCH=amd64
ENV GOOS=linux
ENV CGO_ENABLED=0
ENV APP_VERSION=1
ENV HOSTNAME=localhost
ENV APP_PORT=18606

COPY . /go/src/github.com/sachinmahanin/passwordrepeated
WORKDIR /go/src/github.com/sachinmahanin/passwordrepeated

RUN go build -mod=vendor -o bin/passwordrepeated
RUN ls /go/src/github.com/sachinmahanin/passwordrepeated
RUN chmod +x /go/src/github.com/sachinmahanin/passwordrepeated
EXPOSE 18605/tcp
CMD ./bin/passwordrepeated
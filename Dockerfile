FROM golang:1.13-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

COPY . .

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
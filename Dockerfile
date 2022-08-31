FROM golang:1.18-alpine

WORKDIR /api
RUN apk add --no-cache git make

COPY go.mod go.sum /
RUN go mod download

COPY . /api

RUN go build -o ./api ./main.go
EXPOSE 8080

ENTRYPOINT [ "./api" ]
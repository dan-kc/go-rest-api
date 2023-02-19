FROM golang:1.20 as build

WORKDIR /usr/src/app

COPY . .

RUN go mod tidy

ENV PORT=8080

EXPOSE 8080

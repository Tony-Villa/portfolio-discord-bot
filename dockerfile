FROM golang

WORKDIR /usr/src/app

COPY . .

RUN go mod tidy

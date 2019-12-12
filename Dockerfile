FROM golang:1.13 AS build

WORKDIR /

EXPOSE 31000

COPY ./ /

RUN  go get -u github.com/gorilla/mux && env GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o echo-server .

ENTRYPOINT [ "echo-server" ]
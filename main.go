package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	InitEchoServer()
}

//InitEchoServer - starting the server
func InitEchoServer() {
	var tempServer *http.Server
	r := mux.NewRouter()
	r.HandleFunc("/echo", echoHandler)
	tempServer = &http.Server{
		Addr:    "0.0.0.0:31000",
		Handler: r,
	}
	fmt.Println("Starting Server at port number - 31000 ")
	tempServer.ListenAndServe()
}
func echoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received at echo server ...")
	w.WriteHeader(200)
	fmt.Println("Copying request body in response body ...")
	n, err := io.Copy(w, r.Body)
	if err != nil {
		fmt.Println("Request failed, error while writing data in response ", fmt.Sprintf("%s", err))
	} else {
		fmt.Println("Request Completed, bytes written in response body ", n)
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

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
	fmt.Println("Starting Server at port number - 31000 \n")
	tempServer.ListenAndServe()
}
func echoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[", time.Now(), "]  Request received at echo server ")
	fmt.Println("[", time.Now(), "]  Reading request body")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("[", time.Now(), "]  error reading request body \n", err)
		w.Header().Set("status-code", "515")
		w.Write([]byte("Error " + err.Error()))
		return
	}
	fmt.Println("[", time.Now(), "] number of bytes read from request body ", len(data))
	fmt.Println("[", time.Now(), "] writing data in response body")
	w.Write(data)
	fmt.Println("[", time.Now(), "]  data written in response body\n")
	w.Header().Set("content-type", r.Header.Get("content-type"))
	w.Header().Set("status-code", "200")
	return
}

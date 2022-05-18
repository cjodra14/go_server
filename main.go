package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	httpDirectory = "./static"
	formEndpoint  = "form"
	helloEndpoint = "hello"

	portNumber = "8080"
)

func main() {
	fileServer := http.FileServer(http.Dir(httpDirectory))
	http.Handle("/", fileServer)
	http.HandleFunc("/"+formEndpoint, formHandler)
	http.HandleFunc("/"+helloEndpoint, helloHandler)

	log.Println("Starting server at port ", portNumber)
	if err := http.ListenAndServe(":"+portNumber, nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(responseWrite http.ResponseWriter, request *http.Request) {
	if request.URL.Path != ("/" + helloEndpoint) {
		http.Error(responseWrite, "404 Not found", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(responseWrite, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(responseWrite, "Hello!")
}

func formHandler(responseWrite http.ResponseWriter, request *http.Request) {

}

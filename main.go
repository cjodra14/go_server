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
	if ok := urlPathComprobation(request, responseWrite, ("/" + helloEndpoint)); !ok {
		return
	}

	if ok := isMethodSupported(request, responseWrite, "GET"); !ok {
		return
	}

	fmt.Fprintf(responseWrite, "Hello!")
}

func formHandler(responseWrite http.ResponseWriter, request *http.Request) {
	ok := validateForm(request, responseWrite)
	if !ok {
		return
	}
	fmt.Fprintf(responseWrite, "POST request successful")
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Fprintf(responseWrite, "Name = %s\n", name)
	fmt.Fprintf(responseWrite, "Address = %s\n", address)
}

func validateForm(request *http.Request, responseWrite http.ResponseWriter) bool {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(responseWrite, "ParseForm() err: %v", err)
		return false
	}
	return true
}

func isMethodSupported(request *http.Request, responseWrite http.ResponseWriter, method string) bool {
	if request.Method != method {
		http.Error(responseWrite, "Method is not supported", http.StatusNotFound)
		return false
	}
	return true
}

func urlPathComprobation(request *http.Request, responseWrite http.ResponseWriter, urlPath string) bool {
	if request.URL.Path != (urlPath) {
		http.Error(responseWrite, "404 Not found", http.StatusNotFound)
		return false
	}
	return true
}

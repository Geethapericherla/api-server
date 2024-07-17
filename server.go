package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", handleHelloWorld)
	http.HandleFunc("/health", handleHealth)
	log.Println("Listening and Serving")
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleHelloWorld(writer http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	writeResponse(writer, "Hello World")

}
func handleHealth(writer http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	writeResponse(writer, "OK")
}
func writeResponse(writer http.ResponseWriter, res string) {
	response := []byte(res)
	_, err := writer.Write(response)
	if err != nil {
		fmt.Println(err)
	}
}

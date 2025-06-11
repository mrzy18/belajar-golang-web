package main

import (
	"fmt"
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world"
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	address := ":3000"
	fmt.Printf("Server is running at %s", address)
	err := http.ListenAndServe(address, nil)
	fmt.Println(err.Error())
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("server is running at localhost:3000")
	http.ListenAndServe(":3000", nil)
}

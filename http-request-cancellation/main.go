package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	select {
	case <-r.Context().Done():
		if err := r.Context().Err(); err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "canceled") {
				log.Println("request canceled")
			} else {
				log.Println("unknown error occured", err.Error())
			}
		}
	case <-done:
		log.Println("done")
	}
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":3000", nil)
}

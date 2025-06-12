package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		errorMessage := map[string]string{
			"error": "Method not allowed",
		}
		result, err := json.Marshal(&errorMessage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(result)
		return
	}
	message := map[string]string{
		"message": "Welcome to my server",
	}
	result, err := json.Marshal(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(result)
}

func httpMethod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		message := map[string]string{
			"HTTP Method": "GET",
		}
		json.NewEncoder(w).Encode(&message)
	case http.MethodPost:
		message := map[string]string{
			"HTTP Method": "POST",
		}
		json.NewEncoder(w).Encode(&message)
	case http.MethodPatch:
		message := map[string]string{
			"HTTP Method": "PATCH",
		}
		json.NewEncoder(w).Encode(&message)
	case http.MethodPut:
		message := map[string]string{
			"HTTP Method": "PUT",
		}
		json.NewEncoder(w).Encode(&message)
	case http.MethodDelete:
		message := map[string]string{
			"HTTP Method": "DELETE",
		}
		json.NewEncoder(w).Encode(&message)
	default:
		message := map[string]string{
			"error": "Method not allowed!",
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(&message)
	}
}

func main() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/http-method", httpMethod)

	fmt.Println("Server is running at http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}

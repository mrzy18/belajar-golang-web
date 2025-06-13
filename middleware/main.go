package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/student", ActionStudent)

	var handler http.Handler = mux
	handler = MiddlewareAuth(handler)
	handler = MiddlewareAllowOnlyGet(handler)

	server := new(http.Server)
	server.Addr = ":3000"
	server.Handler = handler

	fmt.Println("Server is running at http://localhost:3000")
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if id := r.URL.Query().Get("id"); id != "" {
		student, err := SelectStudent(id)
		if err != nil {
			http.Error(w, "id not found", http.StatusNotFound)
			return
		}
		OutputJSON(w, student)
		return
	}

	OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, student any) {
	res, err := json.Marshal(student)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

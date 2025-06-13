package main

import (
	"net/http"
)

const USERNAME = "user"
const PASSWORD = "secret"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte("something went wrong"))
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		w.Write([]byte("usersename or password wrong"))
		return false
	}
	return true
}

func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodGet {
		w.Write([]byte("only Method GET is allowed"))
		return false
	}
	return true
}

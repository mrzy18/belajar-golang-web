package main

import (
	"net/http"
)

const USERNAME = "user"
const PASSWORD = "secret"

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("something went wrong"))
			return
		}

		isValid := (username == USERNAME) && (password == PASSWORD)
		if !isValid {
			w.Write([]byte("usersename or password wrong"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func MiddlewareAllowOnlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Write([]byte("only Method GET is allowed"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

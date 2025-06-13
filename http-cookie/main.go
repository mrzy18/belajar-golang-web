package main

import (
	"fmt"
	randMath "math/rand"
	"net/http"
	"time"
)

var cookieName = "cookieData"

func main() {
	http.HandleFunc("/", ActionIndex)
	http.HandleFunc("/delete", ActionDelete)

	fmt.Println("Server is running at http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}

func RandomString(length int) string {
	randMath.New(randMath.NewSource(time.Now().UnixNano()))
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randMath.Intn(len(letters))]
	}

	return string(b)
}

func ActionIndex(w http.ResponseWriter, r *http.Request) {
	cookieName := "cookieData"
	c := &http.Cookie{}
	if storedCookie, _ := r.Cookie(cookieName); storedCookie != nil {
		c = storedCookie
	}

	if c.Value == "" {
		c = &http.Cookie{}
		c.Name = cookieName
		c.Value = RandomString(32)
		c.Expires = time.Now().Add(5 * time.Minute)
		http.SetCookie(w, c)
	}

	w.Write([]byte(c.Value))
}

func ActionDelete(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{}
	c.Name = cookieName
	c.Expires = time.Unix(0, 0)
	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

}

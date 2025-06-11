package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type M map[string]any

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := M{
			"name":    "John Doe",
			"Gender":  "male",
			"Hobbies": []string{"coding", "traveling", "eating"},
		}

		tmpl := template.Must(template.ParseFiles(
			"views/index.html",
			"views/_header.html",
			"views/_message.html",
		))

		err := tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		address := M{
			"Country":  "Utopia",
			"Province": "Tanggaroa",
		}
		person := M{
			"Name":    "John Doe",
			"Gender":  "male",
			"Hobbies": []string{"coding", "traveling", "eating"},
			"Address": address,
		}

		tmpl := template.Must(template.ParseFiles(
			"views/about.html",
			"views/_header.html",
			"views/_message.html",
		))
		err := tmpl.ExecuteTemplate(w, "about", person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Server is running at http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}

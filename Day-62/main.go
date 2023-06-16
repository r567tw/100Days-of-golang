package main

import (
	"html/template"
	"net/http"
)

func main() {
	content := template.Must(template.ParseFiles("./index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{"title": "Hello World!"}
		content.Execute(w, data)
	})

	http.ListenAndServe(":8000", nil)
}

package main

import (
	"net/http"
	"templ-portfolio/internal/templates"
)

func main() {
	fs := http.FileServer(http.Dir("internal"))
	http.Handle("/internal/", http.StripPrefix("/internal/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.Hello("Bob").Render(r.Context(), w)
	})
	http.ListenAndServe("localhost:8080", nil)
}

package main

import (
	"deltawashere/portfolio/views"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	fs := http.FileServer(http.Dir("static"))

	r.Use(middleware.Logger)
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	r.Get("/", templ.Handler(views.Hello("world")).ServeHTTP)
	http.ListenAndServe(":3000", r)
}

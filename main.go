package main

import (
	"deltawashere/portfolio/handlers"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	
	godotenv.Load(".env")
l:= os.Getenv("API_SECRET")
_=l
	r := chi.NewMux()
	fs := http.FileServer(http.Dir("static"))
	
	r.Use(middleware.Logger)

	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	
	r.Get("/", handlers.HomeHandler)
	r.Get("/career", handlers.CareerHandler)
	
	http.ListenAndServe(":3000", r)
}

package handlers

import (
	"deltawashere/portfolio/views"
	"net/http"
)

func CareerHandler(w http.ResponseWriter, r *http.Request){
	views.Career().Render(r.Context(), w)
}
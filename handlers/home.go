package handlers

import (
	"deltawashere/portfolio/models"
	"deltawashere/portfolio/views"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request){
	profile := models.Profile{
		Picture: "",
		Summary: "",
		Name: "",
	}
	c := views.Home(profile)
	c.Render(r.Context(), w)
}
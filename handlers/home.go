package handlers

import (
	"deltawashere/portfolio/models"
	"deltawashere/portfolio/views"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request){
	summary := [6]string{"Howdy, name's delta","Im a full stack software developer","This is my portfolio which it's still an WIP","So.. yeah you can read it if you want to know more 'bout me","I ain't writting everything here lol","bye bye"}


	profile := models.Profile{
		Picture: "",
		Summary: summary,
		Name: "",
	}

	

	c := views.Home(profile)
	c.Render(r.Context(), w)
}
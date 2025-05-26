package handlers

import (
	"deltawashere/portfolio/client"
	models "deltawashere/portfolio/viewmodels"
	"deltawashere/portfolio/views"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	summary := [6]string{"Howdy, name's delta", "Im a full stack software developer", "This is my portfolio which it's still a WIP", "So.. yeah you can read it if you want to know more 'bout me", "I ain't writting everything here lol", "bye bye"}

	client := client.NewClient()

	profileData, err := client.GetProfile(r.Context())

	//todo: do something with the err
	_ = err

	profile := models.Profile{
		Picture: profileData.AvatarURL,
		Summary: summary,
		Name:    profileData.Name,
	}

	c := views.Home(profile)
	c.Render(r.Context(), w)
}

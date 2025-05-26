package handlers

import (
	"deltawashere/portfolio/viewmodels"
	"deltawashere/portfolio/views"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func EventHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		println("err")
	}
	event := getEvent(id)

	c := views.Event(event)
	c.Render(r.Context(), w)
}

func getEvent(id int) viewmodels.Event {
	//open Jsonfile
	dir, _ := os.Getwd()
	jsonFile, err := os.Open(fmt.Sprintf("%s/static/utils/careerevents.json", dir))

	if err != nil {
		fmt.Println(err) // todo error page
	}

	defer jsonFile.Close() //defer is the await of go

	var timeLine [][]viewmodels.Event

	if err := json.NewDecoder(jsonFile).Decode(&timeLine); err != nil {
		print(fmt.Sprintf("Error: %q", err.Error()))
	}

	for i := 0; i < len(timeLine); i++ {
		for j := 0; j < len(timeLine[i]); j++ {
			if timeLine[i][j].Id == id {
				return timeLine[i][j]
			}
		}
	}
	return viewmodels.Event{}
}

package handlers

import (
	"deltawashere/portfolio/viewmodels"
	"deltawashere/portfolio/views"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func CareerHandler(w http.ResponseWriter, r *http.Request) {

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
	//todo implement an algorithm to make the timeline rather than jsut get it alrerady sorted

	career := viewmodels.Career{
		TimeLine: timeLine,
		Earliest: getEarliest(timeLine),
	}

	c := views.Career(career)

	c.Render(r.Context(), w)
}

func getEarliest(timeLine [][]viewmodels.Event) time.Time {
	earliest := time.Now()
	for i := 0; i < len(timeLine); i++ {
		for j := 0; j < len(timeLine[i]); j++ {
			if timeLine[i][j].Start.Time.Compare(earliest) == -1 {
				earliest = timeLine[i][j].Start.Time
			}
		}
	}
	return earliest
}

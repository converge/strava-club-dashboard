package main

import (
	"github.com/converge/strava-club-dashboard/internal/handlers"
	"log"
	"net/http"
)

func main() {
	// using strava user token (account: x@x.com)
	httpClient := http.DefaultClient
	stravaHandler := handlers.StravaHandler{
		HTTPClient: httpClient,
	}
	err := stravaHandler.GetClub()
	if err != nil {
		log.Println(err)
	}
	// request clubs data , todo: what's the club endpoint?
	// stream data to kafka
	// store data / insert if not exist
	// send slack msg
}
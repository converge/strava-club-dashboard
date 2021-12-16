package main

import (
	"github.com/converge/strava-club-dashboard/internal/handlers"
	"github.com/converge/strava-club-dashboard/internal/services"
	"log"
	"net/http"
)

func main() {
	// using strava user token (account: x@x.com)
	httpClient := http.DefaultClient
	//kafkaClient := kafka
	stravaService := services.NewStrava()
	stravaHandler := handlers.NewStrava(httpClient, stravaService)

	// request clubs data
	// stream data to kafka
	err := stravaHandler.GetClubUpdate()
	if err != nil {
		log.Println(err)
	}

	// store data / insert if not exist
	// send slack msg
}

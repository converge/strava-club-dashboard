package main

import (
	"github.com/converge/strava-club-dashboard/internal/handlers"
	"github.com/converge/strava-club-dashboard/internal/services"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	httpClient := http.DefaultClient

	stravaService := services.StravaService{}
	stravaHandler := handlers.StravaHandler{
		HTTPClient: httpClient,
		Service:    stravaService,
	}
	log.Info().Msg("things are running...")

	err := stravaHandler.CheckNewActivities()
	if err != nil {
		log.Err(err)
	}
}

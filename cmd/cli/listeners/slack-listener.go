package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/converge/strava-club-dashboard/internal/handlers"
	"github.com/converge/strava-club-dashboard/internal/services"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

func main() {
	httpClient := http.DefaultClient

	conf := kafka.ConfigMap{}
	bootstrapServer := os.Getenv("BOOTSTRAP_SERVER")
	conf["group.id"] = "kafka-go-getting-started"
	conf["auto.offset.reset"] = "earliest"
	conf["bootstrap.servers"] = bootstrapServer

	kafkaProducer, err := kafka.NewProducer(&conf)
	if err != nil {
		log.Err(err)
		panic(err)
	}
	defer kafkaProducer.Close()


	kafkaConsumer, err := kafka.NewConsumer(&conf)
	if err != nil {
		log.Info().Msg("failed to create consumer")
		panic(err)
	}
	stravaService := services.NewStrava(kafkaProducer, kafkaConsumer)
	stravaHandler := handlers.StravaHandler{
		HTTPClient: httpClient,
		Service:    stravaService,
	}
	log.Info().Msg("things are running...")

	err = stravaHandler.CheckNewActivities()
	if err != nil {
		log.Err(err)
	}
}

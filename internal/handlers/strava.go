package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/converge/strava-club-dashboard/internal/services"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"os"
)

type HTTPClientDoer interface {
	Do(req *http.Request) (*http.Response, error)
}
type StravaHandler struct {
	HTTPClient HTTPClientDoer
	Service    services.StravaService
}

func NewStrava(httpClient *http.Client, stravaService services.StravaService) StravaHandler {
	return StravaHandler{
		HTTPClient: httpClient,
		Service:    stravaService,
	}
}

// GetClubUpdate
// todo: move env. values to a global config object
func (handler StravaHandler) GetClubUpdate() error {

	clubId := os.Getenv("CLUB_ID")
	stravaToken := os.Getenv("STRAVA_TOKEN")

	stravaApiBase := "https://www.strava.com/api/v3"
	stravaGetClubActivities := fmt.Sprintf("clubs/%s/activities?page=1&per_page=10", clubId)
	fullUrl := fmt.Sprintf("%s/%s", stravaApiBase, stravaGetClubActivities)

	authToken := fmt.Sprintf("%s%s", "Bearer ", stravaToken)
	req, err := http.NewRequest(http.MethodGet, fullUrl, nil)
	req.Header.Add("Authorization", authToken)
	if err != nil {
		log.Err(err)
		return err
	}

	res, err := handler.HTTPClient.Do(req)
	if err != nil {
		log.Err(err)
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Err(err)
		return err
	}
	defer res.Body.Close()

	fmt.Println(string(body))

	var rca []services.RiderClubActivity
	err = json.Unmarshal(body, &rca)
	if err != nil {
		log.Err(err)
		return err
	}
	for key, val := range rca {
		log.Printf("sending message id: %d to producer\n", key)
		log.Info().Msgf("val: %v", val)

		err = handler.Service.ProduceKafkaMessage(val)
		if err != nil {
			log.Err(err)
			return err
		}
	}

	return nil
}

func (handler StravaHandler) CheckNewActivities() error {

	err := handler.Service.GetKafkaMessage()
	if err != nil {
		log.Err(err)
		return err
	}
	return nil
}

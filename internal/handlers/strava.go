package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type HTTPClientDoer interface {
	Do(req *http.Request) (*http.Response, error)
}
type StravaHandler struct {
	HTTPClient HTTPClientDoer
}

func NewStrava(httpClient *http.Client) StravaHandler {
	return StravaHandler{
		HTTPClient: httpClient,
	}
}

// GetClub
// todo: move env. values to a global config object
func (service StravaHandler) GetClub () error {

	clubId := os.Getenv("CLUB_ID")
	stravaToken := os.Getenv("STRAVA_TOKEN")

	stravaApiBase := "https://www.strava.com/api/v3"
	stravaGetClubActivities := fmt.Sprintf("clubs/%s/activities?page=1&per_page=10", clubId)
	fullUrl := fmt.Sprintf("%s/%s", stravaApiBase, stravaGetClubActivities)
	fmt.Println(fullUrl)

	authToken := fmt.Sprintf("%s%s", "Bearer ", stravaToken)
	req, err := http.NewRequest(http.MethodGet, fullUrl, nil)
	req.Header.Add("Authorization", authToken)
	if err != nil {
		log.Println(err)
		return err
	}

	res, err := service.HTTPClient.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	defer res.Body.Close()

	fmt.Println(string(body))

	return nil
}


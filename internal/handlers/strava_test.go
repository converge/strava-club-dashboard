package handlers

import (
	"log"
	"net/http"
	"testing"
)

func (mockedHTTPClient MockedHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return nil, nil
}

type MockedHTTPClient struct{}

func TestStravaHandler_GetClub(t *testing.T) {
	mockedHTTPClient := MockedHTTPClient{}
	stravaHandler := StravaHandler{
		HTTPClient: mockedHTTPClient,
	}
	err := stravaHandler.GetClub()
	if err != nil {
		log.Println(err)
	}

}

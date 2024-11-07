package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

type LocationAreasResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endPoint := "/location-area"
	fullURL := baseURL + endPoint

	// req, err := http.NewRequest("GET", fullURL, nil)
	// if err != nil {
	// 	return LocationAreasResp{}, err
	// }

	req, err := http.Get(fullURL)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// req.Body.Close()

	if req.StatusCode > 299 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %d and\nbody: %s", req.StatusCode, body)
	}

	locationArea := LocationAreasResp{}
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationArea, nil
}

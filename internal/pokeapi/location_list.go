package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// getting the list of locations
func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	fullURL := baseURL + "/location-area"
	if pageURL != nil {
		fullURL = *pageURL
	}

	response, err := http.Get(fullURL)
	if err != nil {
		return LocationAreasResp{}, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	if response.StatusCode > 299 {
		return LocationAreasResp{}, fmt.Errorf("failed with bad status code: %d", response.StatusCode)
	}

	locationArea := LocationAreasResp{}

	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationArea, nil
}

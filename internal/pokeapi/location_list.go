package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Getting the list of locations
func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	locationArea := LocationAreasResp{}

	// check if the location exists in the cache
	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &locationArea)
		if err != nil {
			return LocationAreasResp{}, err
		}
		return locationArea, nil
	}

	// else get the location and add it to the cache
	req, err := http.Get(url)
	if err != nil {
		return locationArea, err
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return locationArea, err
	}

	if req.StatusCode > 299 {
		return locationArea, fmt.Errorf("failed with bad status code: %d", req.StatusCode)
	}

	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(url, body)

	return locationArea, nil
}

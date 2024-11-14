package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// Get location information
func (c *Client) GetLocation(name string) (Location, error) {
	url := baseURL + "/location-area/" + name

	locationResp := Location{}

	// check into cache
	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Location{}, err
		}
		return locationResp, nil
	}

	// else get information and add to the cache
	req, err := http.Get(url)
	if err != nil {
		return locationResp, err
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return locationResp, err
	}

	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, body)

	return locationResp, nil
}

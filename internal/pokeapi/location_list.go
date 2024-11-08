package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	fullURL := baseURL + "/location-area"
	if pageURL != nil {
		fullURL = *pageURL
	}

	// req, err := http.NewRequest("GET", fullURL, nil)
	// if err != nil {
	// 	return LocationAreasResp{}, err
	// }

	req, err := http.Get(fullURL)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// defer req.Body.Close()

	// body, err := io.ReadAll(req.Body)
	// if err != nil {
	// 	return LocationAreasResp{}, err
	// }

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// req.Body.Close()

	// if req.StatusCode > 299 {
	// 	return LocationAreasResp{}, fmt.Errorf("bad status code: %d and\nbody: %s", req.StatusCode, body)
	// }

	locationArea := LocationAreasResp{}
	// err = json.Unmarshal(body, &locationArea)
	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationArea, nil
}

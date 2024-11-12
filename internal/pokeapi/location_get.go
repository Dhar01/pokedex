package pokeapi

func (c *Client) GetLocation(name string) (Location, error) {
	url := baseURL + "/location-area/" + name

	locationResp := Location{}

	return locationResp, nil
}

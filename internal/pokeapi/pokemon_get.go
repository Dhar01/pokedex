package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// Get pokemon information
func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	pokemonResp := Pokemon{}

	// check into cache
	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}

	// else get the information and adding to the cache
	req, err := http.Get(url)
	if err != nil {
		return pokemonResp, err
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return pokemonResp, err
	}

	err = json.Unmarshal(body, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, body)

	return pokemonResp, nil
}

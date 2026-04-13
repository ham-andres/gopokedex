package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LocationGet(area string) (LocatePokemon, error)  {
	url := baseURL + "/location-area/"	+ area
	
	if val, ok := c.clientCache.Get(url); ok {
		PokemonList := LocatePokemon{}
		
		err := json.Unmarshal(val, &PokemonList) 
		if err != nil {
			return LocatePokemon{}, err
		}
		return PokemonList, nil

	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocatePokemon{}, err
	}

	responses, err := c.httpClient.Do(req)
	if err != nil {
		return LocatePokemon{}, err
	}
	defer responses.Body.Close()

	dat, err := io.ReadAll(responses.Body)
	if err != nil {
		return LocatePokemon{}, err
	}

	PokemonList := LocatePokemon{}
	err = json.Unmarshal(dat, &PokemonList)
	if err != nil {
		return LocatePokemon{}, err
	}

	c.clientCache.Add(url, dat)

	return PokemonList, nil
}

package pokeapi

import (
	"encoding/json"
	"net/http"
	"io"
)

func (c * Client) GetPokemon( name string) (PokemonInfo, error) {
	url := baseURL + "/pokemon/" + name

	if val, ok := c.clientCache.Get(url); ok {
		PokemonDetail := PokemonInfo{}

		err := json.Unmarshal(val, &PokemonDetail)
		if err != nil {
			return PokemonInfo{}, err
		}
		return PokemonDetail,nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInfo{}, err
	}

	responses, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfo{}, err
	}
	defer responses.Body.Close()

	dat, err := io.ReadAll(responses.Body)
	if err != nil {
		return PokemonInfo{}, err
	}

	PokemonDetail := PokemonInfo{}
	err = json.Unmarshal(dat, &PokemonDetail)
	if err != nil {
		return PokemonInfo{}, err
	}

	c.clientCache.Add(url, dat)

	return PokemonDetail,nil
}


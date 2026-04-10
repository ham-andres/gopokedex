package pokeapi

import (
		"encoding/json"
		"io"
		"net/http"
		
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	c_req, exists := 	c.clientCache.Get(url)
	if exists {

			locationsResp := RespShallowLocations{}
			err := json.Unmarshal(c_req, &locationsResp)
			if err != nil {
				return RespShallowLocations{}, err
			}
			return locationsResp, nil

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}
	



	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.clientCache.Add(url, dat)

	return locationsResp, nil
}

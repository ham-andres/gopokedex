package main 

import (
				"errors"
				"fmt"
)


// this function is for map forward 
func commandMapf(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.NextLocationsURL)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationsResp.Next
	cfg.PrevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

		//
		// url := "https://pokeapi.co/api/v2/location-area/"
		// if cfg.Next != nil {
		// 		url = *cfg.Next
		// } 
		// res, err := http.Get(url)
		//
		// if err != nil {
		// 	return err	
		// }
		// defer res.Body.Close()
		//
		// body, err := io.ReadAll(res.Body)
		// if err != nil {
		// 	return err
		// }
		//
		// var json_res commandMap
		// if err := json.Unmarshal(body, &json_res); err!= nil {
		// 	return err
		// }
		// cfg.Next = json_res.Next
		// cfg.Previous = json_res.Previous
		// for _, j_res := range json_res.Results {
		// 		fmt.Println(j_res.Name)
		//
		// }
		// return nil


// this function is for map backward 

func commandMapb(cfg *config) error {
	if cfg.PrevLocationsURL == nil {
		return errors.New("you're on the first page")
	}
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.PrevLocationsURL)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationResp.Next
	cfg.PrevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
} 


// {
// 	url := "https://pokeapi.co/api/v2/location-area/"
// 	if cfg.Previous == nil {
// 		return errors.New("you're on the first page")
//
// 	}
//
// 	if cfg.Previous != nil {
// 		url = *cfg.Previous
// 	}
// 	res, err := http.Get(url)
//
// 	if err != nil {
// 		return err
// 	}
// 	defer res.Body.Close()
//
// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		return err
// 	}
//
// 	var json_res_back commandMap
// 	if err := json.Unmarshal(body, &json_res_back); err != nil {
// 		return err
// 	}
// 	cfg.Previous = json_res_back.Previous
// 	cfg.Next = json_res_back.Next
// 	for _, j_res_back := range json_res_back.Results{
// 		fmt.Println(j_res_back.Name)
// 	}
// 	return nil
// }

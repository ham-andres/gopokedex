package pokeapi

// pokemon attribute 

type PokemonStat struct {
	BaseStat int      `json:"base_stat"`
	Effort   int      `json:"effort"`
	Stat     StatInfo `json:"stat"`
}

type StatInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Types struct {
	Slot	int		`json:"slot"`
	Type	Types_type 	`json:"type"`
}

type Types_type struct {
	Name 	string	`json:"name"`
	URL		string	`json:"url"`
}


type PokemonInfo struct {
	Name           string        `json:"name"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	BaseExperience int           `json:"base_experience"`
	Types					 []Types 			 `json:"types"`
}



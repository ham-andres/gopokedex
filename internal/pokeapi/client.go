package pokeapi 

import (
		"net/http"
		"time"
		"github.com/ham-andres/gopokedex/internal/pokecache"
)


// Client
type Client struct {
	httpClient http.Client
	clientCache *pokecache.Cache  
}

// Newclient 
func NewClient( timeout, cacheinterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		
		},
		clientCache: pokecache.NewCache(cacheinterval),
	}
}

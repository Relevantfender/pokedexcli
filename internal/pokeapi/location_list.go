package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Relevantfender/pokedexcli/internal/pokecache"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string, cache *pokecache.Cache) (RespShallowLocations, error) {
	// get the url
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	var data []byte

	locationsResp := RespShallowLocations{}

	// check for url in cache
	cachedResponse, ok := cache.Get(url)

	if ok {
		data = cachedResponse

	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespShallowLocations{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespShallowLocations{}, err
		}

		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespShallowLocations{}, err
		}
		cache.Add(url, data)
	}

	err := json.Unmarshal(data, &locationsResp)

	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}

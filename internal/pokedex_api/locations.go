package pokedex_api

import (
	"encoding/json"
	"io"
	"net/http"
)

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// this can be done via anononymous struct
type LocationsResult struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

func (c *Client) ListLocations(pageURL *string) (LocationsResult, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationsResult{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationsResult{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationsResult{}, err
	}

	locationsResp := LocationsResult{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationsResult{}, err
	}

	return locationsResp, nil
}

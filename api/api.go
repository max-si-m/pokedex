package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// this can be done via anononymous struct
type Result struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []Location `json:"results"`
}

const baseURL string = "https://pokeapi.co/api/v2/"

func GetLocations() {
	fullURL := baseURL + "location-area/"
	res, err := http.Get(fullURL)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	var result Result
    err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("error happened during parsing %s\n", err)
	}

	fmt.Printf("%+v", result)
}

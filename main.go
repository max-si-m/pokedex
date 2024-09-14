package main

import (
	"time"

	"github.com/max-si-m/pokedex/internal/pokedex_api"
	"github.com/max-si-m/pokedex/repl"
)

func main(){
	pokeClient := pokedex_api.NewClient(5 * time.Second)
	cfg := &repl.Config{
		PokeApiClient: &pokeClient,
	}

	repl.Start(cfg)
}

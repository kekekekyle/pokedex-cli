package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemon string) (PokemonResponse, error) {
	url := baseURL + "/pokemon/" + pokemon

  val, ok := c.cache.Get(url)
  if ok == true {
    pokemonResp := PokemonResponse{}
    err := json.Unmarshal(val, &pokemonResp)
    if err != nil {
      return PokemonResponse{}, err
    }
    return pokemonResp, nil
  }

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	pokemonResp := PokemonResponse{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return PokemonResponse{}, err
	}

  c.cache.Add(url, dat)

	return pokemonResp, nil
}


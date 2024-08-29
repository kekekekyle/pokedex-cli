package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

  val, ok := c.cache.Get(url)
  if ok == true {
    locationsResp := LocationResponse{}
    err := json.Unmarshal(val, &locationsResp)
    if err != nil {
      return LocationResponse{}, err
    }
    return locationsResp, nil
  }

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationResponse{}, err
	}

	locationsResp := LocationResponse{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationResponse{}, err
	}

  c.cache.Add(url, dat)

	return locationsResp, nil
}


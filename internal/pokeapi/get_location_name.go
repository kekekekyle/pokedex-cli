package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationName(location string) (LocationNameResponse, error) {
	url := baseURL + "/location-area/" + location

  val, ok := c.cache.Get(url)
  if ok == true {
    locationNameResp := LocationNameResponse{}
    err := json.Unmarshal(val, &locationNameResp)
    if err != nil {
      return LocationNameResponse{}, err
    }
    return locationNameResp, nil
  }

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationNameResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationNameResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationNameResponse{}, err
	}

	locationNameResp := LocationNameResponse{}
	err = json.Unmarshal(dat, &locationNameResp)
	if err != nil {
		return LocationNameResponse{}, err
	}

  c.cache.Add(url, dat)

	return locationNameResp, nil
}


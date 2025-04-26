package pokeapi

import (
	"encoding/json"
	)
	
func (c *Client) GetLocationArea(name string) (LocationArea, error) {
	url := baseURL + "/location-area/" + name
	
	dat, err := c.getFromApi(url)
	if err != nil {
		return LocationArea{}, err
	}
	
	locationArea := LocationArea{}
	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	
	return locationArea, nil
	
}
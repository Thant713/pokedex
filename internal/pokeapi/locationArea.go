package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationAreasResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func FetchLocationAreas(url string) (LocationAreasResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return LocationAreasResponse{}, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var areas LocationAreasResponse
	if err := json.NewDecoder(resp.Body).Decode(&areas); err != nil {
		return LocationAreasResponse{}, err
	}

	defer resp.Body.Close()
	return areas, nil
}

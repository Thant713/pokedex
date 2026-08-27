package main

import (
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

	// decode

	defer resp.Body.Close()
}

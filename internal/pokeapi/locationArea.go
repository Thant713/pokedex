package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchLocationAreas(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	// decode and write structs for decoding

	defer resp.Body.Close()








}


package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getHolidays(url string) ([]Holiday, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	// Create a nil slice of items []Item
	var holidays []Holiday
	// Create a new *json.Decoder using json.NewDecoder
	decoder := json.NewDecoder(res.Body)
	// Decode the response body using the decoder's Decode method
	if err := decoder.Decode(&holidays); err != nil {
		return nil, fmt.Errorf("error decoding response body")
	}
	// Return the slice of items
	return holidays, nil
}

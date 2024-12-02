package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getHolidays(url string) ([]Holiday, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	// ?
	// Get the data from the response body using io.ReadAll,
	// creating a slice of bytes []byte.
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	// Create a nil slice of items []Item
	var holidays []Holiday
	// Use json.Unmarshal on the data to get the JSON data
	if err = json.Unmarshal(data, &holidays); err != nil {
		return nil, err
	}
	// Return the items
	return holidays, nil
}

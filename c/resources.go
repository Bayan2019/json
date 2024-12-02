package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

func getResources(url string) ([]map[string]any, error) {
	var resources []map[string]any

	res, err := http.Get(url)
	if err != nil {
		return resources, err
	}

	defer res.Body.Close()

	// Decode the response body into a slice of maps []map[string]interface{} and return it.
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&resources)
	if err != nil {
		return resources, err
	}

	return resources, nil
}

func logResources(resources []map[string]any) {
	// any is an alias for interface{}
	var formattedStrings []string

	// Iterate over the slice of map[string]interface{}
	for _, resource := range resources {
		for key, value := range resource {
			// For each map[string]interface{} get its keys and values
			// using range and append it to formattedStrings as
			// Key: %s - Value: %v, where %s is the key and %v is the value.
			formattedStrings = append(formattedStrings, fmt.Sprintf("Key: %s - Value: %v", key, value))
		}
	}

	sort.Strings(formattedStrings)

	for _, str := range formattedStrings {
		fmt.Println(str)
	}
}

package main

import (
	"encoding/json"
	"log"
)

func marshalAll[T any](items []T) ([][]byte, error) {
	var dataItems [][]byte
	for _, item := range items {
		data, err := json.Marshal(item)
		if err != nil {
			log.Fatal(err)
		}
		dataItems = append(dataItems, data)
	}

	return dataItems, nil
}

package main

import (
	"fmt"
	"log"
)

func main() {
	holidays, err := getHolidays("https://date.nager.at/api/v3/PublicHolidays/2024/kz")
	if err != nil {
		log.Fatal(err)
	}

	for _, h := range holidays {
		fmt.Printf("{date: %s, name: %s, local-name: %s}\n", h.Date, h.Name, h.LocalName)
	}
}

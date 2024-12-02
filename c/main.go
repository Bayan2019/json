package main

import (
	"fmt"
)

const jokesUrl = "https://api.sampleapis.com/jokes/goodJokes"
const bitcoinUrl = "https://api.sampleapis.com/bitcoin/historical_prices"
const healthUrl = "https://api.sampleapis.com/health/professions"

func main() {
	jokes, err := getResources(jokesUrl)
	if err != nil {
		fmt.Println("Error getting jokes:", err)
		return
	}
	fmt.Println("Jokes:")
	logResources(jokes)
	fmt.Println("---")

	bitcoinPrices, err := getResources(bitcoinUrl)
	if err != nil {
		fmt.Println("Error getting bitcoinPrices:", err)
		return
	}
	fmt.Println("BitCoin:")
	logResources(bitcoinPrices)
	fmt.Println("---")

	healths, err := getResources(healthUrl)
	if err != nil {
		fmt.Println("Error getting Health Professionals:", err)
		return
	}
	fmt.Println("Health Professionals:")
	logResources(healths)

}

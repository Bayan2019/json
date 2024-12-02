package main

type Holiday struct {
	Date        string   `json:"date"`
	LocalName   string   `json:"localName"`
	Name        string   `json:"name"`
	CountryCode string   `json:"countryCode"`
	Fixed       bool     `json:"fixed"`
	Global      bool     `json:"global"`
	Counties    []string `json:"counties"`
	LaunchYear  int      `json:"launchYear"`
	Types       []string `json:"types"`
}

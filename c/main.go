package main

import (
	"fmt"
)

const cyclistUrl = "https://raw.githubusercontent.com/freeCodeCamp/ProjectReferenceData/master/cyclist-data.json"
const educationUrl = "https://cdn.freecodecamp.org/testable-projects-fcc/data/choropleth_map/for_user_education.json"
const booksUrl = "https://raw.githubusercontent.com/benoitvallon/100-best-books/refs/heads/master/books.json"

func main() {
	// cyclists, err := getResources(cyclistUrl)
	// if err != nil {
	// 	fmt.Println("Error getting locations:", err)
	// 	return
	// }
	// fmt.Println("Resources:")
	// logResources(cyclists)
	// fmt.Println("---")

	// educations, err := getResources(educationUrl)
	// if err != nil {
	// 	fmt.Println("Error getting items:", err)
	// 	return
	// }
	// fmt.Println("Educations:")
	// logResources(educations)
	// fmt.Println("---")

	books, err := getResources(booksUrl)
	if err != nil {
		fmt.Println("Error getting items:", err)
		return
	}
	fmt.Println("Book:")
	logResources(books)

}

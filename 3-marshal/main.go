package main

import (
	"fmt"
	"log"
)

// Example structs for testing
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {

	types := []any{
		User{"001", "Sir Bedevere the Wise", 25},
		User{"002", "Sir Lancelot the Brave", 30},
		Product{"The Grail", 19.99},
		User{"003", "Sir Galahad the Pure", 19},
		Product{"The Holy Hand Grenade of Antioch", 4.89},
		Product{"The Wooden Rabbit", 1.99},
	}

	bytes, err := marshalAll(types)
	if err != nil {
		log.Fatal(err)
	}
	for _, b := range bytes {
		fmt.Println(string(b))
	}

}

package main

import (
	"encoding/json"
	"log"
	"fmt"
)

//Movie .
type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
	Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
	Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
	Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	fmt.Printf("%s\n", a(movies))
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Printf("%s\n", a1(movies))	
}

func a(s []Movie) []byte {
	data, err := json.Marshal(s)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	return data
}

func a1(s []Movie) []byte {
	data, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	return data
}
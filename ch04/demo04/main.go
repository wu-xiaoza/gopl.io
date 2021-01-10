package main

import "fmt"

type address struct {
    hostname string
    port     int
}

func main() {
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
	for k, v := range hits {
		fmt.Printf("%v\t%v\n", k, v)
	}

	fmt.Println(hits[address{"golang.org", 443}])
}
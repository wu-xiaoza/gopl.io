package main

import "fmt"

func main() {
	a := []string{"qwq", "qdsfg", "gghh", "ghjj"}
	c := a[:]

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", k(c))
}

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}
package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
}

func nonempty(string []string) []string {
	i := 0
	for _, s := range string {
		if s != "" {
			string[i] = s
			i++
		}
	}
	return string[:i]
}

func nonempty2(string []string) []string {
	out := string[:0]
	for _, s := range string {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

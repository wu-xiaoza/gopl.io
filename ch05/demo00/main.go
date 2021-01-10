package main

import (
	"fmt"	
	"strings"
)


func main() {
	// Map returns a copy of the string s with all its characters modified according to the mapping function.
	// If mapping returns a negative value, the character is dropped from the string with no replacement.

	fmt.Println(strings.Map(add1, "VMS"))
}

func add1(r rune) rune { 
	return r+1
}
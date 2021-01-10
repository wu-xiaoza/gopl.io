package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		fmt.Println(arg)
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

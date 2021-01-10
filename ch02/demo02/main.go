package main

import (
	"fmt"
	"os"
)

func main() {
	a, _ := os.Getwd()

	fmt.Println(a)
}

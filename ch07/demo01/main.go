package main

import "fmt"

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d)}

func main() {
	var a dollars = 50
	fmt.Println(a)

	var b dollars = 5
	fmt.Println(b.String())
}
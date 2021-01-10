package main

import "fmt"

func main() {
	var i int = 3
	f(&i)
	fmt.Println(i)
}

func f(a *int) {
	*a++
}

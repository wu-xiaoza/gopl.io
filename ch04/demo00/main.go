package main

import "fmt"

func main() {
	var a [32]byte = [32]byte{'a', 'b', 'c'}
	fmt.Printf("%v\n", a)
	zero(&a)
	fmt.Printf("%v\n", a)

}

func zero(ptr *[32]byte) {
	// for i := range ptr {
	// 	ptr[i] = 0
	// }

	*ptr = [32]byte{}

}
package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := a[3:6]
	c := b[:5]
	fmt.Printf("%v\n", len(a))
	fmt.Printf("%v\n", cap(a))
	fmt.Printf("%v\n", len(b))
	fmt.Printf("%v\n", cap(b))
	fmt.Printf("%v\n", len(c))
	fmt.Printf("%v\n", cap(c))
	fmt.Printf("%v\n%v\n%v\n", a, b, c)
	reverse(c)
	fmt.Printf("%v\n", c)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j =  i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
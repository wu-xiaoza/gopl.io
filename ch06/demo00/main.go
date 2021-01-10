package main

import (
	"fmt"

	"gopl.io/ch6/geometry"
)

func main() {
	q := geometry.Point{1, 2}
	p := geometry.Point{4, 6}

	fmt.Println(geometry.Distance(q, p))
	fmt.Println(p.Distance(q))

	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	//fmt.Println(geometry.PathDistance(perim))
	fmt.Println(perim.Distance())
}

package main

import "fmt"

//Point .
type Point struct{ X, Y int }

//Circle .
type Circle struct {
	Point
	Radius int
}

//Wheel .
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{Circle{Point{8, 8}, 5}, 20}

	var a Wheel
	a = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}

	fmt.Printf("%#v\n", w)
	fmt.Printf("%#v\n", a)

	w.X = 42
	a.Radius = 66

	fmt.Printf("%#v\n", w)
	fmt.Printf("%#v\n", a)
}

package geometry

import (
	"math"
)

//Point .
type Point struct{ x, y float64 }

//Distance .
func Distance(p, q Point) float64 {
	// Hypot returns Sqrt(p*p + q*q), taking care to avoid unnecessary overflow and underflow.
	return math.Hypot(q.x-p.x, q.y-p.y)
}

//Distance .
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

//Path .
type Path []Point

//Distance .
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

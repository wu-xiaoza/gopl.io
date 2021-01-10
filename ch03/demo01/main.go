package main

import "fmt"

//Currency .
type Currency int

//
const (
	USD Currency = iota
	B
	C
	D
)

func main() {
var a = [...]string{USD: "$", B: "￥", C: "€", D: "#"}
fmt.Println(a[B])
}
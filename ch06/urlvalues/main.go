package main

import (
	"fmt"
	"net/url")


func main() {
	m := url.Values{"lang": {"en"}}
	m.Add("itme", "1")
	m.Add("itme", "2")

	fmt.Println(m)
	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("itme"))
	fmt.Println(m["itme"])
	
	m = nil
	fmt.Println(m)
	fmt.Println(m.Get("itme"))
	// m.Add("itme", "3") //assignment to entry in nil map
}
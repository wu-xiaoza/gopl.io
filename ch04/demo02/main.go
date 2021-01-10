package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["alice"] = 1
	ages["charlie"] = 2
	if age, ok := ages["bob"]; !ok {
		fmt.Println(age)
		fmt.Println(!ok)
		} else {
		fmt.Println(age)
		fmt.Println(ok)

	}
	fmt.Println(ages["bob"])
	fmt.Println(true || false)
}
package main

import "fmt"

func main() {
	i := 2
	l := &i         //指针
	fmt.Println(l)  //0xc0000180b0
	fmt.Println(*l) //2
	*l = 5
	fmt.Println(i) //5
}

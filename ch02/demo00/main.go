package main

import (
	"demo/ch2/tempconv"
	"fmt"
)

func main() {

	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)

	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	fmt.Println(tempconv.CToF(tempconv.AbsoluteZeroC))
}

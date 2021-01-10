//ftoc输出两个华氏度 - 摄氏度的转换
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g℉ = %g℃\n", freezingF, fToC(freezingF))
	fmt.Printf("%g℉ = %g℃\n", boilingF, fToC(boilingF))

	fmt.Println(freezingF)
}

func fToC(f float64) float64 {
	//华氏度转摄氏度
	return (f - 32) * 5 / 9
}

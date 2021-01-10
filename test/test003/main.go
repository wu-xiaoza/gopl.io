//题目：一个整数，它加上 100 后是一个完全平方数，再加上 168 又是一个完全平方数，请问该数是多少？
package main

import (
	"fmt"	
	"math"
)


func main() {
	sum := 0
	for {
		x := int(math.Sqrt(float64(sum+100)))
		y := int(math.Sqrt(float64(sum+100+168)))

		if x*x == sum+100 && y*y == sum+100+168 {
			fmt.Printf("答案是：%d\n", sum)
			break
		}
		sum++
	}
}
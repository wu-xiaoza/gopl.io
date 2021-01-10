//题目：判断 101-200 之间有多少个素数，并输出所有素数。
package main

import (
	"fmt"
	"math"
)

func main() {
	var j, k, count int = 0, 0, 0
	for i := 101; i <= 200; i++ {
		k = int(math.Sqrt(float64(i)))
		for j = 2; j <= k; j++ {
			if i%j == 0 {
				break
			}
		}
		fmt.Println(j)
		if j == k+1 {
			fmt.Println(j)
			count++
		}
	}
}

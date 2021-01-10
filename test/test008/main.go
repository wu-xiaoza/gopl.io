// 输出 9*9 乘法口诀表
package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j<= i; j++ {
			fmt.Printf("%d * %d = %d  ", j, i, (i*j))
		}
		fmt.Println()
	}
}
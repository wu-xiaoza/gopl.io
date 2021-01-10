//有 1、2、3、4 这四个数字，能组成多少个互不相同且无重复数字的三位数？都是多少？
package main

import "fmt"

func main() {
	a := 0
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			for k := 1; k < 5; k++ {
				if i != j && j != k && i != k {
					a++
					fmt.Printf("第%d种可能\t%d%d%d\n", a, i, j, k)
				}
			}
		}
	}
	fmt.Println("共有", a, "种可能")
}

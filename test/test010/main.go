// 题目：打印楼梯，同时在楼梯上方打印笑脸。
package main

import "fmt"

func main() {
	for i := 0; i <= 11; i++ {
		if i == 0 {
			fmt.Print("☺\n")
			continue
		} else {
			for j := 0; j <= i-1; j++ {
				fmt.Printf("■ ")
			}
		}
		fmt.Print("☺\n")
	}
}
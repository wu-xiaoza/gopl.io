// 题目：古典问题：有一对兔子，从出生后第 3 个月起每个月都生一对兔子，
// 小兔子长到第三个月后每个月又生一对兔子，假如兔子都不死，问每个月的兔子总数为多少？

package main

import "fmt"

func main() {
	var month int
	fmt.Scanf("%d\n", &month)
	fmt.Printf("第%d个月有%d只兔子\n", month, fun(month))
}

func fun(month int) int {
	if month == 1 || month ==2 {
		return 2
	}
	return (fun(month-1) + fun(month-2))
}
//输入三个整数x，y，z，并把这三个整数由小到大输出
package main

import "fmt"

func main() {
	fmt.Print("请输入三个整数：")
	var x, y, z int = 0, 0, 0
	fmt.Scanf("%d%d%d\n", &x, &y, &z)//不加\n第二次出问题

	if x >= y {
		x, y = y, x
	}
	if x >= z {
		x, z = z, x
	}
	if y >= z {
		y, z = z, y
	}

	fmt.Printf("%d %d %d\n", x, y, z)
}
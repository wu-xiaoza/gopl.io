//打印出所有的 “水仙花数”，所谓 “水仙花数” 是指一个三位数，其各位数字立方和等于该数本身。例如：153 是一个 “水仙花数”，
//因为 153=1 的三次方＋5 的三次方＋3 的三次方。
package main

import "fmt"

func main() {
	for i :=100; i<1000; i++ {
		a := i/100
		b := (i-a*100)/10
		c := (i-a*100-b*10)
		if a*a*a+b*b*b+c*c*c == i {
			fmt.Printf("%d^3+%d^3+%d^3=%d\n", a, b, c, i)
		}
	}
}
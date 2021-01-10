package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// for i := 1; i < len(os.Args); i++ {
	// 	fmt.Printf("  %s\n", comma(os.Args[i]))
	// }

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(comma(input.Text()))
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

package main

import (
	"fmt"
	"time")


func itoa(buf *[]byte, i int, wid int) {
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}

func main() {
	var t time.Time = time.Now()
	var buf []byte
	year, month, day := t.Date()
	fmt.Println(year)
	fmt.Println(month)
	fmt.Println(day)
	
	itoa(&buf, year, 4)
	fmt.Println(year)
	fmt.Println(buf)
	fmt.Printf("%v", buf[0])
}
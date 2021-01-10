package main

import "fmt"

//ByteCounter .
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Println(c)
	// Fprintf formats according to a format specifier and writes to w. It returns the number of bytes written and any write error encountered.
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

	// Stringer is implemented by any value that has a String method, which defines the “native” format for that value. The String method is used to print values passed as an operand to any format that accepts a string or to an unformatted printer such as Print.
	// fmt.Stringer()
}
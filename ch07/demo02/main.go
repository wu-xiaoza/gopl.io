package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)


func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	fmt.Println(f == os.Stdout)

	w = os.Stdout
	rw := w.(io.ReadWriter)
	fmt.Printf("%#v\n", rw)
	
	w = os.Stdout
	f, ok := w.(*os.File)
	fmt.Printf("%v\t%t\n", f, ok)
	b, ok := w.(*bytes.Buffer)
	fmt.Printf("%v\t%t\n", b, ok)
}
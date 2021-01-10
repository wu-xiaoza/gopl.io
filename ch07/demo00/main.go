package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

//Shaper .
type Shaper interface {
	Area() float64
	Perimete() float64
}

//Rect .
type Rect struct {
	width  float64
	height float64
}

//Area .
func (r Rect) Area() float64 {
	return r.width * r.height
}

//Perimete .
func (r Rect) Perimete() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	var s Shaper
	s = Rect{5, 4} //Rect{5, 4}有Shaper方法
	n := Rect{5, 4}
	fmt.Println(s)
	fmt.Println(n)

	var w io.Writer
	fmt.Println(w)
	w = new(bytes.Buffer)
	fmt.Println(w)
	w = os.Stdout              //*os.File有Write方法
	fmt.Println(w)
	// w = time.Second         // compile error: time.Duration lacks Write method

	var rwc io.ReadWriteCloser
	fmt.Println(rwc)
	// rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method
	rwc = os.Stdout
	fmt.Println(rwc)

	w = rwc
	// rwc = w                 // compile error: io.Writer lacks Close method

	os.Stdout.Write([]byte("hello"))
	os.Stdout.Close()          //*os.File有Close方法

	w = os.Stdout
	w.Write([]byte("hello"))  //io.Writer有Write方法
	// w.Close()                // compile error: io.Writer lacks Close method

	// var v io.Writer = new(bytes.Buffer)
	// var _ io.Writer = (bytes.Buffer)(nil)
}

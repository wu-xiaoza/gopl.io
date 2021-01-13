package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)


var (
	//Bool defines a bool flag with specified name, default value, and usage string. The return value is the address of a bool variable that stores the value of the flag.
	n = flag.Bool("n", false, "omit trailling newlinne")
	//String defines a string flag with specified name, default value, and usage string. The return value is the address of a string variable that stores the value of the flag.
	s = flag.String("s", " ", "separator")
)

var out io.Writer = os.Stdout

func main() {

}

func echo(newline bool, sep string, args []string) error {
	// Join concatenates the elements of its first argument to create a single string. The separator string sep is placed between elements in the resulting string.
	fmt.Fprint(out, strings.Join(args, sep))
	if newline {
		fmt.Fprintln(out)
	}
	return nil
}
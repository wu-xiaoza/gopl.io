package tempconv

import (
	"flag"
	"fmt"
)

//Celsius .
type Celsius float64
//Fahrenheit .
type Fahrenheit float64

//CToF .
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5+32)
}

//FToC .
func FToC(f Fahrenheit) Celsius {
	return Celsius((f-32)*5/9)
}

func (c Celsius) String() string {
	// Sprintf formats according to a format specifier and returns the resulting string.
	return fmt.Sprintf("%g℃", c)
}

type celsiusFlag struct{
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	//Sscanf scans the argument string, storing successive space-separated values into successive arguments as determined by the format. It returns the number of items successfully parsed. Newlines in the input must match newlines in the format.
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "℃":
		f.Celsius = Celsius(value)
		return nil
	case "F", "℉":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

//CelsiusFlag .
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{ value }
	//CommandLine is the default set of command-line flags, parsed from os.Args. The top-level functions such as BoolVar, Arg, and so on are wrappers for the methods of CommandLine.
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
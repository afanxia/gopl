package tempconv

import (
	"flag"
	"fmt"

	tc "github.com/afanxia/gopl/ch2/tempconv"
)

// *celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct{ tc.Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C":
		f.Celsius = tc.Celsius(value)
		return nil
	case "F":
		f.Celsius = tc.FToC(tc.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value tc.Celsius, usage string) *tc.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

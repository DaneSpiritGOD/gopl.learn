package temperature

import (
	"flag"
	"fmt"
)

// Celsius _
type Celsius float64

// Fahrenheit _
type Fahrenheit float64

// Kelvin _
type Kelvin float64

const (
	// AbsoluteZeroC _
	AbsoluteZeroC Celsius = -273.15
	// FreezingC _
	FreezingC Celsius = 0
	// BoilingC _
	BoilingC Celsius = 100
	// AbsoluteZeroK _
	AbsoluteZeroK Celsius = -273.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }

// CelsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct{ Celsius }

// Set implements flag Value interface
func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

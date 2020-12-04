package temperature

// CToF convert C to F
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC convert F to C
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC convert K to C
func KToC(k Kelvin) Celsius { return Celsius(k) }

// CToK convert C to K
func CToK(c Celsius) Kelvin { return Kelvin(c) }

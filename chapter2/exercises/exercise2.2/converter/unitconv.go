// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 2.1

package converter

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k + Kelvin(AbsoluteZeroC)) }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k + Kelvin(AbsoluteZeroC)) * 9 / 5 + 32) }

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin { return Kelvin((f - 32) * 5 / 9 - Fahrenheit(AbsoluteZeroC)) }

// FtoM converts a Foot length to Meter.
func FtoM(f Foot) Meter { return Meter(f / 3.281) }

// MtoF converts a Meter length to Foot.
func MtoF(m Meter) Foot { return Foot(m * 3.281) }

// PtoK converts a Pound weight to Kilogram.
func PtoK(p Pound) Kilogram { return Kilogram(p / 2.205) }

// KtoP converts a Kilogram weight to Pound.
func KtoP(k Kilogram) Pound { return Pound(k * 2.205) }

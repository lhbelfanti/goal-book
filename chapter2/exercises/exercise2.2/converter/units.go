// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 2.2
package converter

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const AbsoluteZeroC Celsius = -273.15

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string { return fmt.Sprintf("%g°K", k) }

type Foot float64
type Meter float64

func (f Foot) String() string  { return fmt.Sprintf("%g Foot", f) }
func (m Meter) String() string { return fmt.Sprintf("%g Meters", m) }

type Pound float64
type Kilogram float64

func (p Pound) String() string { return fmt.Sprintf("%g Punds", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%g Kilograms", k) }

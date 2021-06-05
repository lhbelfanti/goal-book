// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Page 43.

// Example:
// go build main.go
// ./main 32
// > Output: 32°F = 0°C, 32°C = 89.6°F
// ./main 212
// > Output: 212°F = 100°C, 212°C = 413.6°F
// ./main -40
// -40°F = -40°C, -40°C = -40°F

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

	"gitlab.com/lhbelfanti/goal-book/chapter2/examples/05tempconv1"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}

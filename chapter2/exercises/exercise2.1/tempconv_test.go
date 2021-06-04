// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Page 39.

package tempconv

import "fmt"

func Example_kelvinConversion() {
	k := CToK(1)
	fmt.Println(k.String())	// "274.15°K"

	c := KToC(274.15)
	fmt.Println(c.String())	// "1°C"

	k = FToK(1)
	fmt.Println(k.String())	// "255.92777777777775°K"

	f := KToF(255.928)
	fmt.Println(f.String())	// "1.0004000000000346°F" ~ 1°F

	// Output:
	// 274.15°K
	// 1°C
	// 255.92777777777775°K
	// 1.0004000000000346°F
}

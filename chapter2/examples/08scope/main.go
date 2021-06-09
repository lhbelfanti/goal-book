// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Page 47.

// Shows an example of how the scope works in golang. Has three different variables called x
// because each declaration appears in a different lexical block.
package main

import "fmt"

func main() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
		}
	}
}

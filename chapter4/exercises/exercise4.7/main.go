// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 4.7
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("Hello　世界")
	reverse(b)
	s := string(b)
	fmt.Println(s)
}

func rev(b []byte) {
	size := len(b)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[size-1-i] = b[size-1-i], b[i]
	}
}

func reverse(b []byte) {
	// Reverse each UTF-8 rune
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		rev(b[i : i+size])
		i += size
	}
	// Reverse the whole bytes holding Reversed-UTF-8
	rev(b)
}

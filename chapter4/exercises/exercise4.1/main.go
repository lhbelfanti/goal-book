// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 4.1
package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	c1 := sha256.Sum256([]byte("xyz"))
	c2 := sha256.Sum256([]byte("XYZ"))

	fmt.Fprintf(os.Stdout, "Different bit count is %d\n", countDiffBits(c1, c2))
}

func countDiffBits(sha1, sha2 [32]uint8) int {
	n := 0
	for i := 0; i < len(sha1); i++ {
		s := sha1[i] ^ sha2[i]
		n += bitCount(s)
	}
	return n
}

func bitCount(x uint8) int {
	n := 0
	for x != 0 {
		x = x & (x - 1)
		n++
	}
	return n
}

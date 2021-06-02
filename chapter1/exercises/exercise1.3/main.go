// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//Exercise 1.3
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	Chapter1example3()
	elapsed := time.Since(start)
	fmt.Printf("----> Potential inefficient version took %s", elapsed)

	fmt.Println()

	start = time.Now()
	Chapter1example4()
	elapsed = time.Since(start)
	fmt.Printf("----> strings.Join version %s", elapsed)
}

func Chapter1example3() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func Chapter1example4() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}


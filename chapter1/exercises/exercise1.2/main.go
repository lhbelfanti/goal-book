// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 1.2
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for idx, arg := range os.Args[1:] {
		sep = " "
		s = sep + strconv.Itoa(idx) + sep + arg
		fmt.Println(s)
	}
}

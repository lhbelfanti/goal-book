// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 4.5
package main

import (
	"fmt"
)

func main() {
	s := []string{"a", "b", "c", "d", "d", "d", "e", "f", "f"}
	fmt.Println(adjacent(s)) // "[a b c d e f]"
}

func adjacent(s []string) []string{
	i := 0
	for _, v := range s {
		if s[i] == v { // Avoid the duplicated values. Non unique value.
			continue
		}
		i++
		s[i] = v // Adds the unique value and evaluates in the next iteration the following one is unique too
	}

	return s[:i+1]
}

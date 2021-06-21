// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 4.4
package main

import (
	"fmt"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rotate(a1, 2)
	fmt.Println(a1) // "[2 3 4 5 6 7 8 9 0 1]"

	a2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rotate(a2, 5)
	fmt.Println(a2) // "[5 6 7 8 9 0 1 2 3 4]"
}

func rotate(s []int, n int) {
	l := len(s)
	if l < n || n < 0 {
		return
	}

	s2 := make([]int, l)
	copy(s2, s[n:])
	for i, v := range s[0:n] {
		s2[l - n + i] = v
	}
	copy(s, s2)
}

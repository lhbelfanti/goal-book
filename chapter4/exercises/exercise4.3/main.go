// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 4.3
package main

import (
	"fmt"
)

func main() {
	a1 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	reverse(&a1)
	fmt.Println(a1) // "[9 8 7 6 5 4 3 2 1 0]"

	a2 := [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	reverse(&a2)
	fmt.Println(a2) // "[0 1 2 3 4 5 6 7 8 9]"
}

func reverse(arr *[10]int) {
	l := len(arr)
	for i := 0; i < l/2; i++ {
		j := l - 1 - i
		arr[i], arr[j] = arr[j], arr[i]
	}
}

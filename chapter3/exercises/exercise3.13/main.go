// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.13
package main

import "fmt"

const (
	Kilo = 1000
	B    = 1
	KB = B * Kilo
	MB = KB * Kilo
	GB = MB * Kilo
	TB = GB * Kilo
	PB = TB * Kilo
	EB = PB * Kilo
	ZB = EB * Kilo
	YB = ZB * Kilo
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB / KB)
	fmt.Println(GB / MB)
	fmt.Println(TB / GB)
	fmt.Println(PB / TB)
	fmt.Println(EB / PB)
	fmt.Println(ZB / EB)
	fmt.Println(YB / ZB)
}



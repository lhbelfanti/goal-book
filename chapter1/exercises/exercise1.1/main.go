// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 1.1
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

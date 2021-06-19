// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.12
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("%t", isAnagram(os.Args[1], os.Args[2]))
}

func isAnagram(s1, s2 string) bool {
	if s1 == s2 {
		return false
	}

	if len(s1) != len(s2) {
		return false
	}

	for _, c := range s1 {
		if strings.Contains(s2, string(c)) {
			s2 = strings.Replace(s2, string(c), "", 1)
		}
	}
	if len(s2) > 0 {
		return false
	}

	return true
}

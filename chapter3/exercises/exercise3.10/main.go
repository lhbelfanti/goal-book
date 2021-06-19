// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.10
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer

	for i := 0; i < len(s); i++ {
		if (i + 1) % 3 == 0 {
			buf.WriteString(",")
		}

		if len(s[i:]) <= 3 {
			buf.WriteString(s[i:])
			break
		}

		buf.WriteString(s[i:i+1])
	}

	return buf.String()
}

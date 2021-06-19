// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.11
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s -> %s\n", os.Args[i], comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer

	n := strings.Index(s, ".")
	if n == -1 {
		n = len(s)
	}

	if strings.HasPrefix(s, "+") ||
		strings.HasPrefix(s, "-") {
		buf.WriteByte(s[0])
		s = s[1:]
		n--
	}

	i := n % 3
	if i == 0 {
		i = 3
	}
	buf.WriteString(s[:i])

	for j, c := range s[i:n] {
		if j%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteRune(c)
	}
	buf.WriteString(s[n:])

	return buf.String()
}

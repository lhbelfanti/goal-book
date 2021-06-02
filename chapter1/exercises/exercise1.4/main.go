// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 1.4
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		// NOTE: ignoring potential errors from input.Err()
		line := input.Text()
		counts[line]++
		if counts[line] > 1 {
			fmt.Printf("File: %v - is adding 1 to \"%s\" - it had %d before\n", filename, line, counts[line] - 1)
		}
	}
}

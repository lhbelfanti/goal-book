// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 4.9
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stdout, "%s <inputfile> <outputfile>\n", os.Args[0])
		return
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}

	filepath := os.Args[2]
	fi, err := os.Create(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while opening file: %s\n", err)
		os.Exit(1)
	}

	words := make(map[string]int)
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		words[input.Text()]++
	}

	fmt.Fprint(fi, "Word\tCount\n")
	var keys []string
	for k := range words {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Fprintf(fi, "%s\t%d\n", k, words[k])
	}

	if err := fi.Close(); err != nil {
		fmt.Printf("Error while closing file: %s\n", err)
		os.Exit(1)
	}
}

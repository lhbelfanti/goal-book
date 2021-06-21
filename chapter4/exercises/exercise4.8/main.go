// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 4.8
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

const (
	Control = "Control"
	Letter = "Letter"
	Lower = "Lowercase Letter"
	Upper = "Uppercase Letter"
	Title = "Titlecase Letter"
	Space = "Spacing Mark"
	Number = "Number"
	Digit = "Digit"
	Mark = "Mark"
	Punct = "Punctuation"
	Symbol = "Symbol"
	Print = "Printable"
	NonPrintable = "Non Printable"
)

func main() {
	counts := make(map[string]map[rune]int)       // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		countCategory(r, counts)
		utflen[n]++
	}
	for c, n := range counts {
		fmt.Printf("\n▶ Unicode Category: %s\n", c)
		fmt.Printf("Rune\tCount\n")
		for k, v := range n {
			fmt.Printf("%q\t%d\n", k, v)
		}
	}
	fmt.Print("\nLen\tCount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func countCategory(r rune, counts map[string]map[rune]int) {
	if unicode.IsControl(r) {
		countRune(r, counts, Control)
	}
	if unicode.IsLetter(r) {
		countRune(r, counts, Letter)
	}
	if unicode.IsLower(r) {
		countRune(r, counts, Lower)
	}
	if unicode.IsUpper(r) {
		countRune(r, counts, Upper)
	}
	if unicode.IsTitle(r) {
		countRune(r, counts, Title)
	}
	if unicode.IsSpace(r) {
		countRune(r, counts, Space)
	}
	if unicode.IsNumber(r) {
		countRune(r, counts, Number)
	}
	if unicode.IsDigit(r) {
		countRune(r, counts, Digit)
	}
	if unicode.IsMark(r) {
		countRune(r, counts, Mark)
	}
	if unicode.IsPunct(r) {
		countRune(r, counts, Punct)
	}
	if unicode.IsSymbol(r) {
		countRune(r, counts, Symbol)
	}
	if unicode.IsPrint(r) {
		countRune(r, counts, Print)
	} else {
		countRune(r, counts, NonPrintable)
	}
}

func countRune(r rune, m map[string]map[rune]int, c string) {
	if m[c] == nil {
		m[c] = make(map[rune]int)
	}
	m[c][r]++
}

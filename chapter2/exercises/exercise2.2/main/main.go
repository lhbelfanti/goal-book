// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 2.2

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gitlab.com/lhbelfanti/goal-book/chapter2/exercises/exercise2.2/converter"
)

var t = flag.Bool("t", false, "convert temperatures")
var w = flag.Bool("w", false, "convert weights")
var l = flag.Bool("l", false, "convert lengths")

func main() {
	flag.Parse()

	if len(os.Args) == 1 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			v, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			fmt.Println()
			printUnits(v)
		}
	} else {
		for _, arg := range os.Args[1:] {
			if strings.HasPrefix(arg, "-") {
				continue
			}
			v, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}

			fmt.Println()

			if len(flag.Args()) > 0 {
				if *t {
					printTemperature(v)
				}
				if *w {
					printWeight(v)
				}
				if *l {
					printLength(v)
				}
			} else {
				printUnits(v)
			}
		}
	}
}

func printUnits(v float64) {
	printTemperature(v)
	printWeight(v)
	printLength(v)
}

func printTemperature(v float64) {
	f := converter.Fahrenheit(v)
	c := converter.Celsius(v)
	k := converter.Kelvin(v)
	fmt.Printf("%s = %s, %s = %s --- %s = %s, %s = %s --- %s = %s, %s = %s\n",
		f, converter.FToC(f), f, converter.FToK(f), c, converter.CToF(c), c, converter.CToK(c), k, converter.KToC(k), k, converter.KToF(k))
}

func printWeight(v float64) {
	p := converter.Pound(v)
	k := converter.Kilogram(v)
	fmt.Printf("%s = %s, %s = %s\n", p, converter.PtoK(p), k, converter.KtoP(k))
}

func printLength(v float64) {
	f := converter.Foot(v)
	m := converter.Meter(v)
	fmt.Printf("%s = %s, %s = %s\n", f, converter.FtoM(f), m, converter.MtoF(m))
}

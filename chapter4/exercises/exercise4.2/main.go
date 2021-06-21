// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Example:
// go run main.go -sha384 -sha512 true true

// Exercise 4.2
package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var sha = flag.Int("sha", 256, "hash width (384 or 512)")

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	flag.Parse()
	switch *sha {
	case 256:
		fmt.Printf("Sha256: %x\n", sha256.Sum256([]byte(text)))
	case 384:
		fmt.Printf("Sha384: %x\n", sha512.Sum384([]byte(text)))
	case 512:
		fmt.Printf("Sha512: %x\n", sha512.Sum512([]byte(text)))
	}
}

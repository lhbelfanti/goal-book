// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 1.10
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	filepath := os.Args[1]

	for _, url := range os.Args[2:] {
		go fetch(url, ch) // start a goroutine
	}

	fi, err := os.Create(filepath)
	if err != nil {
		fmt.Printf("Error while opening file: %s\n", err)
		os.Exit(1)
	}

	for range os.Args[2:] {
		fi.WriteString(<-ch) // receive from channel ch
	}

	fi.WriteString(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds()))

	if err := fi.Close(); err != nil {
		fmt.Printf("Error while closing file: %s\n", err)
		os.Exit(1)
	}
}

func fetch(url string, ch chan<- string) {
	fmt.Printf("Fetching %s...\n", url)
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s\n", secs, nbytes, url)
}

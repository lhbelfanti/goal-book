// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Example:
//
// To create the index:
// go run xkcd.go index ./index.json
//
// To get a specific comic:
// go run xkcd.go get 123
//
// After creating the index you can search for one or more words in the comics of the index:
// go run xkcd.go search ./index.json "dog cute"

// Exercise 4.12
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Comic struct {
	Num int
	Year, Month, Day string
	Title string
	Transcript string
	Alt string
	Link string  `json:"img"`
}

const InfoURL = "https://xkcd.com/info.0.json"

const usage = `xkcd get N
xkcd index OUTPUT_FILE
xkcd search INDEX_FILE QUERY`

func usageDie() {
	fmt.Println(usage)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	cmd := os.Args[1]
	switch cmd {
	case "get":
		if len(os.Args) != 3 {
			usageDie()
		}
		n, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "N (%s) must be an int", os.Args[1])
			usageDie()
		}
		comic, err := getComic(n)
		if err != nil {
			log.Fatal("Error getting comic", err)
		}
		fmt.Println(comic)
	case "index":
		if len(os.Args) != 3 {
			usageDie()
		}
		err := createIndex(os.Args[2])
		if err != nil {
			log.Fatal("Error serializing indexes", err)
		}
	case "search":
		if len(os.Args) != 4 {
			usageDie()
		}
		filename := os.Args[2]
		query := os.Args[3]
		err := search(query, filename)
		if err != nil {
			log.Fatal("Error searching index", err)
		}
	default:
		usageDie()
	}
}

func createIndex(filename string) error {
	comics := make(map[string]Comic)
	c, err := getComic(-1)
	if err != nil {
		return err
	}

	comics[strconv.Itoa(c.Num)] = c
	max := c.Num

	workers := 100

	comicNums := make(chan int, workers)
	done := make(chan int, 0)
	lock := sync.RWMutex{}

	for i := 1; i < workers; i++ {
		go downloader(comicNums, comics, done, &lock)
	}

	go dispatcher(comicNums, max)

	<-done

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error while opening file to create the index: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	err = enc.Encode(comics)
	if err != nil {
		return err
	}

	fmt.Printf("\nIndex saved into: %s\n", filename)

	return nil
}

func dispatcher(comicNums chan int, max int) {
	for i := 1; i < max; i++ {
		comicNums <- i
	}
	close(comicNums)
}

func downloader(comicNums chan int, comics map[string]Comic, done chan int, lock *sync.RWMutex) {
	for n := range comicNums {
		comic, err := getComic(n)
		if err != nil {
			log.Printf("Can't get comic %d: %s", n, err)
			continue
		}

		lock.Lock()
		comics[strconv.Itoa(comic.Num)] = comic
		fmt.Printf("Adding comic %d to index\n", comic.Num)
		lock.Unlock()

	}
	fmt.Println("Downloader goroutine finished")
	done <- 1
}

func getComic(n int) (Comic, error){
	var url string
	if n == -1 {
		url =  InfoURL
	} else {
		url = fmt.Sprintf("https://xkcd.com/%d/info.0.json", n)
	}

	var comic Comic
	resp, err := http.Get(url)
	if err != nil {
		return comic, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return comic, fmt.Errorf("can't get comic %d: %s", n, resp.Status)
	}
	if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return comic, err
	}
	return comic, nil
}

func search(query, filename string) error {
	file, e := os.Open(filename)
	if e != nil {
		fmt.Println("Error while opening the file: %s. Error: %v", filename, e)
		return e
	}
	defer file.Close()

	var comics map[string]Comic
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&comics)
	if err != nil {
		fmt.Println("Error while decoding the json: %v", err)
		return err
	}

	words := strings.Fields(query)
	for _, v := range comics {
		ti := wordsInComic(words, strings.Fields(v.Title))
		if ti {
			printComic(v)
		} else {
			tr := wordsInComic(words, strings.Fields(v.Transcript))
			if tr {
				printComic(v)
			}
		}
	}

	return nil
}


func printComic(c Comic) {
	fmt.Printf(
`------------------------------------------------------------
Comic: %d
Title: %s
Link: %s
Transcript: %q
`, c.Num, c.Title, c.Link, c.Transcript)
}

func wordsInComic(words []string, comicWords []string) bool {
	for _, w := range words {
		for _, cw := range comicWords {
			if w == cw {
				return true
			}
		}
	}

	return false
}

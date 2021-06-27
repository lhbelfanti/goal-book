// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 4.13
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Movie struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Released string `json:"Released"`
	Duration  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Writer   string `json:"Writer"`
	Actors   string `json:"Actors"`
	Plot     string `json:"Plot"`
	Poster   string `json:"Poster"`
}

const usage = `poster API_KEY MOVIE`

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	movie := os.Args[2]
	m, err := getMovieInfo(os.Args[1], movie)
	if err != nil {
		fmt.Printf("Failed to get %s poster. Error: %v", movie, err)
		os.Exit(1)
	} else {
		err = downloadPoster(m.Title, m.Poster)
		if err != nil {
			fmt.Printf("Failed to download %s poster. Error: %v", movie, err)
			os.Exit(1)
		}

		fmt.Printf("\n%s poster downloaded correctly.\n", movie)
		fmt.Printf("Info:\n" +
			"Title: %s\n" +
			"Year: %s\n" +
			"Released: %s\n" +
			"Duration: %s\n" +
			"Genre: %s\n" +
			"Director: %s\n" +
			"Writer: %s\n" +
			"Actors: %s\n" +
			"Plot: %s\n",
			m.Title, m.Year, m.Released, m.Duration, m.Genre, m.Director, m.Writer, m.Actors, m.Plot)
	}
}


func getMovieInfo(api, movieName string) (*Movie, error) {
	movieName = strings.ReplaceAll(movieName, " ", "+")
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&t=%s", api, movieName)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Can't get %s movie. Error: %s", movieName, resp.Status)
	}

	var movie Movie
	if err = json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return nil, err
	}
	return &movie, nil
}

func downloadPoster(movieName, url string) error{
	fileName := movieName + ".jpg"

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Can't download %s poster. Error: %s", movieName, resp.Status)
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

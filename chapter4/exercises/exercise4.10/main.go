// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 4.10
package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	MONTH = "Less than a month"
	YEAR = "Less than a year"
	YEARS = "More than a year"
	OTHER = "Other issues"
)

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	sorted := sortByDate(result.Items)
	fmt.Println("\nIssues reported by age:")

	for k, v := range sorted {
		fmt.Printf("\n%s:\n", k)
		for _, item := range v {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
}

func sortByDate(issues []*Issue) map[string][]*Issue{
	sortMap := make(map[string][]*Issue, 4)

	now := time.Now()
	month := now.AddDate(0, -1, 0)
	year := now.AddDate(-1, 0, 0)

	for _, item := range issues {
		switch {
		case item.CreatedAt.After(month):
			sortMap[MONTH] = append(sortMap[MONTH], item)
		case item.CreatedAt.After(year):
			sortMap[YEAR] = append(sortMap[YEAR], item)
		case item.CreatedAt.Before(year):
			sortMap[YEARS] = append(sortMap[YEARS], item)
		default:
			sortMap[OTHER] = append(sortMap[OTHER], item)
		}
	}

	return sortMap
}

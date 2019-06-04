package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/deputatovn/gopl-solutions/ch04/ex4.10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	timeNow := time.Now()
	monthAgo := timeNow.AddDate(0, -1, 0)
	yearAgo := timeNow.AddDate(-1, 0, 0)

	issueMonthAgo := &github.IssuesCategory{}
	issueYearAgo := &github.IssuesCategory{}
	issueOverAYearAgo := &github.IssuesCategory{}

	for _, item := range result.Items {
		if item.CreatedAt.After(monthAgo) {
			issueMonthAgo.Items = append(issueMonthAgo.Items, item)
			issueMonthAgo.TotalCount++
		} else if item.CreatedAt.Before(monthAgo) && item.CreatedAt.After(yearAgo) {
			issueYearAgo.Items = append(issueYearAgo.Items, item)
			issueYearAgo.TotalCount++
		} else {
			issueOverAYearAgo.Items = append(issueOverAYearAgo.Items, item)
			issueOverAYearAgo.TotalCount++
		}
	}

	if issueMonthAgo.TotalCount > 0 {
		fmt.Printf("%d issues less than a month old:\n", issueMonthAgo.TotalCount)
		for i, item := range issueMonthAgo.Items {
			fmt.Printf("%-3d %s #%-5d %9.9s %.55s\n",
				i, item.CreatedAt, item.Number, item.User.Login, item.Title)
		}
		fmt.Println()
	}

	if issueYearAgo.TotalCount > 0 {
		fmt.Printf("%d issues less than a year old:\n", issueYearAgo.TotalCount)
		for i, item := range issueYearAgo.Items {
			fmt.Printf("%-3d %s #%-5d %9.9s %.55s\n",
				i, item.CreatedAt, item.Number, item.User.Login, item.Title)
		}
		fmt.Println()
	}

	if issueOverAYearAgo.TotalCount > 0 {
		fmt.Printf("%d issues more than a year old:\n", issueOverAYearAgo.TotalCount)

		for i, item := range issueOverAYearAgo.Items {
			fmt.Printf("%-3d %s #%-5d %9.9s %.55s\n",
				i, item.CreatedAt, item.Number, item.User.Login, item.Title)
		}
	}
}

package main

import (
	"TheGoProgrammingLanguage/gosnippets/github"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var result github.IssuesSearchResult

// SearchIssues queries the GitHub issue tracker.
/*
The SearchIssues function makes an HTTP request and decodes the result as JSON.
Since the query terms presented by a user could contain characters
like ? and & that have special meaning in a URL, we use url.QueryEscape
to ensure that they are taken literally.
*/

func SearchIssues(terms []string) (*github.IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	// resp, err := http.Get(github.IssuesURL + "?q=" + q)
	// if err != nil {
	// 	return nil, err
	// }

	// For long-term stability, instead of http.Get, use the
	// variant below which adds an HTTP request header indicating
	// that only version 3 of the GitHub API is acceptable.
	//
	req, err := http.NewRequest("GET", github.IssuesURL+"?q="+q, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(
		"Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

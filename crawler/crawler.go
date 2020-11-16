package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

//InitCrawler initializes the basic BFS logic
func InitCrawler(opts CrawlerOptions, state State) State {
	fmt.Println("[Crawler] Starting crawler")

	visited := getAlreadyVisitedURLs(state)
	URLQueue := []string{opts.baseURL}
	indexedCount := 0

	if visited[opts.baseURL] {
		fmt.Println("[Crawler] BaseURL is already stored in the database. Consider changing the baseURL.")
		os.Exit(0)
	}
	visited[opts.baseURL] = true

	for len(URLQueue) > 0 && indexedCount < opts.limit {
		url := URLQueue[0]
		URLQueue = URLQueue[1:]

		goQueryDoc, err := getURLDocument(url)
		if err != nil {
			// Uncomment the log below for lots of spamming
			// fmt.Println("[Crawler] Url cannot be fetched with error:", url, err)
			continue
		}

		document := parseGoQueryDocument(url, goQueryDoc, opts)
		state.DF = populateDF(state.DF, document)
		state.documents = append(state.documents, document)
		indexedCount++

		fmt.Println("[Crawler] Total parsed urls:", indexedCount)
		fmt.Println("[Crawler] Visiting url:", url)

		for _, url := range document.neighbors {
			if _, ok := visited[url]; ok {
				// fmt.Println("[Crawler] URL is already in the visited slice", url)
				continue
			}

			visited[url] = true
			URLQueue = append(URLQueue, url)
		}
	}

	return state
}

func getURLDocument(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("[Crawler] Url cannot be fetched", url)
		return nil, errors.New("Url cannot be fetched")
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("[Crawler] Url is broken:", url)
		return nil, errors.New("Url returned errored status")
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func getAlreadyVisitedURLs(state State) map[string]bool {
	visited := make(map[string]bool)

	for _, document := range state.documents {
		visited[document.url] = true
	}

	return visited
}

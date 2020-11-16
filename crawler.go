package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

// Basic BFS logic
func initCrawler(opts CrawlerOptions, state State) State {
	fmt.Println("[Crawler] Starting crawler")

	visited := getAlreadyVisitedURLs(state)
	URLQueue := []string{opts.baseURL}
	indexedCount := 0

	if visited[opts.baseURL] {
		fmt.Println("[Crawler] BaseURL is already stored in the database. Consider changing the baseURL.")
		os.Exit(0)
	}

	for len(URLQueue) > 0 && indexedCount < opts.limit {
		url := URLQueue[0]
		URLQueue = URLQueue[1:]

		goQueryDoc, err := getURLDocument(url)
		if err != nil {
			// Uncomment the log below for lots of spamming
			// fmt.Println("[Crawler] Url cannot be fetched with error:", url, err)
			continue
		}

		document := parseGoQueryDocument(url, goQueryDoc)
		state.DF = populateDF(state.DF, document)
		state.documents = append(state.documents, document)
		indexedCount++

		fmt.Println("[Crawler] Total parsed urls:", indexedCount)
		fmt.Println("[Crawler] Visiting url:", url)

		for _, url := range document.neighbors {
			if _, ok := visited[url]; ok {
				// fmt.Println("[Crawler] URL has been already visited", url)
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
	checkErr(err)
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("[CRAWLER] Status code error: %d %s", res.StatusCode, res.Status)
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

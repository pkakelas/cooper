package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Basic BFS logic
func initCrawler(baseURL string, limit int) {
	visited := make(map[string]bool)
	parsed := []string{}
	URLQueue := []string{baseURL}

	fmt.Println("Starting crawler")

	for len(URLQueue) > 0 && len(parsed) < limit {
		url := URLQueue[0]
		URLQueue = URLQueue[1:]

		goQueryDoc, err := getURLDocument(url)
		if err != nil {
			// Uncomment the log below for lots of spamming
			// fmt.Println("[Crawler] Url cannot be fetched with error:", url, err)
			continue
		}

		document := parseGoQueryDocument(url, goQueryDoc)
		fmt.Println(document.stems)
		parsed = append(parsed, url)

		fmt.Println("[Crawler] Total parsed urls:", len(parsed))
		fmt.Println("[Crawler] Visiting url:", url)

		for _, url := range document.neighbors {
			if _, ok := visited[url]; ok {
				continue
			}

			visited[url] = true
			URLQueue = append(URLQueue, url)
		}
	}
}

func getURLDocument(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

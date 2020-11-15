package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Basic BFS logic
func initCrawler(baseURL string, limit int) (indexed []Document, DF DocumentFrequency) {
	indexed = []Document{}
	DF = make(DocumentFrequency)

	visited := make(map[string]bool)
	URLQueue := []string{baseURL}

	fmt.Println("Starting crawler")

	for len(URLQueue) > 0 && len(indexed) < limit {
		url := URLQueue[0]
		URLQueue = URLQueue[1:]

		goQueryDoc, err := getURLDocument(url)
		if err != nil {
			// Uncomment the log below for lots of spamming
			// fmt.Println("[Crawler] Url cannot be fetched with error:", url, err)
			continue
		}

		document := parseGoQueryDocument(url, goQueryDoc)
		DF = populateDF(DF, document)
		indexed = append(indexed, document)

		fmt.Println("[Crawler] Total parsed urls:", len(indexed))
		fmt.Println("[Crawler] Visiting url:", url)

		for _, url := range document.neighbors {
			if _, ok := visited[url]; ok {
				continue
			}

			visited[url] = true
			URLQueue = append(URLQueue, url)
		}
	}

	return
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

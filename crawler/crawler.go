package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

//InitCrawler initializes the basic BFS logic
func InitCrawler(opts CrawlerOptions, state State) State {
	worklist := make(chan []string)  // All urls found from the goroutines
	unseenLinks := make(chan string) // Only unseed urls
	seen := getAlreadyVisitedURLs(state)
	var wg sync.WaitGroup

	if seen[opts.baseURL] {
		fmt.Println("[Crawler] BaseURL is already stored in the database. Consider changing the baseURL.")
		os.Exit(0)
	}

	go func() {
		worklist <- []string{opts.baseURL}
	}()

	// Create crawler's goroutines to fetch and parse site
	for i := 0; i < opts.threads; i++ {
		go func() {
			for link := range unseenLinks {
				wg.Add(1)
				foundLinks := crawl(link, &state, opts)
				wg.Done()

				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	for list := range worklist {
		for _, link := range list {
			if len(seen) == opts.limit {
				wg.Wait()
				return state
			}
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}

	wg.Wait()
	return state
}

func crawl(url string, state *State, opts CrawlerOptions) []string {
	fmt.Println("[Crawler] Crawling", url)

	goQueryDoc, err := getURLDocument(url)
	if err != nil {
		fmt.Print("[Crawler] Url cannot be fetched")
		return []string{}
	}

	document := parseGoQueryDocument(url, goQueryDoc, opts)
	(*state).DF = populateDF((*state).DF, document)
	(*state).documents = append((*state).documents, document)

	return document.neighbors
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

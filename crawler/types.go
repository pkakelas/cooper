package main

// State is a wraper containing all the state vars of the crawler
type State struct {
	documents []Document
	DF        DocumentFrequency
}

// CrawlerOptions contains all the options given from the CLI
type CrawlerOptions struct {
	baseURL            string
	limit              int
	threads            int
	loadData           bool
	includeQueryParams bool
	serverMode         bool
}

// Document represents a parsed HTML file
type Document struct {
	tf        TermFrequency
	id        string
	title     string
	url       string
	neighbors []string
	stems     []string
}

// QueryResult represents a result-document of the query
type QueryResult struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
	sim   float64
}

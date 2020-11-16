package main

// State is a wraper containing all the state vars of the crawler
type State struct {
	documents []Document
	DF        DocumentFrequency
}

// CrawlerOptions contains all the options given from the CLI
type CrawlerOptions struct {
	baseURL  string
	limit    int
	loadData bool
	threads  int
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
	id    string
	title string
	url   string
	sim   float64
}

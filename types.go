package main

// CrawlerOptions contains all the options given from the CLI
type CrawlerOptions struct {
	baseURL  string
	maxSites int
	keepData bool
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

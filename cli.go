package main

import (
	"flag"
	"fmt"
	"os"
)

func initCLI() CrawlerOptions {
	fmt.Println("Welcome to Cooper, an simple and lightweight crawler written in Golang!\n ")
	return parseFlags()
}

func parseFlags() CrawlerOptions {
	baseURLPtr := flag.String("base_url", "", "The maximum sites that Cooper should visit")
	maxSitesPtr := flag.Int("max", 50, "The maximum sites that Cooper should visit")
	keepDataPtr := flag.Bool("keep_data", false, "Whether or not the existing crawls should be kept")
	threadsPtr := flag.Int("threads", 1, "How many crawl threads Cooper should create")
	flag.Parse()

	if len(*baseURLPtr) == 0 {
		fmt.Println("The base URL should be defined.\nUse --help for more info.")
		os.Exit(0)
	}

	return CrawlerOptions{
		baseURL:  *baseURLPtr,
		maxSites: *maxSitesPtr,
		keepData: *keepDataPtr,
		threads:  *threadsPtr,
	}
}

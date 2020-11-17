package main

import (
	"flag"
	"fmt"
	"os"
)

//InitCLI handles all the command interface and returns CrawlerOptions
func InitCLI() CrawlerOptions {
	fmt.Println("Welcome to Cooper, an simple and lightweight crawler written in Golang!")
	fmt.Println(getCooper())
	return parseFlags()
}

func parseFlags() CrawlerOptions {
	serverModePtr := flag.Bool("server_mode", false, "Work in server mode for serving data to the cooper frontend")
	baseURLPtr := flag.String("base_url", "", "The maximum sites that Cooper should visit")
	limitPtr := flag.Int("limit", 50, "The maximum sites that Cooper should visit")
	loadDataPtr := flag.Bool("load_existed_data", true, "Whether or not the existing crawls should be loaded")
	threadsPtr := flag.Int("threads", 1, "How many crawl threads Cooper should create")
	includeQueryParamsPtr := flag.Bool("include_query_params", true, "Should Cooper consider test.com?query and test.com as the same document?")
	flag.Parse()

	if !*serverModePtr {
		if len(*baseURLPtr) == 0 {
			fmt.Println("The base URL should be defined.\nUse --help for more info.")
			os.Exit(0)
		}
		if !isValidURI(*baseURLPtr) {
			fmt.Println("The base URL is not valid. Please use a url like https://github.com.")
			os.Exit(0)
		}
	}

	return CrawlerOptions{
		baseURL:            *baseURLPtr,
		serverMode:         *serverModePtr,
		limit:              *limitPtr,
		threads:            *threadsPtr,
		loadData:           *loadDataPtr,
		includeQueryParams: *includeQueryParamsPtr,
	}
}

func getCooper() string {
	return `
	   _=,_
	o_/6 /#\
	\__ |##/
	='|--\
	/   #'-.
	\#|_   _'-. /
	 |/ \_( # |"
	C/ ,--___/
	`
}

package main

import (
	"flag"
	"fmt"
	"os"
)

func initCLI() CrawlerOptions {
	fmt.Println("Welcome to Cooper, an simple and lightweight crawler written in Golang!")
	fmt.Println(getCooper())
	return parseFlags()
}

func parseFlags() CrawlerOptions {
	baseURLPtr := flag.String("base_url", "", "The maximum sites that Cooper should visit")
	limitPtr := flag.Int("limit", 50, "The maximum sites that Cooper should visit")
	loadDataPtr := flag.Bool("load_existed_data", true, "Whether or not the existing crawls should be loaded")
	threadsPtr := flag.Int("threads", 1, "How many crawl threads Cooper should create")
	flag.Parse()

	if len(*baseURLPtr) == 0 {
		fmt.Println("The base URL should be defined.\nUse --help for more info.")
		os.Exit(0)
	}

	return CrawlerOptions{
		baseURL:  *baseURLPtr,
		limit:    *limitPtr,
		loadData: *loadDataPtr,
		threads:  *threadsPtr,
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
	C/ ,--___/`
}

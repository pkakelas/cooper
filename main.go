package main

func main() {
	options := initCLI()
	state := initState(options)
	state = initCrawler(options, state)
	saveState(state)
}

func initState(options CrawlerOptions) (state State) {
	if options.loadData {
		state = loadState()
		return
	}

	state = State{
		documents: []Document{},
		DF:        make(DocumentFrequency),
	}

	return
}

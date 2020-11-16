package main

func main() {
	options := InitCLI()
	state := initState(options)
	state = initCrawler(options, state)
	SaveState(state)
}

func initState(options CrawlerOptions) (state State) {
	if options.loadData {
		return LoadState()
	}

	state = State{
		documents: []Document{},
		DF:        make(DocumentFrequency),
	}

	return
}

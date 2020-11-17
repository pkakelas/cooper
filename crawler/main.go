package main

func main() {
	options := InitCLI()
	state := initState(options)

	if options.serverMode {
		InitServer(state)
	} else {
		state = InitCrawler(options, state)
		SaveState(state)
	}
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

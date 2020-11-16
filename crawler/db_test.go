package main

import (
	"testing"
)

func TestStateSave(t *testing.T) {
	document := Document{
		tf:        map[string]float64{"word1": 0.01},
		title:     "title",
		url:       "url",
		id:        "id",
		neighbors: []string{"url2"},
	}
	df := DocumentFrequency{
		"word1": []string{"doc1", "doc2"},
	}
	state := State{
		documents: []Document{document},
		DF:        df,
	}

	createNewDatabase()
	SaveState(state)

	got := LoadState()
	gotDoc := got.documents[0]
	gotDF := got.DF
	if gotDoc.title != "title" || gotDoc.url != "url" || gotDoc.id != "id" || gotDoc.neighbors[0] != "url2" {
		t.Errorf("TestLoadState failed for testcase: %+v. Got %+v ", state, got)
	}

	if len(gotDF) != 1 || len(gotDF["word1"]) != 2 {
		t.Errorf("TestLoadState failed for testcase: %+v. Got %+v ", state, got)
	}
}

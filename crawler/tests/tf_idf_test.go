package main

import (
	"math"
	"reflect"
	"testing"
)

func TestTermFrequency(t *testing.T) {
	testCases := []struct {
		in  []string
		out TermFrequency
	}{
		{[]string{}, TermFrequency{}},
		{[]string{"a"}, TermFrequency{"a": 1}},
		{[]string{"a", "b"}, TermFrequency{"a": 0.5, "b": 0.5}},
		{[]string{"a", "b", "b", "c"}, TermFrequency{"a": 0.25, "b": 0.5, "c": 0.25}},
	}

	for _, testCase := range testCases {
		got := generateTermFrequency(testCase.in)
		if !reflect.DeepEqual(got, testCase.out) {
			t.Errorf("TestTermFrequencyFailed for testcase: %+v. Got %+v ", testCase, got)
		}
	}
}

func TestPopulateDF(t *testing.T) {
	document := Document{
		id:        "id5",
		tf:        TermFrequency{"a": 5, "c": 5},
		title:     "new page",
		url:       "https://mypage.com",
		stems:     []string{"a", "b"},
		neighbors: []string{"https://google.com"},
	}
	DF := DocumentFrequency{
		"a": {"id1", "id2", "id3"},
		"b": {"id1", "id2", "id4"},
	}

	got := populateDF(DF, document)

	if !stringInSlice("id5", got["a"]) || stringInSlice("id5", got["b"]) || !stringInSlice("id5", got["c"]) {
		t.Errorf("TestPopulateDF failed for testcase:\n(%+v, %s) => %+v", DF, "id5", got)
	}
}

func TestInversedDocumentFrequency(t *testing.T) {
	DF := DocumentFrequency{
		"a": {"id1"},
		"b": {"id1", "id2", "id3"},
	}
	got := inversedDocumentFrequency("a", DF, 15)

	if got != math.Log(15/float64(2)) {
		t.Errorf("TestInversedDocumentFreq failed for testcase:\n(%s, %+v, %d) => %+v", "a", DF, 15, got)
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

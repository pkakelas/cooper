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
	TF := TermFrequency{"a": 5, "c": 5}
	DF := DocumentFrequency{
		"a": {1, 2, 3},
		"b": {1, 2, 4},
	}

	got := populateDF(DF, 5, TF)

	if !intInSlice(5, got["a"]) || intInSlice(5, got["b"]) || !intInSlice(5, got["c"]) {
		t.Errorf("TestPopulateDF failed for testcase:\n(%+v, %d, %+v) => %+v", DF, 5, TF, got)
	}
}

func TestInversedDocumentFrequency(t *testing.T) {
	DF := DocumentFrequency{
		"a": {1},
		"b": {1, 2, 3},
	}
	got := inversedDocumentFrequency("a", DF, 15)

	if got != math.Log(15/float64(2)) {
		t.Errorf("TestInversedDocumentFreq failed for testcase:\n(%s, %+v, %d) => %+v", "a", DF, 15, got)
	}
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

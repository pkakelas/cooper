package main

import (
	"math"
	"strings"
)

// DocumentFrequency is the Document Frequency Map
type DocumentFrequency = map[string][]int

// TermFrequency is the Document Frequency Map
type TermFrequency = map[string]float64

func generateTermFrequency(doc string) (TF TermFrequency) {
	words := strings.Fields(doc)
	docsize := len(words)
	TF = map[string]float64{}

	for _, word := range words {
		if _, ok := TF[word]; ok {
			TF[word] += 1 / float64(docsize)
		} else {
			TF[word] = 1 / float64(docsize)
		}
	}

	return TF
}

func populateDF(DF DocumentFrequency, docID int, TF TermFrequency) DocumentFrequency {
	for word := range TF {
		if _, ok := DF[word]; ok {
			DF[word] = append(DF[word], docID)
		} else {
			DF[word] = []int{docID}
		}
	}

	return DF
}

func inversedDocumentFrequency(word string, DF DocumentFrequency, docsCount int) float64 {
	occurencies := float64(len(DF[word]))

	ret := math.Log(float64(docsCount) / (1 + occurencies))

	if ret < 0.0 {
		return 0.0
	}

	return ret
}

func generateTfIdfForWord(word string, tf map[string]float64, DF DocumentFrequency, docsCount int) float64 {
	return tf[word] * inversedDocumentFrequency(word, DF, docsCount)
}

package main

import (
	"math"

	"github.com/gonum/floats"
)

// DocumentFrequency is the Document Frequency Map
type DocumentFrequency = map[string][]string

// TermFrequency is the Term Frequency Map
type TermFrequency = map[string]float64

func generateTermFrequency(stems []string) (TF TermFrequency) {
	docsize := len(stems)
	TF = map[string]float64{}

	for _, stem := range stems {
		if _, ok := TF[stem]; ok {
			TF[stem] += 1 / float64(docsize)
		} else {
			TF[stem] = 1 / float64(docsize)
		}
	}

	return TF
}

func populateDF(DF DocumentFrequency, document Document) DocumentFrequency {
	for word := range document.tf {
		if _, ok := DF[word]; ok {
			DF[word] = append(DF[word], document.id)
		} else {
			DF[word] = []string{document.id}
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

func generateTfIdfForWord(term string, termFrequency float64, DF DocumentFrequency, docsCount int) float64 {
	return termFrequency * inversedDocumentFrequency(term, DF, docsCount)
}

func calculateTfIdfForDoc(terms []string, TF TermFrequency, DF DocumentFrequency, docsCount int) []float64 {
	TFIDF := make([]float64, len(terms))

	for idx, term := range terms {
		TFIDF[idx] = generateTfIdfForWord(term, TF[term], DF, docsCount)
		idx++
	}

	return TFIDF
}

// Vector length Math.sqrt(d[0]^2 + d[1]^2 + ..)
func vectorLength(v []float64) float64 {
	length := 0.0

	for _, dim := range v {
		length += dim * dim
	}

	return math.Sqrt(length)
}

// Cosine Similarity (d1, d2) =  Dot product(d1, d2) / ||d1|| * ||d2||
func cosineSimilarity(d1 []float64, d2 []float64) float64 {
	return floats.Dot(d1, d2) / (vectorLength(d1) * vectorLength(d1))
}

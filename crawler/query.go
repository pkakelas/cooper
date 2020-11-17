package main

import (
	"math"
	"sort"
	"strings"
)

func makeQuery(query string, state State) []QueryResult {
	terms := strings.Fields(query)
	stems := stemize(terms)
	results := calculateSimilarity(stems, state)

	sort.SliceStable(results, func(i, j int) bool {
		return results[i].sim > results[j].sim
	})

	return results
}

func calculateSimilarity(words []string, state State) []QueryResult {
	TF := generateTermFrequency(words)
	similarities := []QueryResult{}

	// Calculate tf-idf for query
	queryTFIDF := calculateTfIdfForDoc(words, TF, state.DF, len(state.documents))

	//Calculate tf-idf for all documents
	for _, doc := range state.documents {
		tfidf := calculateTfIdfForDoc(words, doc.tf, state.DF, len(state.documents))
		sim := cosineSimilarity(tfidf, queryTFIDF)

		if math.IsNaN(sim) {
			continue
		}

		similarities = append(similarities, QueryResult{
			ID:    doc.id,
			Title: doc.title,
			URL:   doc.url,
			sim:   cosineSimilarity(tfidf, queryTFIDF),
		})
	}

	return similarities
}

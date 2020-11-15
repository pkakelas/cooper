package main

import (
	"math"
	"sort"
	"strings"
)

func makeQuery(query string, documents []Document, DF DocumentFrequency) []QueryResult {
	terms := strings.Fields(query)
	stems := stemize(terms)
	results := calculateSimilarity(stems, documents, DF)

	sort.SliceStable(results, func(i, j int) bool {
		return results[i].sim > results[j].sim
	})

	return results
}

func calculateSimilarity(words []string, documents []Document, DF DocumentFrequency) []QueryResult {
	TF := generateTermFrequency(words)
	similarities := []QueryResult{}

	// Calculate tf-idf for query
	queryTFIDF := calculateTfIdfForDoc(words, TF, DF, len(documents))

	//Calculate tf-idf for all documents
	for _, doc := range documents {
		tfidf := calculateTfIdfForDoc(words, doc.tf, DF, len(documents))
		sim := cosineSimilarity(tfidf, queryTFIDF)

		if math.IsNaN(sim) {
			continue
		}

		similarities = append(similarities, QueryResult{
			id:    doc.id,
			title: doc.title,
			url:   doc.url,
			sim:   cosineSimilarity(tfidf, queryTFIDF),
		})
	}

	return similarities
}

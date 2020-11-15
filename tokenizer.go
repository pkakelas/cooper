package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	"github.com/kljensen/snowball"
)

//TODO: Find a stem library which handles non-english words
func extractStems(document *goquery.Document) (parsed []string) {
	text := extractText(document)
	words := strings.Fields(text)

	for _, word := range words {
		text := strings.TrimFunc(word, trimAllButLetters)

		stemmed, err := snowball.Stem(text, "english", true)
		if err != nil || len(stemmed) == 0 {
			fmt.Println("[TEXT_PARSER] Word couldn't be parsed:", word, err)
			continue
		}

		parsed = append(parsed, stemmed)
	}

	return
}

//TODO: Convert relative urls to absolute
func extractLinks(document *goquery.Document) (urls []string) {
	document.Find("a").Each(func(i int, s *goquery.Selection) {
		if href, ok := s.Attr("href"); ok {
			urls = append(urls, href)
		}
	})

	return
}

//TODO: Fix filtering of scripts when extracting text
func extractText(document *goquery.Document) string {
	document.Find("body script").Remove()
	document.Find("body frame").Remove()

	return document.Find("body").Text()
}

func extractTitle(document *goquery.Document) string {
	return document.Find("title").Text()
}

func trimAllButLetters(r rune) bool {
	return !unicode.IsLetter(r)
}

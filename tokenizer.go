package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	"github.com/kljensen/snowball"
)

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

func extractLinks(document *goquery.Document) (urls []string) {
	document.Find("a").Each(func(i int, s *goquery.Selection) {
		if href, ok := s.Attr("href"); ok {
			urls = append(urls, href)
		}
	})

	return
}

func extractText(document *goquery.Document) (text string) {
	htmlTags := "h1, h2, h3, h3, h5, h6, p, li, a"

	document.Find(htmlTags).Each(func(i int, s *goquery.Selection) {
		text += " " + s.Text()
	})

	return
}

func trimAllButLetters(r rune) bool {
	return !unicode.IsLetter(r)
}

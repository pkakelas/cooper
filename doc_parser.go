package main

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	"github.com/kljensen/snowball"
	"github.com/rs/xid"
	"golang.org/x/net/html"
)

// Document represents a parsed HTML file
type Document struct {
	tf        TermFrequency
	id        string
	title     string
	url       string
	neighbors []string
	stems     []string
}

func parseGoQueryDocument(url string, document *goquery.Document) Document {
	stems := extractStems(document)

	return Document{
		id:        xid.New().String(),
		title:     extractTitle(document),
		neighbors: extractLinks(url, document),
		tf:        generateTermFrequency(stems),
		url:       url,
		stems:     stems,
	}
}

//TODO: Find a stem library which handles non-english words
func extractStems(document *goquery.Document) []string {
	text := extractText(document)
	terms := strings.Fields(text)

	return stemize(terms)
}

func stemize(terms []string) (stems []string) {
	for _, term := range terms {
		text := strings.TrimFunc(term, trimAllButLetters)

		stemmed, err := snowball.Stem(text, "english", true)
		if err != nil || len(stemmed) == 0 {
			// Uncomment the log below for lots of spamming
			// fmt.Println("[TEXT_PARSER] Word couldn't be parsed:", word, err)
			continue
		}

		stems = append(stems, stemmed)
	}

	return
}

func extractLinks(url string, document *goquery.Document) (urls []string) {
	document.Find("a").Each(func(i int, s *goquery.Selection) {
		if href, ok := s.Attr("href"); ok {
			if len(href) > 0 && href[1:] == "/" {
				// Convert relative to absolut path
				href = extractDomainFromURI(url) + href
			}

			urls = append(urls, href)
		}
	})

	return
}

// Heavily inspired by: https://stackoverflow.com/a/44454014
func extractText(document *goquery.Document) string {
	str := ""
	code, _ := document.Html()
	domDocTest := html.NewTokenizer(strings.NewReader(code))

	previousStartToken := domDocTest.Token()
loop:
	for {
		tt := domDocTest.Next()
		switch {
		case tt == html.ErrorToken:
			break loop
		case tt == html.StartTagToken:
			previousStartToken = domDocTest.Token()
		case tt == html.TextToken:
			if previousStartToken.Data == "script" {
				continue
			}

			txtContent := strings.TrimSpace(html.UnescapeString(string(domDocTest.Text())))
			if len(txtContent) > 0 {
				str = str + " " + txtContent
			}
		}
	}
	return str
}

func extractTitle(document *goquery.Document) string {
	return document.Find("title").Text()
}

func extractDomainFromURI(url string) string {
	re := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`)
	return re.FindAllString(url, 1)[0]
}

func trimAllButLetters(r rune) bool {
	return !unicode.IsLetter(r)
}

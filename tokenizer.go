package main

import "github.com/PuerkitoBio/goquery"

func extractText(document *goquery.Document) (urls []string) {
	htmlTags := "h1, h2, h3, h3, h5, h6, p, li, a"

	document.Find(htmlTags).Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		urls = append(urls, text)
	})

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

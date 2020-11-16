package main

import (
	"net/url"
	"os"
	"strings"
	"unicode"
)

func existsInStringSlice(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func trimAllButLetters(str string) string {
	return strings.TrimFunc(str, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func isValidURI(uri string) bool {
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		return false
	}
	return true
}

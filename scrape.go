package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GenerateStaticFiles() {
	Directory("Juicestus/TorqueLearn/tree/master/pages/")
}

func Directory(path string) {
	fmt.Printf("[PATH] %s\n", path)

	res, err := http.Get("https://github.com/" + path)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	entries := doc.Find(".js-navigation-open")

	dir := false
	for i := range entries.Nodes {
		if i > 4 {
			href, exists := entries.Eq(i).Attr("href")
			if !exists {
				continue
			}

			Directory(href)
			dir = true
		}
	}

	if !dir {
		File(path)
	}
}

func File(path string) string {
	fmt.Printf("\t[FILE] %s\n", path)

	path = strings.Replace(path, "/blob/", "/", 1)
	res, err := http.Get("https://raw.githubusercontent.com/" + path)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	return doc.Find("html").Eq(0).Text()
}

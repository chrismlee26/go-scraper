package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Image struct {
	URL   string
	Title string
	alt   string
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	//  Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// Find link using an attribute selector
		// Matches any element that includes href=""
		link := e.Attr("href")

		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)

		// Visit link
		e.Request.Visit(link)
	})

	// Output to terminal
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("https://www.thesislabs.com")
}

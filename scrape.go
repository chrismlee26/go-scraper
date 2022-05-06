package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

// Struct for all URL's
type Url struct {
	Text string
	Link string
}

// Instantiate variable for single URL
var (
	FileURL string
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Slice to store URLs with length of 100
	var foundLinks = make([]Url, 100)
	fmt.Printf("foundLinks: \tLen: %v \tCap: %v\n", len(foundLinks), cap(foundLinks))

	// Instantiate default collector
	c := colly.NewCollector(
		// Limit Domain (Prevent travel to external sites)
		colly.AllowedDomains("thesislabs.com"),
	)

	// Limit Rules to prevent getting banned by the site
	c.Limit(&colly.LimitRule{
		DomainGlob: "*thesislabs.com",
		Delay:      2 * time.Second,
	})

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)

		if len(foundLinks) < 100 {
			foundLink := Url{
				Text: e.Text,
				Link: link,
			}
			// append(foundLinks, foundLink)
		}
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

	// Start scraping
	c.Visit("https://thesislabs.com")

	// All logic goes here after visit.
	// Create file name and folder name
	// createFileName()

	// Find all image links
	// c.OnHTML("img", func(e *colly.HTMLElement) {
	// 	link := e.Attr("src")
	// 	// store image links in array
	// 	foundImage := Image{
	// 		URL: link,
	// 	}
	// 	append(foundImages, foundImage)
	// })

	// Export links to json file
	// exportToJSON()
}

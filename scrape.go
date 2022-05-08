package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gocolly/colly"
)

// Struct for all URL's
type Url struct {
	Text string
	Link string
}

// Helper Functions

// Check Error
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// JSON Conversion

// Append to File

func main() {
	// Create Slice to store URLs, length 100
	var appendLinks = make([]Url, 100)
	var linkByTest = map[string]string{}

	// Create a new collector
	// foundLink := []Url{}
	fmt.Printf("foundLinks: \tLen: %v \tCap: %v\n", len(appendLinks), cap(appendLinks))

	// Instantiate default collector
	c := colly.NewCollector(
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

		linkByTest[e.Text] = link
		// appendFile(link)
		// append(appendLinks, foundLink{
		// 	Text: e.Text,
		// 	Link: link,
		// })
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
	// Find all image links
	// appendLinks()

	// Export links to json file
	// exportToJSON()

	linkByTextJSON, err := json.MarshalIndent(linkByTest, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(linkByTextJSON))

	ioutil.WriteFile("linksByTest.json", linkByTextJSON, 0777)
}

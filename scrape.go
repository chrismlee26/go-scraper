package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gocolly/colly"
)

//  ---------------------
// Struct for all URL's
//  ---------------------
type Url struct {
	Text string
	Link string
}

//  ---------------------
//  Check Error
//  ---------------------
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

//  ---------------------
//  Main
//  ---------------------
func main() {
	// Create Slice to store URLs, length 100
	var appendLinks = make([]Url, 100)
	var linkByText = map[string]string{}
	var images []string

	// Create a new collector
	// foundLink := []Url{}
	fmt.Printf("foundLinks: \tLen: %v \tCap: %v\n", len(appendLinks), cap(appendLinks))

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("thesislabs.com", "images.squarespace-cdn.com", "https://thesislabs.com/web", "https://thesislabs.com/interiors", "https://thesislabs.com/fashion"),
	)

	// Limit Rules to prevent getting banned by the site
	c.Limit(&colly.LimitRule{
		DomainGlob: "thesislabs.com*",
		Delay:      0 * time.Second,
	})

	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		src := e.Attr("src")
		fmt.Printf("Image found: %q => %s \n", e.Text, src)

		images = append(images, src)
		appendLinks = append(appendLinks, Url{Text: e.Text, Link: src})

		// Debug
		// fmt.Println("~~~~~~~~~~~~~~~~~~~~~~\n")
		// fmt.Printf("%v\n", linkByText)
		// fmt.Printf("%v\n", images)
		// fmt.Printf("%v\n", len(linkByText))
		// fmt.Printf("%v\n", len(images))
		// fmt.Printf("%v\n", cap(images))
		// fmt.Println("~~~~~~~~~~~~~~~~~~~~~~\n")
	})

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)

		linkByText[link] = e.Text

		// Go to link found on page
		nextPage := "https://thesislabs.com" + e.Request.AbsoluteURL(e.Attr("href"))
		c.Visit(nextPage)
	})

	//  ---------------------
	// Output to terminal
	//  ---------------------
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)

		if r.Request.URL.String() == "images.squarespace-cdn.com" {
			fmt.Println("Found image:", r.Request.URL)
			// Append to slice that goes to output.json
			images = append(images, r.Request.URL.String())

			// TODO: Download images in link directory
			// Why won't it download????
			// ioutil.WriteFile(r.Request.URL.String(), r.Body, 0777)
			// os.Create(r.Request.URL.String() + ".jpg")
			// fmt.Println("Saving", r.Request.URL)
			// if err := os.Rename(r.Request.URL.String(), r.Request.URL.String()+".jpg"); err != nil {
			// 	panic(err)
			// }

			fmt.Println("Error, not saving", r.Request.URL)
		}

	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	//  ---------------------
	// Start scraping
	//  ---------------------
	// May need a piece of logic to extend to other links instead of hardcoding these :S

	// TODO: auto select links
	c.Visit("https://thesislabs.com")

	//  ---------------------
	// Convert links to JSON
	//  ---------------------
	linkByTextJSON, err := json.MarshalIndent(linkByText, "", "  ")
	checkError(err)
	ioutil.WriteFile("output.json", linkByTextJSON, 0777)
}

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
// Check Error
//  ---------------------
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

//  ---------------------
// Main()
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
		colly.AllowedDomains("thesislabs.com", "images.squarespace-cdn.com"),
	)

	// Limit Rules to prevent getting banned by the site
	c.Limit(&colly.LimitRule{
		DomainGlob: "thesislabs.com*",
		Delay:      2 * time.Second,
	})

	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		src := e.Attr("src")
		fmt.Printf("Image found: %q => %s \n", e.Text, src)

		images = append(images, src)
	})

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)

		linkByText[e.Text] = link
		absoluteURL := e.Request.AbsoluteURL(link)
		c.Visit(absoluteURL)
		// appendFile(link)
		// append(appendLinks, foundLink{
		// 	Text: e.Text,
		// 	Link: link,
		// })
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
			// Why won't it download????
			ioutil.WriteFile("images.squarespace-cdn.com"+""+r.Request.URL.String(), r.Body, 0777)
			fmt.Println("Saving", r.Request.URL)
		}

	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	//  ---------------------
	// Start scraping
	//  ---------------------
	// May need a piece of logic to extend to other links

	c.Visit("https://thesislabs.com")
	c.Visit("https://thesislabs.com/web")
	c.Visit("https://thesislabs.com/interiors")
	c.Visit("https://thesislabs.com/fashion")

	//  ---------------------
	// Convert links to JSON
	//  ---------------------

	linkByTextJSON, err := json.MarshalIndent(linkByText, "", "  ")
	checkError(err)

	fmt.Println(string(linkByTextJSON))

	ioutil.WriteFile("output.json", linkByTextJSON, 0777)

	imagesBySrcJSON, err := json.MarshalIndent(images, "", "  ")
	checkError(err)

	fmt.Println(string(imagesBySrcJSON))

	ioutil.WriteFile("imagesBySrc.json", linkByTextJSON, 0777)

	//  ---------------------
	// Download Image Links
	//  ---------------------
	for _, image := range images {
		c.Visit(image)
	}

	//  ---------------------
	// Save Images to Disk
	//  ---------------------

}

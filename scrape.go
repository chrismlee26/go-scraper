package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

// Struct for data to return
type Data struct {
	Text  []Text
	Image []Image
}

// Text Struct
type Text struct {
	MainText string
	SubText  string
}

// Image Struct
type Image struct {
	URL   string
	Title string
	alt   string
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	c.Limit(&colly.LimitRule{
		DomainGlob: "*",
		Delay:      2 * time.Second,
	})

	// Limit Domain (Prevent travel to external sites)
	c.AllowedDomains = []string{"thesislabs.com"}

	//  Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		e.Request.Visit(link)
	})

	// Visit site and print image links
	// c.OnHTML("img", func(e *colly.HTMLElement) {
	// 	link := e.Attr("src")
	// 	fmt.Printf("%q\n", e.Attr("src"))
	// 	e.Request.save(link)
	// })

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

	c.Visit("https://thesislabs.com")
}

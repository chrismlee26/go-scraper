package main

import (
	"fmt"
	"net/url"
	"strings"
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

var (
	folderName   string
	fileName     string
	foundFileUrl string
)

// func createFolders() {
// 	// Create folders
// 	file, err := os.Create()
// 	checkError(err)
// 	return file
// }

func createFileName() {
	// change foundFileUrl into image paths found by scraper
	fileURL, err := url.Parse(foundFileUrl)
	checkError(err)

	path := fileURL.Path
	splitPaths := strings.Split(path, "/")
	fileName = splitPaths(len(splitPaths) - 1)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
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

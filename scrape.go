package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"
	"encoding/json"

	"github.com/gocolly/colly"
)

// Struct for data to return
type Data struct {
	Text  []Text
	Image []Image
}

// Text Struct
type Url struct {
	Text string
	Link  string
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

foundLinks := make([]Url, 100)

func createFileName() {
	// change foundFileUrl into image paths found by scraper
	fileURL, err := url.Parse(foundFileUrl)
	checkError(err)

	path := fileURL.Path
	splitPaths := strings.Split(path, "/")

	// Filename for downloaded items
	fileName = splitPaths[len(splitPaths)-1]

	// Folder name for downloaded items
	folderName = strings.Join(splitPaths[:len(splitPaths)-1], "/")
	// os.Mkdir(folderName, 0777)
	os.MkdirAll(folderName, os.ModePerm)
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
		// fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		foundLink := Url{
			Text: e.Text,
			Link: link,
		}
		append(foundLinks, foundLink)
			

		// Visit link found on page
		// e.Request.Visit(link)
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

	// Start scraping
	c.Visit("https://thesislabs.com")

	// All logic goes here after visit. 
	// Create file name and folder name
	createFileName()
	
	// Find all image links
	c.OnHTML("img", func(e *colly.HTMLElement) {
		link := e.Attr("src")
	// store image links in array
		foundImage := Image{
			URL: link,
		}
		append(foundImages, foundImage)
	})
	// Export links to json file
	exportToJSON()
}

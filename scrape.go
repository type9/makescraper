package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Headline struct {
	Title  string
	Author string
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	fName := "headlines.json"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	// Instantiate default collector
	c := colly.NewCollector()

	headlines := make([]Headline, 0, 100)
	// On every title element
	c.OnHTML(".c-card__header", func(e *colly.HTMLElement) {
		author := e.ChildText(".c-card__byline")
		headline := Headline{
			e.Text,
			author,
		}
		headlines = append(headlines, headline)
		// Print link
		fmt.Printf("Title Found: %q by %s\n", e.Text, author)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.rollingstone.com/music/music-news/")

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	// Dump json to the standard output
	enc.Encode(headlines)
}

package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/PuerkitoBio/goquery"
)

// Scrape the latest headlines from a news website
func scrapeHeadlines(url string) {
    // Make an HTTP request
    res, err := http.Get(url)
    if err != nil {
        log.Fatalf("Failed to make request: %v", err)
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        log.Fatalf("Failed to fetch data: %d %s", res.StatusCode, res.Status)
    }

    // Parse the HTML document
    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        log.Fatalf("Failed to parse HTML: %v", err)
    }

    // Find and print article titles
    doc.Find("article h2").Each(func(i int, s *goquery.Selection) {
        title := s.Text()
        fmt.Printf("%d: %s\n", i+1, title)
    })
}

func main() {
    // URL to scrape (for example, replace this with any website you want to scrape)
    url := "https://example.com/news"
    fmt.Println("Scraping headlines from:", url)
    scrapeHeadlines(url)
}
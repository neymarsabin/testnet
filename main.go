package main

import (
	"fmt"
	colly "github.com/gocolly/colly"
	"time"
)

type Speed struct {
	value     int
	timestamp time.Time
}

func main() {
	fmt.Println("Hello world")

	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("http://go-colly.org/")
}

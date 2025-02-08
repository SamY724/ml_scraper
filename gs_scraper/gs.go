package gs_scraper

import (
	"fmt"
	"github.com/gocolly/colly"
)

type GSArtifact struct {
	source string
	date_posted string
	paywall bool
}


func ScrapeGS(s string) int {

	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	return 6
}
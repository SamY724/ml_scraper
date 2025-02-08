package main


import (
	"fmt"
	"github.com/SamY724/ml_scraper/hn_scraper"
	"github.com/SamY724/ml_scraper/gs_scraper"
)

func main(){

	fmt.Println("Hai")
	fmt.Println(hn_scraper.ScrapeHN("hello"))
	fmt.Println(gs_scraper.ScrapeGS("hello"))

}
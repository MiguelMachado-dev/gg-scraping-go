package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/gocolly/colly"
)

type Stats struct {
	Players string
}

func CreateLog(players string) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + players + " playing\n")

	file.Close()
}

func GuiltyGearScraper() {
	fmt.Println("Start scraping")

	c := colly.NewCollector(
		colly.AllowedDomains("steamcharts.com"),
	)

	c.OnHTML("#app-heading div.app-stat:first-of-type", func(e *colly.HTMLElement) {
		stats := Stats{}
		stats.Players = e.ChildText(".num")
		println(time.Now().Format("02/01/2006 15:04:05") + " - " + stats.Players + " playing\n")
		CreateLog(stats.Players)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status code:", r.StatusCode)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://steamcharts.com/app/1384160")
}

func main() {
	my_scheduler := gocron.NewScheduler(time.UTC)
	my_scheduler.Every(2).Minute().Do(GuiltyGearScraper)
	my_scheduler.StartBlocking()
}

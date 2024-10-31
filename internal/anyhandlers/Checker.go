package anyHandlers

import (
	"log/slog"
	"net/url"
	"sync"

	"github.com/gocolly/colly"
)

func IsLink(msg string) bool {

	_, err := url.Parse(msg)

	return err == nil
}

func ScrapHeader(url string) string {
	var title string

	c := colly.NewCollector()
	var wg sync.WaitGroup

	wg.Add(1)

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		title = e.Text
		wg.Done()
	})

	c.OnError(func(r *colly.Response, err error) {
		if err != nil {
			slog.Error("Scraper error:", err)
		}
	})

	c.Visit(url)
	wg.Wait()
	return title

}

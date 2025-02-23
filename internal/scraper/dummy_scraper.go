package scraper

import (
	"fmt"
	"math/rand"
	"scraper/internal/bus"
	"scraper/internal/object"
)

type Scraper interface {
	Scrape(job object.Job)
}

type DummyScraper struct {
	bus bus.Bus
}

func NewDummyScraper(bus_object bus.Bus) *DummyScraper {
	return &DummyScraper{
		bus: bus_object,
	}
}

func (d *DummyScraper) Scrape(job object.Job) {
	chance := rand.Intn(100)
	error_statuses := []int{
		401,
		403,
		404,
		500,
		502,
	}

	// simulate a 10% chance of failure
	if chance < 10 {
		d.bus.Error <- object.NewErrorData(
			fmt.Errorf("Failed to fetch data"),
			"",
			job.GetUrl(),
			error_statuses[rand.Intn(len(error_statuses))],
		)
		return
	}

	d.bus.ScrapedData <- object.NewData(
		job.GetUrl(),
		"Dummy data",
		200,
		map[string]any{
			"data": "Dummy data",
			"val":  123,
		},
	)
}

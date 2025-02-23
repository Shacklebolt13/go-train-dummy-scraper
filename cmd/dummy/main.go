package main

import (
	"scraper/config"
	"scraper/internal/bus"
	"scraper/internal/scheduler"
	"scraper/internal/scraper"
)

func main() {
	conf := config.DefaultConfig()
	conf.ParseFromArgs()
	defer conf.Close()

	size := len(conf.GetUrlsToScrape())

	busObject := bus.NewBus(size)
	scraperObject := scraper.NewDummyScraper(*busObject)

	schedulerObject := scheduler.NewGoRoutineScheduler(
		scraperObject,
		*busObject,
		conf,
	)

	for _, job := range conf.GetUrlsToScrape() {
		schedulerObject.AppendJob(job)
	}

	schedulerObject.Start()
}

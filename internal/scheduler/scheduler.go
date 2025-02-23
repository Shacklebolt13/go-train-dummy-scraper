package scheduler

import (
	"fmt"
	"scraper/config"
	"scraper/internal/bus"
	"scraper/internal/object"
	"scraper/internal/scraper"
	"sync"
)

type Scheduler interface {
	Start()
	AppendJob(job object.Job)
}

type GoRoutineScheduler struct {
	scraper scraper.Scraper
	wg      sync.WaitGroup
	bus     bus.Bus
	conf    config.Config
	jobs    []object.Job
}

func NewGoRoutineScheduler(scraper scraper.Scraper, bus bus.Bus, conf config.Config) *GoRoutineScheduler {
	return &GoRoutineScheduler{
		scraper: scraper,
		bus:     bus,
		conf:    conf,
	}
}

func (s *GoRoutineScheduler) Start() {
	go func() {
		s.listener()
		s.wg.Done()
	}()
	for _, job := range s.jobs {
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			s.scraper.Scrape(job)
		}()
	}
	s.wg.Wait()
	s.wg.Add(1)
	s.bus.Close()
	s.wg.Wait()
}

func (s *GoRoutineScheduler) AppendJob(job object.Job) {
	s.jobs = append(s.jobs, job)
}

func (s *GoRoutineScheduler) listener() {
	for {
		select {
		case scrapedData := <-s.bus.ScrapedData:
			if scrapedData == nil {
				return
			}
			s.conf.GetRawDataHandler().HandleData(scrapedData)
			fmt.Printf("\n[%d] Scraped data [%s] : %v", scrapedData.GetStatus(), scrapedData.GetPageURL(), scrapedData.GetDataMap())
		case errData := <-s.bus.Error:
			if errData == nil {
				return
			}
			s.conf.GetRawDataHandler().HandleData(errData)
			fmt.Printf("\n[%d] Error while scraping [%s] : %v", errData.GetStatus(), errData.GetPageURL(), errData.GetError())
		}
	}
}

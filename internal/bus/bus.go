package bus

import "scraper/internal/object"

type Bus struct {
	// scraped data channel
	ScrapedData chan object.PageData
	// error channel
	Error chan object.ErrorData
}

func NewBus(jobSize int) *Bus {
	return &Bus{
		ScrapedData: make(chan object.PageData),
		Error:       make(chan object.ErrorData),
	}
}

func (b *Bus) Close() {
	close(b.ScrapedData)
	close(b.Error)
}
